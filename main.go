package main

import (
	"github.com/adamrdrew/mosh/auth"
	"github.com/adamrdrew/mosh/config"
)

var APP_ID = "f79d7735-864b-4ed7-a6dc-a3971824843b"
var APP_NAME = "Mosh"

func main() {
	conf := config.Config{}
	conf.Load()

	authorizer := auth.Auth{
		ID:    APP_ID,
		Name:  APP_NAME,
		Token: conf.Token,
	}
	authorizer.CheckToken()

	conf.Token = authorizer.Token
	conf.Save()

}
