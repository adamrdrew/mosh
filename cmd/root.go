package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mosh",
	Short: "MOSH is a Plex music player for the shell",
	Long: `MOSH (music on shell) is a Plex music client for the shell written in Go. 
	It rocks, and so do you. 
	Made with aggression by Adam Drew at https://github.com/adamrdrew/mosh`,
}

func Execute() error {
	return rootCmd.Execute()
}
