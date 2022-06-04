package printer

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/responses"
	"github.com/adamrdrew/mosh/shortcuts"
	"github.com/adamrdrew/mosh/track_ascii"
	"github.com/jedib0t/go-pretty/v6/table"
)

const TABS = "\t\t\t"

func shortcutOrBlank(id string) string {
	shortcut := shortcuts.ReverseResolve(id)
	if shortcut == id {
		shortcut = ""
	}
	return shortcut
}

func Artists(source []responses.ResponseArtistDirectory) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Shortcut"})
	for _, artist := range source {
		shortcut := shortcutOrBlank(artist.RatingKey)
		t.AppendRow(table.Row{artist.RatingKey, artist.Title, shortcut})
	}
	t.Render()
}

func Albums(source []responses.ResponseAlbumDirectory) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Artist", "Shortcut"})
	for _, album := range source {
		shortcut := shortcutOrBlank(album.RatingKey)
		t.AppendRow(table.Row{album.RatingKey, album.Title, album.ParentTitle, shortcut})
	}
	t.Render()
}

func Shortcuts(cuts map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Shortcut", "ID"})
	for key, val := range cuts {
		t.AppendRow(table.Row{key, val})
	}
	t.Render()
}

func Songs(source []responses.ResponseTrack) {
	showAlbumArt(source[0].Image)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Track", "Title", "Album", "Artist"})
	for _, track := range source {
		t.AppendRow(table.Row{track.RatingKey, track.Index, track.Title, track.ParentTitle, track.GrandParentTitle})
	}
	t.Render()
}

func PlayQueue(source ipc.Response) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Playing", "Track", "Title", "Album", "Artist"})
	for _, track := range source.Responses {
		nowPlaying := ""
		if track.Code == "PLAYING" {
			nowPlaying = "X"
		} else {
			nowPlaying = ""
		}
		t.AppendRow(table.Row{nowPlaying, track.Message, track.Song, track.Album, track.Artist})
	}
	t.Render()
}

func showAlbumArt(source string) {
	conf := config.GetConfig()
	if conf.ShowArt {
		ascii := track_ascii.MakeConverter(source)
		fmt.Print(ascii.GetAscii())
	}
}

func NowPlaying(source ipc.ResponseItem) {
	showAlbumArt(source.Image)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Track", "Artist", "Album"})
	t.AppendRow(table.Row{source.Song, source.Artist, source.Album})
	t.Render()
	current, _ := strconv.ParseInt(source.CurrentTime, 10, 64)
	max, _ := strconv.ParseInt(source.TotalTime, 10, 64)
	RenderStatusBar(current, max)
}

func RenderStatusBar(val int64, maxVal int64) {
	pct := math.Round((float64(val) / float64(maxVal)) * 100.00)
	reduced := pct / 5
	outArr := []string{"["}
	max := 20
	for i := 0; i < int(reduced); i++ {
		outArr = append(outArr, "#")
	}
	for i := 0; i < max-int(reduced); i++ {
		outArr = append(outArr, "-")
	}
	outArr = append(outArr, "]")
	outStr := strings.Join(outArr, "")
	timePlayed := (float64(val) / 44100.00) / 60
	totalTime := (float64(maxVal) / 44100.00) / 60
	timePlayedString := strings.Replace(fmt.Sprintf("%.2f", timePlayed), ".", ":", 1)
	totalTimeString := strings.Replace(fmt.Sprintf("%.2f", totalTime), ".", ":", 1)
	fmt.Println(timePlayedString, "/", totalTimeString, outStr, pct, "%")
}
