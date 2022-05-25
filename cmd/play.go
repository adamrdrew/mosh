package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/adamrdrew/mosh/filehandler"
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/vorbis"
	"github.com/gabriel-vasile/mimetype"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCommand)
	playCommand.AddCommand(playAlbumCommand)
	playCommand.AddCommand(playSongCommand)
}

var playCommand = &cobra.Command{
	Use:   "play [album | song]",
	Short: "Play a song or album",
	Long:  `Play a song or album`,
}

var playAlbumCommand = &cobra.Command{
	Use:   "album [album ID]",
	Short: "Play an album",
	Long:  `Play an album by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		server := getServer()
		songs := server.GetSongsForAlbum(strings.Join(args, " "))
		fileHandler := filehandler.GetFileHandler(server, songs[0])
		songPath := fileHandler.GetTrackFile()
		PlaySong(songPath)

	},
}

func PlaySong(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	mtype, err := mimetype.DetectFile(path)

	var streamer beep.StreamSeekCloser
	var format beep.Format
	var decErr error

	switch mtype.String() {
	case "audio/flac":
		streamer, format, decErr = flac.Decode(f)
	case "audio/mpeg":
		streamer, format, decErr = mp3.Decode(f)
	case "audio/ogg":
		streamer, format, decErr = vorbis.Decode(f)
	default:
		return
	}

	if decErr != nil {
		panic(decErr)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {}
}

var playSongCommand = &cobra.Command{
	Use:   "song [song ID]",
	Short: "Play song",
	Long:  `Play a song by ID`,
	Run: func(cmd *cobra.Command, args []string) {
		//server := getServer()
		fmt.Println("Play song")
	},
}
