package cmd

import (
	"github.com/adamrdrew/mosh/printer"
	"github.com/adamrdrew/mosh/shortcuts"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCommand)
	lsCommand.AddCommand(lsArtistCommand)
	lsCommand.AddCommand(lsAlbumCommand)
}

var lsCommand = &cobra.Command{
	Use:   "ls [album ID|artist ID]",
	Short: "List children. Shortcuts supported.",
	Long:  `List children of an artist or album, being albums or songs respectively. Supports shortcuts.`,
}

var lsArtistCommand = &cobra.Command{
	Use:   "artist [artist ID]",
	Short: "List albums for an artist",
	Long:  `List albums by artist ID. Use the search artist command to get the ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := shortcuts.Resolve(args[0])
		albums := server.GetAlbumsForArtist(query)
		printer.Albums(albums)
	},
}

var lsAlbumCommand = &cobra.Command{
	Use:   "album [album ID]",
	Short: "List songs for an album",
	Long:  `List songs by album ID. Use the search album command to get the ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := shortcuts.Resolve(args[0])
		songs := server.GetSongsForAlbum(query)
		printer.Songs(songs)
	},
}
