package ipc

type Message struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

type Response struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	CurrentTime string `json:"currenttime"`
	TotalTime   string `json:"totaltime"`
}
