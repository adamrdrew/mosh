package server

//https://github.com/Arcanemagus/plex-api/wiki/Plex-Web-API-Overview

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/plex_urls"
	"github.com/adamrdrew/mosh/responses"
)

func GetServer(config *config.Config) Server {
	server := Server{
		Config:   config,
		PlexURLs: plex_urls.GetPlexURLs(config),
	}
	server.getServerData()
	return server
}

type Server struct {
	Config   *config.Config
	PlexURLs plex_urls.PlexURLs
}

func (s *Server) panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (s *Server) doGet(url string) []byte {
	var client = http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	s.panic(err)

	response, err := client.Do(req)
	s.panic(err)

	body, err := ioutil.ReadAll(response.Body)
	s.panic(err)

	return body

}

func (s *Server) getServerData() {
	//This URL isn't in PlexURLs because that type provides
	//Plex server queries. This is a plex.tv query. It is a one-off.
	url := "https://plex.tv/pms/servers.xml?X-Plex-Token=" + s.Config.Token

	body := s.doGet(url)

	var serverResponse = new(responses.ServerMediaContainer)
	xmlError := xml.Unmarshal(body, &serverResponse)
	s.panic(xmlError)

	s.Config.Address = serverResponse.Server.Address
	s.Config.Port = serverResponse.Server.Port
	s.Config.Save()
}

func (s *Server) GetLibraries() responses.LibraryMediaContainer {
	url := s.PlexURLs.GetLibraries()
	body := s.doGet(url)

	var serverResponse = new(responses.LibraryMediaContainer)
	xmlError := xml.Unmarshal(body, &serverResponse)
	s.panic(xmlError)

	return *serverResponse
}
