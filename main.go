package main

import (
	"fmt"

	"github.com/adamrdrew/mosh/auth"
	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/server"
)

func main() {
	//Get or create and then get the local config
	conf := config.GetConfig()

	//Get the authorizer. This will auth to plex.tv if we don't have a valid token.
	authorizer := auth.GetAuthorizer(conf)

	//Set the token and save the config
	conf.SetToken(authorizer.Token)

	//Get the server info
	//TODO: If you have multiple servers you are SOL we're just using the first in the list
	server := server.GetServer(conf)
	fmt.Println(server)

}
