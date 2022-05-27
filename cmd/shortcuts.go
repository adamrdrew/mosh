package cmd

import (
	"fmt"

	"github.com/adamrdrew/mosh/printer"
	"github.com/adamrdrew/mosh/shortcuts"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shortcutCommand)
	shortcutCommand.AddCommand(shortcutAddCommand)
	shortcutCommand.AddCommand(shortcutDelCommand)
	shortcutCommand.AddCommand(shortcutListCommand)

}

var shortcutCommand = &cobra.Command{
	Use:     "shortcuts",
	Aliases: []string{"sh", "shortcut"},
	Short:   "Manage shortcuts",
	Long:    `Manage shortcuts for frequently accessed IDs`,
}

var shortcutAddCommand = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a shortcut for an ID",
	Long:    `Example: shortcut add nin 3220`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Provide the shortcut and the ID")
			return
		}
		shortcuts.Add(args[0], args[1])
	},
}

var shortcutDelCommand = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del"},
	Short:   "Remove a shortcut",
	Long:    `Example: shortcut delete nin`,
	Run: func(cmd *cobra.Command, args []string) {
		shortcuts.Delete(args[0])
	},
}

var shortcutListCommand = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all shortcuts",
	Long:    `List all shortcuts in a table`,
	Run: func(cmd *cobra.Command, args []string) {
		cuts := shortcuts.GetAll()
		printer.Shortcuts(cuts)
	},
}
