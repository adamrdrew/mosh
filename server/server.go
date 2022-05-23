package server

//https://github.com/Arcanemagus/plex-api/wiki/Plex-Web-API-Overview

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/adamrdrew/mosh/config"
)

func GetServer(config config.Config) Server {
	server := Server{
		Address: "",
		Port:    "",
		Config:  config,
	}
	server.getServerData()
	return server
}

type XMLMediaContainer struct {
	Server XMLServer `xml:"Server"`
}

type XMLServer struct {
	Name    string `xml:"name,attr"`
	Address string `xml:"address,attr"`
	Port    string `xml:"port,attr"`
}

type Server struct {
	Address string
	Port    string
	Config  config.Config
}

func (s *Server) panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (s *Server) getServerData() {
	var client = http.Client{}
	url := "https://plex.tv/pms/servers.xml?X-Plex-Token=" + s.Config.Token
	req, err := http.NewRequest("GET", url, nil)
	s.panic(err)

	response, err := client.Do(req)
	s.panic(err)

	body, err := ioutil.ReadAll(response.Body)
	s.panic(err)

	var serverResponse = new(XMLMediaContainer)
	xmlError := xml.Unmarshal(body, &serverResponse)
	s.panic(xmlError)

	s.Address = serverResponse.Server.Address
	s.Port = serverResponse.Server.Port
}
