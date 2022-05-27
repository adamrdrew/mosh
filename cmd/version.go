package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver", "v"},
	Short:   "Version number",
	Long:    `Display version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", config.Version)
	},
}
