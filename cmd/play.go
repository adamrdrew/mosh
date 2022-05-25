package cmd

import (
	"fmt"
	"strings"

	"github.com/adamrdrew/mosh/filehandler"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCommand)
	playCommand.AddCommand(playAlbumCommand)
	playCommand.AddCommand(playSongCommand)
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
		songs := server.GetSongsForAlbum(strings.Join(args, " "))
		fileHandler := filehandler.GetFileHandler(server, songs[0])
		songPath := fileHandler.GetTrackFile()
		PlaySong(songPath)

	},
}

var playSongCommand = &cobra.Command{
	Use:   "song [song ID]",
	Short: "Play song",
	Long:  `Play a song by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		//server := getServer()
		fmt.Println("Play song")
	},
}
