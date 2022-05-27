package main

// This page was very helpful here https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/player"
	"github.com/adamrdrew/mosh/server"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/sevlyar/go-daemon"
)

var musicPlayer player.Player

//Entrypoint for the daemon
func main() {
	fmt.Println("Starting moshd...")
	cntxt := &daemon.Context{
		PidFileName: config.GetPidDir() + "moshd.pid",
		PidFilePerm: 0644,
		LogFileName: config.GetLogir() + "moshd.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("moshd started")

	musicPlayer = player.Player{}
	musicPlayer.Init()

	defaultSampleRate := beep.SampleRate(44100)
	speaker.Init(defaultSampleRate, defaultSampleRate.N(time.Second/10))

	startListener()
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

	response.Add(responseItem)
	return response
}
