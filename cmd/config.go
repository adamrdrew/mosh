package cmd

import (
	"fmt"
	"strconv"

	"github.com/adamrdrew/mosh/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCommand)
	configCommand.AddCommand(configCacheMaxSize)
	configCommand.AddCommand(configCacheMaxDays)
	configCommand.AddCommand(configShowArt)
}

var configCommand = &cobra.Command{
	Use:     "config",
	Aliases: []string{"conf"},
	Short:   "Manage the config",
	Long:    `Manage the config`,
}

var configCacheMaxSize = &cobra.Command{
	Use:     "cache-max-size",
	Short:   "Set the max cache size in MB",
	Aliases: []string{"max-size", "cache-size"},
	Long:    "Set the max cache size in MB",
	Run: func(cmd *cobra.Command, args []string) {
		errText := "Error: Invalid value. Enter a number between 256 and 16348."
		conf := config.GetConfig()
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(errText)
			return
		}
		if val < 256 || val > 16348 {
			fmt.Println(errText)
			return
		}
		conf.CacheMaxSizeMB = val
		conf.Save()
	},
}

var configCacheMaxDays = &cobra.Command{
	Use:     "cache-max-days",
	Short:   "Set the max cache age in days",
	Aliases: []string{"max-days", "cache-days"},
	Long:    "Set the max cache age in days",
	Run: func(cmd *cobra.Command, args []string) {
		errText := "Error: Invalid value. Enter a number between 1 and 30."
		conf := config.GetConfig()
		val, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(errText)
			return
		}
		if val < 1 || val > 30 {
			fmt.Println(errText)
			return
		}
		conf.CacheMaxAgeDays = val
		conf.Save()
	},
}

var configShowArt = &cobra.Command{
	Use:     "show-art [true|false]",
	Short:   "Enable or disable showing album art in get commands",
	Aliases: []string{"art"},
	Long:    "Enable or disable showing album art in get commands",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.GetConfig()
		val, err := strconv.ParseBool(args[0])
		if err != nil {
			fmt.Println("Error: Invalid value. Enter true or false.")
			return
		}
		conf.ShowArt = val
		conf.Save()
	},
}
