package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/kkdai/youtube/v2"
)

var (
	userOperatingSystem string
)

// used for printing out help menu, since I can't
// figure out how to make no options count as -h
const (
	Outfiletext = "-o              Output filepath. Do not specify a file extension. Default is $HOME/something.mp4"
	URLText     = "-u              URL of video to download"
	TypeText    = "-a              Use this flag if you want an (a)udio file, rather than a video."
)

func main() {

	//Define flags
	//I can't figure out "help" yet
	videoOutfile := flag.String("o", "", "Output filepath. Do not specify a file extension. Default is $HOME/something.mp4")
	videoURL := flag.String("u", "https://www.youtube.com/watch?v=dQw4w9WgXcQ", "URL of video to download")
	videoType := flag.Bool("a", false, "Use this flag if you want an (a)udio file instead.")
	// help := flag.String("h", "", "Outputs a help dialog, showing how to use the program.")

	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Println("=========================HELP=========================\n")
		fmt.Println(Outfiletext + "\n\n" + URLText + "\n\n" + TypeText + "\n")
		fmt.Println("=========================HELP=========================")
		os.Exit(1234) //There is no reason for this number
	}

	//Determine user operating system
	userOperatingSystem = runtime.GOOS

	*videoOutfile = getDefaults(*videoOutfile, userOperatingSystem)
	client := youtube.Client{}
	videoID := getVideoID(*videoURL)
	videoData := getVideoMetadata(videoID, client)

	videoOutfileMP4 := downloadVideo(&client, videoData, *videoType, *videoOutfile)

	if *videoType == true {
		convertVideo(videoOutfileMP4)
	}

}
