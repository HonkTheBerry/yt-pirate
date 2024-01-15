package main

import "fmt"

func displayHelp() {
	fmt.Println("Usage: yt-pirate [options]")
	fmt.Println("Options: ")
	fmt.Println(" -h    Display help information")
	fmt.Println(" -o    Output. The file you want to save as. Do NOT type file extension")
	fmt.Println(" -t    Filetype to be received. Acceptable arguments are 'audio' or 'video'.")
	fmt.Println(" -u    Video URL. Just paste the whole URL.")
}
