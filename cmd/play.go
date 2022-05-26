package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/shortcuts"
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
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "queue-album",
			Data:    shortcuts.Resolve(args[0]),
		})
		response := responseList.First()
		fmt.Println(response.Message)
		if response.Code == "OK" {
			ipc.SendMessageToDaemon(ipc.Message{
				Command: "play-queue",
			})
		}
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
