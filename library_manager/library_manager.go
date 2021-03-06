package library_manager

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/server"
)

const MUSIC_LIB_TYPE = "artist"

func SetLibraryInConf(serv server.Server, conf *config.Config) LibrarySelector {
	ls := LibrarySelector{
		Config: conf,
		Server: serv,
	}
	//The library key is already in the config so we don't need to do
	//anything else
	if ls.isLibInConf() {
		return ls
	}
	ls.setLibInConf()
	return ls
}

type LibrarySelector struct {
	Server server.Server
	Config *config.Config
}

//Returns true if there's a library key in the config file
func (l *LibrarySelector) isLibInConf() bool {
	return l.Server.Config.Library != config.UNINITIALIZED
}

//Gets a list of libraries from the server and then filters them
//for music libraries
func (l *LibrarySelector) getMusicLibs() []responses.LibraryDirectory {
	libs := l.Server.GetLibraries()

	var libList []responses.LibraryDirectory

	for _, libDir := range libs.Directories {
		if libDir.Type == MUSIC_LIB_TYPE {
			libList = append(libList, libDir)
		}
	}

	return libList
}

//Returns an array of strings that represent valid indicies for music libraries in the list
func (l *LibrarySelector) getValidLibraryIndexes(musicLibs []responses.LibraryDirectory) []string {
	var validIndexes []string

	for i, lib := range musicLibs {
		stringIndex := fmt.Sprint(i)
		fmt.Println("    " + stringIndex + ") " + lib.Title)
		validIndexes = append(validIndexes, stringIndex)
	}

	return validIndexes
}

//Polls for user input to select a library by index
//If the user provides bogus input we force them to repeat
func (l *LibrarySelector) getUserSelection(musicLibs []responses.LibraryDirectory) responses.LibraryDirectory {
	validIndexes := l.getValidLibraryIndexes(musicLibs)

	validSelectionMade := false
	userInput := "0"

	for {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		userInput = input.Text()
		for _, validIndex := range validIndexes {
			validSelectionMade = userInput == validIndex
			if validSelectionMade {
				break
			}
		}
		if validSelectionMade {
			fmt.Println("Library set.")
			break
		} else {
			fmt.Println("Invalid selection. Please try again.")
		}
	}

	//I don't bother handling the error here because we sanitized the input
	i, _ := strconv.Atoi(userInput)

	return musicLibs[i]
}

func (l *LibrarySelector) setLibInConf() {
	fmt.Println("Select a music library to use.")
	fmt.Println("Type the number to the left of the name of the library you want.")
	musicLibs := l.getMusicLibs()

	selectedLib := l.getUserSelection(musicLibs)

	l.Config.Library = selectedLib.Key
	l.Config.Save()
}
