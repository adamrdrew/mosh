package ipc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Command string   `json:"command"`
	Data    string   `json:"data"`
	Args    []string `json:"args"`
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

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func SendMessageToDaemon(message Message) Response {
	host := "http://0.0.0.0"
	port := "9666"
	route := "listener"
	url := host + ":" + port + "/" + route

	var client = http.Client{}

	body, _ := json.Marshal(message)

	req, reqErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
	maybePanic(reqErr)

	response, doErr := client.Do(req)
	maybePanic(doErr)

	body, readErr := ioutil.ReadAll(response.Body)
	maybePanic(readErr)

	var retval Response

	umErr := json.Unmarshal(body, &retval)
	maybePanic(umErr)

	return retval
}
