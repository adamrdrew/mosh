package player

import "github.com/adamrdrew/mosh/responses"

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

}

func (p *Player) Forward() {

}
