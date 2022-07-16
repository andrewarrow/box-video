package main

import (
	"box-video/audio"
	"box-video/video"
	"fmt"
	"io/ioutil"
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
	} else if command == "clip" {
		file := os.Args[2]
		words := os.Args[3]
		audio.PlayForClip(file, words)
	} else if command == "stars" {
		video.MakeStars()
	} else if command == "words" {
		file := os.Args[2]
		video.MakeWords(file)
	} else if command == "sound" {
		v := os.Args[2]
		a := os.Args[3]
		video.AddSound(v, a)
	} else if command == "eight" {
		video.MakeEight()
	} else if command == "frames" {
		file := os.Args[2]
		name := os.Args[3]
		video.MakeFrames(file, name)
	} else if command == "title" {
		title := os.Args[2]
		dir := "/Users/aa/watts/joy/keep"
		files, _ := ioutil.ReadDir(dir)
		for i, file := range files {
			name := file.Name()
			path := dir + "/" + name
			video.MakeTitle(title, path, i*30)
		}
	}
}
