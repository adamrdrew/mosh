package plex_urls

import (
	"github.com/adamrdrew/mosh/config"
)

const RT_LIBRARIES = "/library/sections"

func GetPlexURLs(conf config.Config) PlexURLs {
	um := PlexURLs{
		Config: conf,
	}
	return um
}

type PlexURLs struct {
	Config config.Config
}

func (u *PlexURLs) Server() string {
	return "http://" + u.Config.Address + ":" + u.Config.Port
}

func (u *PlexURLs) GetLibraries() string {
	return u.makeURL(RT_LIBRARIES)
}

func (u *PlexURLs) makeURL(route string) string {
	return u.withToken(u.Server() + route)
}

func (u *PlexURLs) withToken(url string) string {
	return url + "?X-Plex-Token=" + u.Config.Token
}
