package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MessageHandler struct {
}

func (m *MessageHandler) Handle(message IPCMessage) IPCResponse {
	response := IPCResponse{}
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

type IPCMessage struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

type IPCResponse struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	CurrentTime string `json:"currenttime"`
	TotalTime   string `json:"totaltime"`
}

func Listener(w http.ResponseWriter, r *http.Request) {
	log.Print("Message recieved...")

	var message IPCMessage

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		log.Fatal("Message decode failed: ", err)
	}

	handler := MessageHandler{}
	response := handler.Handle(message)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	// Add handle funcs for 3 pages.
	http.HandleFunc("/listener", Listener)

	log.Print("Starting listener....")

	// Run the web server.
	log.Fatal(http.ListenAndServe(":9666", nil))
}
