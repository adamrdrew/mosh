package auth

// https://forums.plex.tv/t/authenticating-with-plex/609370

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os/exec"
	"time"
)

var DEFAULT_HEADER = http.Header{
	"Content-Type": {"application/json"},
	"accept":       {"application/json"},
}

type AuthToken struct {
	Response string `json:"authToken"`
}

type Pin struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

func (p *Pin) GetID() string {
	return fmt.Sprint(p.ID)
}

type Auth struct {
	ID    string
	Name  string
	Token string
	Pin   Pin
}

func (a *Auth) CheckToken() {
	fmt.Println("Checking Plex authorization status...")
	checkTokenURL := "https://plex.tv/api/v2/user?" +
		"X-Plex-Product=" + a.Name +
		"&X-Plex-Client-Identifier=" + a.ID +
		"&X-Plex-Token=" + a.Token

	response, err := http.Get(checkTokenURL)
	a.panic(err)

	switch response.StatusCode {
	case 401:
		fmt.Println("    FAILED: Authorization required.")
		a.ObtainToken()
	case 200:
		fmt.Println("    SUCCESS: Authorized.")
	}

}

func (a *Auth) ObtainToken() {
	fmt.Println("Obtaining token...")
	a.ObtainPin()
	a.OpenPinURL()
	a.PollPIN()
}

func (a *Auth) printResponse(response *http.Response) {
	respDump, _ := httputil.DumpResponse(response, true)
	fmt.Printf("DEBUG RESPONSE:\n%s", string(respDump))
}

func (a *Auth) panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (a *Auth) ObtainPin() {
	obtainPINURL := "https://plex.tv/api/v2/pins?strong=true&X-Plex-Product=" + a.Name + "&X-Plex-Client-Identifier=" + a.ID

	client := http.Client{}
	req, err := http.NewRequest("POST", obtainPINURL, nil)
	a.panic(err)

	req.Header = DEFAULT_HEADER

	response, err := client.Do(req)
	a.panic(err)

	body, err := ioutil.ReadAll(response.Body)
	a.panic(err)

	pin := new(Pin)
	jsonError := json.Unmarshal(body, &pin)
	a.panic(jsonError)
	a.Pin = *pin
}

func (a *Auth) OpenPinURL() {
	url := "https://app.plex.tv/auth#?clientID=" + a.ID + "&code=" + a.Pin.Code + "&context%5Bdevice%5D%5Bproduct%5D=" + a.Name
	fmt.Println("")
	fmt.Println("Authorize Mosh with a web browser. It should open automatically. If it doesn't then open it manually with the following URL:")
	fmt.Println(url)
	fmt.Println("")
	exec.Command("gio open", url)
}

func (a *Auth) PollPIN() {
	authorized := false
	fmt.Println("Waiting for authorization...")
	url := "https://plex.tv/api/v2/pins/" + a.Pin.GetID() + "?code=" + a.Pin.Code + "&X-Plex-Client-Identifier=" + a.ID
	for start := time.Now(); time.Since(start) < (time.Second * 60); {
		print(".")
		var client = http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		a.panic(err)

		req.Header = DEFAULT_HEADER

		response, err := client.Do(req)
		a.panic(err)

		body, err := ioutil.ReadAll(response.Body)
		a.panic(err)

		var authToken = new(AuthToken)
		jsonError := json.Unmarshal(body, &authToken)
		a.panic(jsonError)

		authorized = authToken.Response != ""

		if authorized {
			a.printResponse(response)
			a.Token = authToken.Response
			break
		}

		time.Sleep(time.Second)
	}

	if authorized {
		fmt.Println("Authorized.")
	} else {
		fmt.Println("\nAuthorization timed out or failed. Exiting.")
	}

}
