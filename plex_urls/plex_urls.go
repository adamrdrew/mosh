package plex_urls

import (
	"github.com/adamrdrew/mosh/config"
)

const RT_LIBRARIES = "/library/sections"
const XPLEXTOKEN = "X-Plex-Token="
const RT_ARTIST_SEARCH = "/all?title="
const RT_ALBUM_SEARCH = "/albums?title="

func GetPlexURLs(conf *config.Config) PlexURLs {
	um := PlexURLs{
		Config: conf,
	}
	return um
}

type PlexURLs struct {
	Config *config.Config
}

func (u *PlexURLs) Server() string {
	return "http://" + u.Config.Address + ":" + u.Config.Port
}

func (u *PlexURLs) GetLibraries() string {
	return u.makeURL(RT_LIBRARIES)
}

func (u *PlexURLs) SearchArstists(query string) string {
	return u.libraryQuery(RT_ARTIST_SEARCH, query)
}

func (u *PlexURLs) SearchAlbums(query string) string {
	return u.libraryQuery(RT_ALBUM_SEARCH, query)
}

func (u *PlexURLs) libraryQuery(endpoint string, query string) string {
	return u.Server() + RT_LIBRARIES + "/" + u.Config.Library + endpoint + query + "&" + XPLEXTOKEN + u.Config.Token
}

func (u *PlexURLs) makeURL(route string) string {
	return u.withToken(u.Server() + route)
}

func (u *PlexURLs) withToken(url string) string {
	return url + "?" + XPLEXTOKEN + u.Config.Token
}
