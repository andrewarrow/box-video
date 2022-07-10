package video

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

type Word struct {
	Word string
	Time int
}

var wordsTo = 35

func MakeWords(filename string) {
	RmRfBang()

	lines := ReadWordTimes(filename)
	fmt.Println(lines)
	wordsFromLines(lines[0:wordsTo])
}

func wordsFromLines(lines [][]Word) {

	for _, line := range lines {
		dc := gg.NewContext(1920, 1080)
		dc.SetRGB(0, 200, 200)
		dc.Clear()
		dc.LoadFontFace("arialbd.ttf", 96)

		for i, _ := range line {
			drawWordsWithColorOn(dc, i, line)
		}
	}

	dir := "nine"
	files, _ := ioutil.ReadDir(dir)
	for i, file := range files {
		if i > wordsTo {
			break
		}
		name := file.Name()
		fmt.Println(name)
		makeBoxFrame(i, dir, name)
	}
	ffmpeg("9")
}

func makeBoxFrame(i int, dir, name string) {
	existing, e := gg.LoadPNG(fmt.Sprintf("data/img%07d.png", i))
	if e != nil {
		return
	}
	dc := gg.NewContextForImage(existing)

	file, _ := os.Open(dir + "/" + name)
	im, _ := png.Decode(file)
	rgba := im.(*image.RGBA)

	x := 660
	y := 191
	w := 609 - 16
	h := 346 - 8
	cropped := rgba.SubImage(image.Rect(x, y, w+x, h+y))

	color := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
	pattern := gg.NewSolidPattern(color)
	dc.SetFillStyle(pattern)
	dc.DrawRectangle(650, 180, float64(w+20), float64(h+20))
	dc.Fill()
	dc.DrawImage(cropped, 0, 0)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", i))
}

func drawWordsWithColorOn(dc *gg.Context, index int, words []Word) {
	x := 200.0
	y := 900.0
	dc.SetRGB(1, 1, 1)
	for i, word := range words {
		w, _ := dc.MeasureString(word.Word)
		fmt.Println(w, word.Word, x)
		dc.SetRGB(0, 0, 0)
		dc.DrawString(word.Word, x, y)
		dc.SetRGB(1, 1, 1)
		if i == index {
			dc.SetRGB(255, 1, 1)
		}
		dc.DrawString(word.Word, x-3, y-3)
		x += w + 23
	}

	time := words[index].Time

	count := 0
	for {
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		count++
		if count > time {
			break
		}
	}
}

func ReadWordTimes(filename string) [][]Word {
	b, _ := ioutil.ReadFile(filename)

	wordLines := [][]Word{}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}

		tokens := strings.Split(trimmed, "|")
		allWords := tokens[0]
		allTimes := tokens[1]

		words := strings.Split(allWords, " ")
		times := strings.Split(allTimes, ",")

		if len(words) != len(times) {
			fmt.Println("!!!!")
			os.Exit(1)
			break
		}

		wordLine := []Word{}
		for i, word := range words {
			time, _ := strconv.Atoi(times[i])
			w := Word{word, time}
			wordLine = append(wordLine, w)
		}
		wordLines = append(wordLines, wordLine)
	}

	return wordLines
}
