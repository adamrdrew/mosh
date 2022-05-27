package filehandler

import (
	"os"
	"testing"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/server"
	"github.com/stretchr/testify/assert"
)

func setup() FileHandler {
	c := config.GetConfig()
	s := server.GetServer(&c)
	tr := responses.ResponseTrack{}
	tr.Type = "song"
	tr.Media.Part.Key = "library/parts/77708/1574343480/file.mp3"
	tr.Key = "/library/parts/77708/1574343480/file.mp3"
	tr.RatingKey = "3333"
	fh := GetFileHandler(s, tr)
	return fh
}

func TestGetFileHandler(t *testing.T) {
	fh := setup()
	assert.IsType(t, FileHandler{}, fh)
	assert.Equal(t, os.Getenv("MOSH_CACHE_DIR")+"/", fh.cacheDir)
}

/*
This is beyond my test fu
func TestGetTrackFile(t *testing.T) {
}
*/
