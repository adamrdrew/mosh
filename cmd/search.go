package cmd

import (
	"strings"

	"github.com/adamrdrew/mosh/config"
	"github.com/adamrdrew/mosh/printer"
	"github.com/adamrdrew/mosh/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(searchCommand)
	searchCommand.AddCommand(artistCommand)
	searchCommand.AddCommand(albumCommand)
}

var searchCommand = &cobra.Command{
	Use:     "search",
	Aliases: []string{"s", "se"},
	Short:   "Search the for an album or artist",
	Long:    `Search the library for albums or artists`,
}

var artistCommand = &cobra.Command{
	Use:     "artist",
	Aliases: []string{"ar"},
	Short:   "Search for an artist",
	Long:    `Search for artist by name. Partial names are fine.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := strings.Join(args, " ")
		artists := server.SearchArtists(query)
		printer.Artists(artists)
	},
}

var albumCommand = &cobra.Command{
	Use:     "album",
	Aliases: []string{"al"},
	Short:   "Search the library",
	Long:    `Search for an album by name`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		query := strings.Join(args, " ")
		albums := server.SearchAlbums(query)
		printer.Albums(albums)
	},
}

func getServer() server.Server {
	conf := config.GetConfig()
	return server.GetServer(&conf)
}
