package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
)

// This gets reasonable defaults if the user doesn't enter anything
var (
	specialchar string
	unixtime    string
)

func getDefaults(videoOutFile string, userOperatingSystem string) string {
	if strings.Contains(userOperatingSystem, "win") {
		specialchar = "\\Videos\\"
		unixtime = "defaultvideoname"
	} else {
		specialchar = "/"
		layout := "Jan-2-15:04:05"
		unixtime = time.Now().Format(layout)
	}

	if videoOutFile == "" {
		fmt.Println("No video output file specified. Generating default...")
		pwd, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Failed to get current working directory: ")
			panic(err)
		}
		videoOutFile = pwd + specialchar + unixtime

	}
	fmt.Printf("Default path generated: %s\n", videoOutFile)
	return videoOutFile

}

// Gets the video ID, removes the youtube.com/watch?v= before it
func getVideoID(videoURL string) string {
	fmt.Println("Extracting Video ID (Removing everything before watch?v=)")
	videoURL, err := youtube.ExtractVideoID(videoURL)
	if err != nil {
		fmt.Println("Failed to extract Video ID, please see below:")
		panic(err)
	}
	fmt.Println("Video ID extracted successfully!")
	return videoURL

}

func getVideoMetadata(videoID string, client youtube.Client) (videoData *youtube.Video) {
	fmt.Println("Getting video metadata (Name, Author, Description, etc)")
	//Gets video metadata, like name/author/description/views/etc
	videoData, err := client.GetVideo(videoID)
	if err != nil {
		fmt.Println("Failed to retrieve metadata, see below:")
		panic(err)
	}
	fmt.Println("Metadata retrieved succesfully")
	return
}

func downloadVideo(client *youtube.Client, videoData *youtube.Video, videoType bool, videoOutFile string) string {
	//Checks if videoOutFile ends in mp4, adds mp4 if not
	index := strings.LastIndex(videoOutFile, ".")
	if index != -1 {
		videoOutFile = videoOutFile[:index] //removes file extensions since I'm too lazy to check for them
	}
	videoOutFile += ".mp4"

	//Format and the accompanying if/else statement are used to determine whether to download video as audio or video format
	var format *youtube.Format
	formatlist := videoData.Formats.WithAudioChannels()
	if videoType == true {
		format = &formatlist[3]
	} else {
		format = &formatlist[0]
	}

	//Getting the actual data stream, still needs to be written to a file
	stream, _, err := client.GetStream(videoData, format)
	if err != nil {
		fmt.Println("Failed to get video stream, please see below:")
		panic(err)
	}
	defer stream.Close()

	//Creating the output .mp4 to copy data stream to (creates actualy file contents of mp4/mp3)
	fileptr, err := os.Create(videoOutFile)
	if err != nil {
		fmt.Println("Failed to create video file, please see below:")
	}
	defer fileptr.Close()
	fmt.Println("Video file created")

	io.Copy(fileptr, stream)
	fmt.Printf("Your file was downloaded to: %s\n", videoOutFile)

	return videoOutFile
}

// mp4 to mp3 conversion
// TODO: Remove old mp4 after converting
func convertVideo(videoOutFile string) {
	file, err := os.Open(videoOutFile)
	if err != nil {
		fmt.Println(videoOutFile, "does not exist. Creating..")
		_, err = os.Create(videoOutFile)
		if err != nil {
			fmt.Println("Error creating file: ")
			panic(err)
		}
	}
	defer os.Remove(videoOutFile)
	defer file.Close()

	//handle if user input specifies mp4 or not
	newvideoOutFile := strings.Replace(videoOutFile, ".mp4", ".mp3", -1)

	conversion := exec.Command("ffmpeg", "-i", videoOutFile, newvideoOutFile)
	conversion.Stdout = os.Stdout //this should make ffmpeg's output visible to user????
	err = conversion.Run()
	if err != nil {
		fmt.Println("Failed to run conversion command. You may not have ffmpeg installed on your machine.")
		panic(err)
	}

	fmt.Println("Your conversion is done, your new file is located at: \n" + newvideoOutFile + "\nEnjoy! :DD")
}
