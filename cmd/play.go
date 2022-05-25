package cmd

import (
	"strings"

	"github.com/adamrdrew/mosh/printer"
	"github.com/spf13/cobra"
)

func init() {
	//rootCmd.AddCommand(playCommand)
	//searchCommand.AddCommand(playAlbumCommand)
	//searchCommand.AddCommand(playSongCommand)
}

var playCommand = &cobra.Command{
	Use:   "play [album | song]",
	Short: "Play a song or album",
	Long:  `Play a song or album`,
}

var playAlbumCommand = &cobra.Command{
	Use:   "album [album ID]",
	Short: "Play an album",
	Long:  `Play an album by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := strings.Join(args, " ")
		artists := server.SearchArtists(query)
		printer.Artists(artists)
	},
}

var playSongCommand = &cobra.Command{
	Use:   "song [song ID]",
	Short: "Play song",
	Long:  `Play a song by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := strings.Join(args, " ")
		albums := server.SearchAlbums(query)
		printer.Albums(albums)
	},
}
