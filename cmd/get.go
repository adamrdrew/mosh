package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCommand)
	getCommand.AddCommand(getPlayingCommand)
	getCommand.AddCommand(getQueueCommand)
}

var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Get info ",
	Long:  `Get info on things like song, queue, etc`,
}

var getPlayingCommand = &cobra.Command{
	Use:   "playing",
	Short: "See what's playing",
	Long:  `Get detailed info on what is currently playing`,
	Run: func(cmd *cobra.Command, args []string) {
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "now-playing",
		})
		response := responseList.First()
		if response.Code == "OK" {
			printer.NowPlaying(response)
		} else {
			fmt.Println(response.Message)
		}

	},
}

var getQueueCommand = &cobra.Command{
	Use:   "queue",
	Short: "See the play queue",
	Long:  `Get detailed info on the play queue`,
	Run: func(cmd *cobra.Command, args []string) {
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "get-queue",
		})
		printer.PlayQueue(responseList)
	},
}