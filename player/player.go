package player

import "github.com/adamrdrew/mosh/responses"

/*
For our queue we've got an array of ResponseTrack instances.
We can get part of the MP3 URL from ResponseTrack.GetPath().
However, that's only the route - not the server URL, port, etc.
Server knows all that stuff, so something needs to work with both
Server and ResponseTrack. Additionally, we want the option for
caching, so whatever does the work of determining the location
of a music file should probably be where cache happens too.
*/

func New(queue []responses.ResponseTrack) Player {
	p := Player{
		queue:           queue,
		Playing:         false,
		NowPlayingIndex: 0,
	}
	return p
}

type Player struct {
	queue           []responses.ResponseTrack
	Playing         bool
	NowPlayingIndex int
}

func (p *Player) maxIndex() int {
	return len(p.queue) - 1
}

func (p *Player) Play() {
	if p.Playing {
		return
	}
	p.Playing = true

}

func (p *Player) Pause() {
	if !p.Playing {
		return
	}
	p.Playing = false

}

func (p *Player) Stop() {

}

func (p *Player) Back() {
	if p.NowPlayingIndex == 0 {
		return
	}
}

func (p *Player) Forward() {
	if p.NowPlayingIndex == p.maxIndex() {
		return
	}
}
