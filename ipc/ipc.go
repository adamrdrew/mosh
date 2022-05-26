package ipc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/adamrdrew/mosh/config"
)

//The type we use for sending messages to moshd
type Message struct {
	Command string   `json:"command"`
	Data    string   `json:"data"`
	Args    []string `json:"args"`
}

//The type we use for responses from moshd
type ResponseItem struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	Song        string `json:"song"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	CurrentTime string `json:"currenttime"`
	TotalTime   string `json:"totaltime"`
}

//In some rare cases we need to get multiple reponses from
//moshd at once
type Response struct {
	Responses []ResponseItem `json:"items"`
}

func (r *Response) First() ResponseItem {
	return r.Responses[0]
}

func (r *Response) Add(ri ResponseItem) {
	r.Responses = append(r.Responses, ri)
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func daemonTransport(message Message) []byte {
	host := "http://0.0.0.0"
	port := config.GetPort()
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

	return body
}

func SendMessageToDaemon(message Message) Response {

	body := daemonTransport(message)

	var retval Response

	umErr := json.Unmarshal(body, &retval)
	maybePanic(umErr)

	return retval
}
