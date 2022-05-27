package plex_urls

import (
	"testing"

	"github.com/adamrdrew/mosh/config"
	"github.com/stretchr/testify/assert"
)

var urls PlexURLs
var conf config.Config
var baseURL string

const query = "666"

func TestServer(t *testing.T) {
	got := urls.Server()
	assert.Equal(t, baseURL, got)
}

func TestGetLibraries(t *testing.T) {
	got := urls.GetLibraries()
	want := baseURL + RT_LIBRARIES + "?" + XPLEXTOKEN + conf.Token
	assert.Equal(t, want, got)
}

func TestGetChildren(t *testing.T) {
	got := urls.GetChildren(query)
	want := baseURL + RT_METADATA + urls.encodeQuery(query) + RT_CHILDREN + "?" + XPLEXTOKEN + conf.Token
	assert.Equal(t, want, got)
}

func TestSearchArtists(t *testing.T) {
	got := urls.SearchArstists(query)
	want := baseURL + RT_LIBRARIES + "/" + conf.Library + RT_ARTIST_SEARCH + urls.encodeQuery(query) + "&" + XPLEXTOKEN + urls.Config.Token
	assert.Equal(t, want, got)
}
func TestSearchAlbums(t *testing.T) {
	got := urls.SearchAlbums(query)
	want := baseURL + RT_LIBRARIES + "/" + conf.Library + RT_ALBUM_SEARCH + urls.encodeQuery(query) + "&" + XPLEXTOKEN + urls.Config.Token
	assert.Equal(t, want, got)
}
func TestMain(m *testing.M) {
	conf = config.GetConfig()
	urls = GetPlexURLs(&conf)
	baseURL = "http://" + conf.Address + ":" + conf.Port
}
