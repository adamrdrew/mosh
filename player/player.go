package player

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/filehandler"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/server"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gabriel-vasile/mimetype"
)

type Player struct {
	Streamer     beep.StreamSeekCloser
	Queue        []responses.ResponseTrack
	CurrentIndex int
	MaxIndex     int
	Server       server.Server
	Config       config.Config
	StopPlayLoop bool
}

func (p *Player) Init() {
	p.Config = config.GetConfig()
	p.Server = server.GetServer(&p.Config)
	p.CurrentIndex = 0
	p.MaxIndex = 0
	p.StopPlayLoop = false
}

func (p *Player) GetPlayQueue() ipc.Response {
	response := ipc.Response{}

	if len(p.Queue) == 0 {
		item := ipc.ResponseItem{
			Song:        "",
			Album:       "",
			Artist:      "",
			TotalTime:   "0",
			CurrentTime: "0",
			Message:     "Nothing in queue. I feel empty.",
			Code:        "EMPTY",
		}
		response.Add(item)
		return response
	}

	for i, track := range p.Queue {
		nowPlaying := "NOTPLAYING"
		if i == p.CurrentIndex {
			nowPlaying = "PLAYING"
		}
		item := ipc.ResponseItem{
			Song:    track.Title,
			Album:   track.ParentTitle,
			Artist:  track.GrandParentTitle,
			Message: strconv.FormatInt(int64(i+1), 10),
			Code:    nowPlaying,
		}
		response.Add(item)
	}

	return response
}

func (p *Player) GetNowPlaying() ipc.ResponseItem {
	if len(p.Queue) == 0 {
		return ipc.ResponseItem{
			Song:        "",
			Album:       "",
			Artist:      "",
			TotalTime:   "0",
			CurrentTime: "0",
			Message:     "Nothing playing. Kinda quiet in here.",
			Code:        "EMPTY",
		}
	}
	track := p.Queue[p.CurrentIndex]
	return ipc.ResponseItem{
		Song:        track.Title,
		Album:       track.ParentTitle,
		Artist:      track.GrandParentTitle,
		TotalTime:   strconv.FormatInt(int64(p.Streamer.Len()), 10),
		CurrentTime: strconv.FormatInt(int64(p.Streamer.Position()), 10),
		Message:     "",
		Code:        "OK",
	}
}

func (p *Player) SetQueue(queue []responses.ResponseTrack) {
	p.Queue = queue
	p.CurrentIndex = 0
	p.MaxIndex = len(queue) - 1
}

func (p *Player) QueueAlbum(albumID string) ipc.ResponseItem {
	response := ipc.ResponseItem{}
	response.Code = "OK"
	songs := p.Server.GetSongsForAlbum(albumID)
	if len(songs) == 0 {
		response.Code = "NOTFOUND"
		response.Message = "Album not found"
		return response
	}
	p.SetQueue(songs)
	response.Message = songs[0].ParentTitle + " by " + songs[0].GrandParentTitle + " is now playing."
	return response
}

func (p *Player) NowPlayingSongString() string {
	song := p.Queue[p.CurrentIndex]
	return song.Title + " by " + song.GrandParentTitle + " from the album " + song.ParentTitle
}

func (p *Player) PlayQueue() {
	for i := p.CurrentIndex; i < p.MaxIndex; i++ {
		if p.StopPlayLoop {
			p.StopPlayLoop = false
			break
		}
		p.CurrentIndex = i
		song := p.Queue[i]
		fileHandler := filehandler.GetFileHandler(p.Server, song)
		path := fileHandler.GetTrackFile()
		p.PlaySongFile(path)
	}
}

func (p *Player) StopQueue() ipc.ResponseItem {
	p.StopPlayLoop = true
	p.CurrentIndex = 0
	p.Streamer.Close()
	p.Streamer = nil
	return ipc.ResponseItem{
		Status:  "OK",
		Message: "Playback stopped.",
	}
}

func (p *Player) GoBackInQueue() ipc.ResponseItem {
	if len(p.Queue) == 0 {
		return ipc.ResponseItem{
			Status:  "NOPE",
			Message: "Nothing in queue",
		}
	}
	currentIndex := p.CurrentIndex
	log.Println("GoBackInQueue currentIndex", currentIndex)
	if currentIndex == 0 {
		return ipc.ResponseItem{
			Status:  "NOPE",
			Message: "Already at first track",
		}
	}
	p.StopPlayLoop = true
	p.Streamer.Close()
	p.Streamer = nil
	if !p.waitForPlayThreadToDie() {
		return ipc.ResponseItem{
			Status:  "ERROR",
			Message: "Stopping the play queue thread failed.",
		}
	}

	p.CurrentIndex = currentIndex - 1
	log.Println("GoBackInQueue p.CurrentIndex", p.CurrentIndex)
	go p.PlayQueue()
	return ipc.ResponseItem{
		Status:  "OK",
		Message: "Went back. Next up: " + p.NowPlayingSongString(),
	}
}

func (p *Player) GoForwardInQueue() ipc.ResponseItem {
	if len(p.Queue) == 0 {
		return ipc.ResponseItem{
			Status:  "NOPE",
			Message: "Nothing in queue",
		}
	}
	currentIndex := p.CurrentIndex
	if currentIndex == p.MaxIndex {
		return ipc.ResponseItem{
			Status:  "NOPE",
			Message: "Already at last track",
		}
	}
	p.StopPlayLoop = true
	p.Streamer.Close()
	p.Streamer = nil
	if !p.waitForPlayThreadToDie() {
		return ipc.ResponseItem{
			Status:  "ERROR",
			Message: "Stopping the play queue thread failed.",
		}
	}
	p.CurrentIndex = currentIndex + 1
	go p.PlayQueue()
	return ipc.ResponseItem{
		Status:  "OK",
		Message: "Went forward. Next up: " + p.NowPlayingSongString(),
	}
}

func (p *Player) waitForPlayThreadToDie() bool {
	retVal := false
	start := time.Now()
	for {
		if p.StopPlayLoop == false {
			retVal = true
			break
		}
		if time.Since(start) >= 30*time.Second {
			retVal = false
			break
		}
	}
	return retVal
}

func (p *Player) PlaySongFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	mtype, err := mimetype.DetectFile(path)

	var decErr error
	var format beep.Format
	switch mtype.String() {
	case "audio/flac":
		p.Streamer, format, decErr = flac.Decode(file)
	case "audio/mpeg":
		p.Streamer, format, decErr = mp3.Decode(file)
	case "audio/ogg":
		p.Streamer, format, decErr = vorbis.Decode(file)
	default:
		return
	}

	if decErr != nil {
		panic(decErr)
	}
	defer p.Streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	log.Print("Playing: ", path)
	done := make(chan bool)
	speaker.Play(beep.Seq(p.Streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}
