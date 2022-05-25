package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adamrdrew/mosh/ipc"
	"github.com/sevlyar/go-daemon"
)

//Entrypoint for the daemon
func main() {
	cntxt := &daemon.Context{
		PidFileName: "moshd.pid",
		PidFilePerm: 0644,
		LogFileName: "moshd.log",
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

	startListener()
}

//Listens for HTTP requests on port 9666 and passes them off to httpListener
func startListener() {
	log.Print("Starting listener....")
	http.HandleFunc("/listener", httpListener)
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

//Takes a message and dispatches to the right code to handle it
func handleMessage(message ipc.Message) ipc.Response {
	response := ipc.Response{}
	response.Code = "OK"

	switch message.Command {
	case "play":
		response.Message = "Playing song"
		log.Print("Play song: ", message.Data)
	case "stop":
		response.Message = "Stopped song"
		log.Print("Stop song")
	case "pause":
		response.Message = "Paused Song"
		log.Print("Pause song")
	case "previous":
		response.Message = "Previous song"
		log.Print("Previous song")
	case "next":
		response.Message = "Next song"
		log.Print("Next song")
	case "status":
		response.Message = "Status"
		log.Print("Get status")
	case "default":
		response.Code = "UNKNOWN"
		response.Message = "Unknown command recieved."
	}

	return response
}
