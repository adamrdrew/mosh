package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/ipc"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(controlCommand)
	controlCommand.AddCommand(controlStopCommand)
	controlCommand.AddCommand(controlPlayCommand)
	controlCommand.AddCommand(controlNextCommand)
	controlCommand.AddCommand(controlBackCommand)
}

var controlCommand = &cobra.Command{
	Use:     "control",
	Aliases: []string{"c", "co", "con", "ctrl"},
	Short:   "Control playback",
	Long:    `These are the buttons on your tape deck (are you old enough to remember tape decks?)`,
}

var controlStopCommand = &cobra.Command{
	Use:     "stop",
	Short:   "Stop the play queue",
	Aliases: []string{"s", "st"},
	Long:    `Stops the play queue but leaves it loaded. Returns the play position to the first track.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
		if !guardSetupCompleted() {
			return
		}
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "stop",
		})
		response := responseList.First()
		fmt.Println(response.Message)
	},
}

var controlPlayCommand = &cobra.Command{
	Use:     "play",
	Aliases: []string{"p", "pl"},
	Short:   "Play the queue from the current position",
	Long:    `Starts playing the play queue from the current position`,
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
		if !guardSetupCompleted() {
			return
		}
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "play-queue",
		})
		response := responseList.First()
		fmt.Println(response.Message)
	},
}

var controlNextCommand = &cobra.Command{
	Use:     "next",
	Aliases: []string{"n", "ne"},
	Short:   "Skip to the next track in queue",
	Long:    `Skip to the next track in queue`,
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
		if !guardSetupCompleted() {
			return
		}
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "next",
		})
		response := responseList.First()
		fmt.Println(response.Message)
	},
}

var controlBackCommand = &cobra.Command{
	Use:     "back",
	Aliases: []string{"b", "ba"},
	Short:   "Skip to the previous track in queue",
	Long:    `Skip to the previous track in queue`,
	Run: func(cmd *cobra.Command, args []string) {
		if !guardDaemonIsRunning() {
			return
		}
		if !guardSetupCompleted() {
			return
		}
		responseList := ipc.SendMessageToDaemon(ipc.Message{
			Command: "back",
		})
		response := responseList.First()
		fmt.Println(response.Message)
	},
}
