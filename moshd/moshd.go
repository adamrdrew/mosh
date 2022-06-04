package moshd

// This page was very helpful here https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	cachemanager "github.com/adamrdrew/mosh/cache_manager"
	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/player"
	"github.com/adamrdrew/mosh/server"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/sevlyar/go-daemon"
)

var musicPlayer player.Player
var cacheLastPruneTime time.Time

func StartWaitAndCheck() {
	StartDaemon()
	time.Sleep(1 * time.Second)
	CheckDaemonStatus()
}

func readPidFile() (string, error) {
	fileContent, err := ioutil.ReadFile(pidFile())
	if err != nil {
		return "", err
	}

	// Convert []byte to string
	text := string(fileContent)
	return text, nil
}

func pidFile() string {
	return config.GetPidDir() + "moshd.pid"
}

func CheckDaemonStatus() *os.Process {
	pid, err := readPidFile()
	if err != nil {
		fmt.Println("PID file not found. Daemon not running?")
		return nil
	}
	pidInt, convErr := strconv.ParseInt(pid, 10, 64)
	if convErr != nil {
		fmt.Println("PID file malformed. Something is wrong.")
		return nil
	}
	proc, err := os.FindProcess(int(pidInt))
	if err != nil || proc == nil {
		fmt.Println("Process not found. Daemon not running?")
		return nil
	}

	return proc
}

func KillDaemon() {
	proc := CheckDaemonStatus()

	if proc == nil {
		fmt.Println("Can't stat daemon. Something wrong.")
		return
	}

	// Kill the process
	kErr := proc.Kill()
	if kErr != nil {
		fmt.Println("Kill process failed. Something is wrong. Daemon not running?")
		return
	}

	os.Remove(pidFile())

	fmt.Println("OK - Daemon stopped.")
}

//Entrypoint for the daemon
func StartDaemon() {
	fmt.Println("Starting mosh daemon...")
	cntxt := &daemon.Context{
		PidFileName: pidFile(),
		PidFilePerm: 0644,
		LogFileName: config.GetLogir() + "moshd.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
		return
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("moshd started")
	fmt.Println("Mosh daemon started.")

	musicPlayer = player.Player{}
	musicPlayer.Init()

	defaultSampleRate := beep.SampleRate(44100)
	speaker.Init(defaultSampleRate, defaultSampleRate.N(time.Second/10))

	//Check and prune cache on daemon start
	cache := cachemanager.MakeCacheManager()
	cache.PruneCache()
	cacheLastPruneTime = time.Now()

	startListener()
}

func pruneCache() {
	//If the last time we pruned was after yesterday bail
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	if cacheLastPruneTime.After(yesterday) {
		return
	}
	log.Print("Need to prune cache...")
	cacheLastPruneTime = today
	cache := cachemanager.MakeCacheManager()
	cache.PruneCache()
}

//Listens for HTTP requests on port 9666 and passes them off to httpListener
func startListener() {
	log.Print("Starting listener....")
	http.HandleFunc("/listener", httpListener)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+config.GetPort(), nil))
}

//The HTTP handler function. Gets messages, decodes them, and sends them through
//the message handler.
func httpListener(w http.ResponseWriter, r *http.Request) {
	log.Print("Message recieved...")

	var message ipc.Message

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Fatal("Message decode failed: ", err)
	}

	response := handleMessage(message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func getServer() server.Server {
	conf := config.GetConfig()
	return server.GetServer(&conf)
}

//Takes a message and dispatches to the right code to handle it
func handleMessage(message ipc.Message) ipc.Response {
	response := ipc.Response{}
	responseItem := ipc.ResponseItem{}
	responseItem.Code = "OK"

	switch message.Command {
	case "queue-album":
		responseItem = musicPlayer.QueueAlbum(message.Data)
	case "play-queue":
		go musicPlayer.PlayQueue()
		responseItem.Message = "Playing: " + musicPlayer.NowPlayingSongString()
	case "now-playing":
		responseItem = musicPlayer.GetNowPlaying()
	case "get-queue":
		return musicPlayer.GetPlayQueue()
	case "stop":
		responseItem = musicPlayer.StopQueue()
	case "back":
		responseItem = musicPlayer.GoBackInQueue()
	case "next":
		responseItem = musicPlayer.GoForwardInQueue()
	}

	go pruneCache()

	response.Add(responseItem)
	return response
}
