package video

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/fogleman/gg"
)

var text = `So then you see the deeper layer
of reality of oh, everyone is acting
perfectly according to their state
of consciousness. And that's why
reality is perfect.
`

type Word struct {
	Word         string
	Milliseconds int
}

func MakeWords() {
	RmRfBang()

	w := Word{"So", 250}
	line := []Word{w}
	w = Word{"then", 250}
	line = append(line, w)
	w = Word{"you", 250}
	line = append(line, w)
	w = Word{"see", 250}
	line = append(line, w)
	w = Word{"the", 125}
	line = append(line, w)
	w = Word{"deeper", 250}
	line = append(line, w)
	w = Word{"layer", 250}
	line = append(line, w)

	lines := [][]Word{line}

	w = Word{"of", 250}
	line = []Word{w}
	w = Word{"reality", 250}
	line = append(line, w)
	w = Word{"of", 125}
	line = append(line, w)
	w = Word{"oh,", 250}
	line = append(line, w)
	w = Word{"everyone", 250}
	line = append(line, w)
	w = Word{"is", 250}
	line = append(line, w)
	w = Word{"acting", 250}
	line = append(line, w)

	lines = append(lines, line)

	w = Word{"perfectly", 250}
	line = []Word{w}
	w = Word{"according", 250}
	line = append(line, w)
	w = Word{"to", 250}
	line = append(line, w)
	w = Word{"their", 250}
	line = append(line, w)
	w = Word{"state", 250}
	line = append(line, w)
	lines = append(lines, line)

	w = Word{"of", 125}
	line = []Word{w}
	w = Word{"consciousness.", 500}
	line = append(line, w)
	w = Word{"And", 250}
	line = append(line, w)
	w = Word{"that's", 250}
	line = append(line, w)
	w = Word{"why", 250}
	line = append(line, w)
	lines = append(lines, line)

	w = Word{"reality", 500}
	line = []Word{w}
	w = Word{"is", 250}
	line = append(line, w)
	w = Word{"perfect.", 250}
	line = append(line, w)
	lines = append(lines, line)

	fmt.Println(lines)
	wordsFromLines(lines)
}

func makeBoxFrame(i int, dir, name string) {
	existing, _ := gg.LoadPNG(fmt.Sprintf("data/img%07d.png", i))
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
		name := file.Name()
		fmt.Println(name)
		makeBoxFrame(i, dir, name)
	}
	ffmpeg("9")
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

	ms := words[index].Milliseconds
	breakAt := 1
	if ms == 250 {
		breakAt = 2
	} else if ms == 375 {
		breakAt = 3
	} else if ms == 500 {
		breakAt = 4
	} else if ms == 625 {
		breakAt = 5
	} else if ms == 750 {
		breakAt = 6
	} else if ms == 875 {
		breakAt = 7
	} else if ms == 1000 {
		breakAt = 8
	} else if ms == 1125 {
		breakAt = 9
	}

	count := 0
	for {
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		count++
		if count > breakAt {
			break
		}
	}
}
