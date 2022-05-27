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
	Use:     "play",
	Short:   "Play a song or album",
	Long:    `Play a song or album`,
	Aliases: []string{"p", "pl"},
}

var playAlbumCommand = &cobra.Command{
	Use:     "album",
	Short:   "Play an album",
	Long:    `Play an album by ID`,
	Aliases: []string{"a", "al"},
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
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
	Use:     "song",
	Short:   "Play song",
	Long:    `Play a song by ID`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
		//server := getServer()
		fmt.Println("Play song")
	},
}
