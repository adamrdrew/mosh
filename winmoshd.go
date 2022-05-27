package main

// This page was very helpful here https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/player"
	"github.com/adamrdrew/mosh/server"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

var musicPlayer player.Player

func setupLog() *os.File {
	f, err := os.OpenFile(os.Getenv("MOSH_LOG_DIR")+"/winmoshd.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return f
}

//Entrypoint for the daemon
func main() {
	fmt.Println("Starting winmoshd...")

	logFile := setupLog()
	defer logFile.Close()
	log.SetOutput(logFile)

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
	//Number of the beast... with a 9 in front of it.
	log.Fatal(http.ListenAndServe("0.0.0.0:"+config.GetPort(), nil))
	for {
	}
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
