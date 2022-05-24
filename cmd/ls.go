package cmd

import (
	"github.com/adamrdrew/mosh/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCommand)
	lsCommand.AddCommand(lsArtistCommand)
	lsCommand.AddCommand(lsAlbumCommand)
}

var lsCommand = &cobra.Command{
	Use:   "ls [album ID|artist ID]",
	Short: "List children",
	Long:  `List children of an artist or album, being albums or songs respectively`,
}

var lsArtistCommand = &cobra.Command{
	Use:   "artist [artist ID]",
	Short: "List albums for an artist",
	Long:  `List albums by artist ID. Use the search artist command to get the ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := args[0]
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
		query := args[0]
		songs := server.GetSongsForAlbum(query)
		printer.Songs(songs)
	},
}
