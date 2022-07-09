package video

import (
	"fmt"

	"github.com/fogleman/gg"
)

func MakeWords() {
	RmRfBang()

	words := []string{"Words", "are", "in", "a", "nice", "font."}
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()
	dc.LoadFontFace("arialbd.ttf", 96)

	/*
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored("Words are in a nice font.", x, y, 0.5, 0.5)
		dc.SetRGB(1, 1, 1)
		dc.DrawStringAnchored("Words are in a nice font.", x-3, y-3, 0.5, 0.5)
	*/

	drawWordsWithColorOn(dc, -1, words)
	for i, _ := range words {
		drawWordsWithColorOn(dc, i, words)
	}
	ffmpeg()
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
