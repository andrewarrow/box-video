package main

import (
	"box-video/audio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "source" {
		file := os.Args[2]
		fmt.Println(file)
	} else if command == "audio" {
		audio.PlayTest()
	} else if command == "mark" {
		file := os.Args[2]
		audio.PlayForMark(file)
	} else if command == "merge" {
		audio.Merge()
	}
}
