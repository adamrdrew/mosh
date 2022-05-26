package main

// This page was very helpful here https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
	"encoding/json"
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
	cntxt := &daemon.Context{
		PidFileName: "moshd/moshd.pid",
		PidFilePerm: 0644,
		LogFileName: "moshd/moshd.log",
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
	//Number of the beast... with a 9 in front of it.
	log.Fatal(http.ListenAndServe(":9666", nil))
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
	response.Code = "OK"

	switch message.Command {
	case "queue-album":
		response = musicPlayer.QueueAlbum(message.Data)
	case "play-queue":
		go musicPlayer.PlayQueue()
		response.Message = "Playing: " + musicPlayer.NowPlayingSongString()
	case "now-playing":
		response = musicPlayer.GetNowPlaying()
	case "stop":
		response = musicPlayer.StopQueue()
	case "back":
		response = musicPlayer.GoBackInQueue()
	case "next":
		response = musicPlayer.GoForwardInQueue()
	}

	return response
}
