package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/moshd"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(daemonCommand)
	daemonCommand.AddCommand(daemonStartCommand)
	daemonCommand.AddCommand(daemonKillCommand)
	daemonCommand.AddCommand(daemonStatusCommand)
}

var daemonCommand = &cobra.Command{
	Use:     "daemon",
	Aliases: []string{"d", "da"},
	Short:   "Manage the daemon",
	Long:    `Manage the mosh daemon`,
}

var daemonStartCommand = &cobra.Command{
	Use:     "start",
	Short:   "Start the mosh daemon",
	Aliases: []string{"s", "st"},
	Long:    `Starts the mosh daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		moshd.StartWaitAndCheck()
	},
}

var daemonKillCommand = &cobra.Command{
	Use:     "kill",
	Short:   "Kill the mosh daemon",
	Aliases: []string{"k", "ki", "stop"},
	Long:    `Kill the mosh daemon. Probably the most metal command.`,
	Run: func(cmd *cobra.Command, args []string) {
		moshd.KillDaemon()
	},
}

var daemonStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "Get status of mosh daemon",
	Long:  `Get the status of the mosh daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		proc := moshd.CheckDaemonStatus()
		if proc != nil {
			fmt.Println("Daemon status OK - PID:", proc.Pid)
		}
	},
}
