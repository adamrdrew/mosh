package cmd

import (
	"github.com/adamrdrew/mosh/ipc"
	"github.com/adamrdrew/mosh/printer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCommand)
}

var statusCommand = &cobra.Command{
	Use:   "status",
	Short: "See what's playing",
	Long:  `Get detailed info on what is currently playing`,
	Run: func(cmd *cobra.Command, args []string) {
		response := ipc.SendMessageToDaemon(ipc.Message{
			Command: "now-playing",
		})
		printer.NowPlaying(response)
	},
}
