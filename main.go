package main

import (
	"fmt"
	"os"

	"github.com/adamrdrew/mosh/auth"
	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/library_manager"
	"github.com/adamrdrew/mosh/server"
)

func main() {
	fmt.Println("Welcome to MOSH! 🎧💿🐧")

	//Get or create and then get the local config
	conf := config.GetConfig()

	//Get the authorizer. This will auth to plex.tv if we don't have a valid token.
	authorizer := auth.GetAuthorizer(&conf)

	if !authorizer.Authorized {
		fmt.Println("Authorization failed. Quitting.")
		os.Exit(1)
	}

	//Set the token and save the config
	conf.SetToken(authorizer.Token)

	//Get the server info
	//TODO: If you have multiple servers you are SOL we're just using the first in the list
	server := server.GetServer(&conf)

	library_manager.GetLibraryManager(server, &conf)

	fmt.Println("We are ready to ROCK!!!!")

}
