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
	MilliSeconds int
}

func MakeWords() {
	RmRfBang()

	w := Word{"So", 500}
	line := []Word{w}
	w = Word{"then", 500}
	line = append(line, w)
	w = Word{"you", 500}
	line = append(line, w)
	w = Word{"see", 500}
	line = append(line, w)
	w = Word{"the", 500}
	line = append(line, w)
	w = Word{"deeper", 900}
	line = append(line, w)
	w = Word{"layer", 800}
	line = append(line, w)

	lines := [][]Word{line}

	w = Word{"of", 500}
	line = []Word{w}
	w = Word{"reality", 900}
	line = append(line, w)
	w = Word{"of", 300}
	line = append(line, w)
	w = Word{"oh,", 300}
	line = append(line, w)
	w = Word{"everyone", 900}
	line = append(line, w)
	w = Word{"is", 200}
	line = append(line, w)
	w = Word{"acting", 900}
	line = append(line, w)

	lines = append(lines, line)

	w = Word{"perfectly", 1100}
	line = []Word{w}
	w = Word{"according", 1000}
	line = append(line, w)
	w = Word{"to", 300}
	line = append(line, w)
	w = Word{"their", 400}
	line = append(line, w)
	w = Word{"state", 900}
	line = append(line, w)
	lines = append(lines, line)

	w = Word{"of", 100}
	line = []Word{w}
	w = Word{"consciousness.", 1600}
	line = append(line, w)
	w = Word{"And", 300}
	line = append(line, w)
	w = Word{"that's,", 400}
	line = append(line, w)
	w = Word{"why", 400}
	line = append(line, w)
	lines = append(lines, line)

	w = Word{"reality", 1000}
	line = []Word{w}
	w = Word{"is", 100}
	line = append(line, w)
	w = Word{"perfect.", 600}
	line = append(line, w)
	lines = append(lines, line)

	fmt.Println(lines)
}

func foo() {
	dir := "nine"
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		name := file.Name()
		fmt.Println(name)
		makeBoxFrame(dir, name)
	}
	ffmpeg("9")
}

func makeBoxFrame(dir, name string) {
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()

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
	count := 0
	for {
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		count++
		if count > 0 {
			break
		}
	}
}

func MakeWords2() {
	RmRfBang()

	words := []string{"Words", "are", "in", "a", "nice", "font."}
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()
	dc.LoadFontFace("arialbd.ttf", 96)

	drawWordsWithColorOn(dc, -1, words)
	for i, _ := range words {
		drawWordsWithColorOn(dc, i, words)
	}
	ffmpeg("29.97")
}

func drawWordsWithColorOn(dc *gg.Context, index int, words []string) {
	x := 200.0
	y := 900.0
	dc.SetRGB(1, 1, 1)
	for i, word := range words {
		w, _ := dc.MeasureString(word)
		fmt.Println(w, word, x)
		dc.SetRGB(0, 0, 0)
		dc.DrawString(word, x, y)
		dc.SetRGB(1, 1, 1)
		if i == index {
			dc.SetRGB(255, 1, 1)
		}
		dc.DrawString(word, x-3, y-3)
		x += w + 23
	}

	count := 0
	for {
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		count++
		if count > 30 {
			break
		}
		frameCount++
	}
}
