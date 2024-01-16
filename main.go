package main

import (
	"flag"
	"fmt"
	"os"

	// "strings"

	"github.com/kkdai/youtube/v2"
	// "github.com/kkdai/youtube/v2"
)

const (
	Outfiletext = "-o              Output filepath. Do not specify a file extension. Default is $HOME/something.mp4"
	URLText     = "-u              URL of video to download"
	TypeText    = "-t              Accepts 'audio' or 'video' as input."
)

func main() {

	//Define flags
	videoOutfile := flag.String("o", "", "Output filepath. Do not specify a file extension. Default is $HOME/something.mp4")
	videoURL := flag.String("u", "https://www.youtube.com/watch?v=dQw4w9WgXcQ", "URL of video to download")
	videoType := flag.String("t", "video", "Accepts 'audio' or 'video' as input.")
	// help := flag.String("h", "", "Outputs a help dialog, showing how to use the program.")

	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("=========================HELP=========================\n")
		fmt.Println(Outfiletext + "\n\n" + URLText + "\n\n" + TypeText + "\n")
		fmt.Println("=========================HELP=========================")
		os.Exit(1234)
	}

	*videoOutfile = getDefaults(*videoOutfile)
	client := youtube.Client{}
	videoID := getVideoID(*videoURL)
	videoData := getVideoMetadata(videoID, client)

	videoOutfileMP4 := downloadVideo(&client, videoData, *videoType, *videoOutfile)

	if *videoType == "audio" {
		convertVideo(videoOutfileMP4)
	}

}
