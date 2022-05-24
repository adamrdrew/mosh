package printer

import (
	"os"

	"github.com/adamrdrew/mosh/responses"
	"github.com/jedib0t/go-pretty/v6/table"
)

const TABS = "\t\t\t"

func Artists(source []responses.ResponseArtistDirectory) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title"})
	for _, artist := range source {
		t.AppendRow(table.Row{artist.RatingKey, artist.Title})
	}
	t.Render()
}

func Albums(source []responses.ResponseAlbumDirectory) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title", "Artist"})
	for _, album := range source {
		t.AppendRow(table.Row{album.RatingKey, album.Title, album.ParentTitle})
	}
	t.Render()
}
