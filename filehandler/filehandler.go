package filehandler

import (
	"os"
	"strings"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/server"
	"github.com/cavaliergopher/grab/v3"
)

func GetFileHandler(server server.Server, track responses.ResponseTrack) FileHandler {
	fh := FileHandler{
		server:   server,
		track:    track,
		cacheDir: config.GetCacheDir(),
	}
	fh.createCacheDir()
	return fh
}

//File handler acts as the abstraction between the rest of the app and getting paths to files
//that we can play.
type FileHandler struct {
	server   server.Server
	track    responses.ResponseTrack
	cacheDir string
}

//Public API. Returns the path to a local file for the song in question
//Encapsulates caching, downloading, etc
func (f *FileHandler) GetTrackFile() string {
	f.warmCache()
	return f.getCachePath()
}

//Checks to see if the cached file exists
//If it doesn't we download it
func (f *FileHandler) warmCache() {
	path := f.getCachePath()
	_, statErr := os.Stat(path)
	if !os.IsNotExist(statErr) {
		//Found cache, bail
		return
	}
	//No cache found, so cache it baby
	_, dlErr := grab.Get(path, f.getTrackURL())
	if dlErr != nil {
		panic(dlErr)
	}
}

func (f *FileHandler) getPath() string {
	return f.track.GetPath()
}

//For a server and track get the full file URL
func (f *FileHandler) getTrackURL() string {
	return f.server.MakeURL(f.getPath())
}

//Our cached file's name is just the track path with /'s converted to -'s
//and the first char removed
//i.e. /library/parts/77708/1574343480/file.mp3 == library-parts-77708-1574343480-file.mp3
func (f *FileHandler) getCacheFileName() string {
	return strings.Replace(f.getPath(), "/", "-", -1)[1:]
}

//This is the full path to the file
//i.e.  /tmp/mosh/library-parts-77708-1574343480-file.mp3
func (f *FileHandler) getCachePath() string {
	return f.cacheDir + f.getCacheFileName()
}

//If the cache dir doesn't exist we create it
//TODO: should probably be moved to setup or something
func (f *FileHandler) createCacheDir() {
	_, statErr := os.Stat(f.cacheDir)
	if !os.IsNotExist(statErr) {
		return
	}
	mkErr := os.Mkdir(f.cacheDir, 0755)
	if mkErr == nil {
		return
	}
	panic(mkErr)
}
