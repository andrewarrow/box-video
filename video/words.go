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
	x := 200.0 //float64(1920 / 2)
	y := float64(900)

	/*
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored("Words are in a nice font.", x, y, 0.5, 0.5)
		dc.SetRGB(1, 1, 1)
		dc.DrawStringAnchored("Words are in a nice font.", x-3, y-3, 0.5, 0.5)
	*/

	dc.SetRGB(1, 1, 1)
	for _, word := range words {
		w, _ := dc.MeasureString(word)
		fmt.Println(w, word, x)
		dc.SetRGB(1, 1, 1)
		dc.DrawString(word, x, y)

		x += w + 23
	}

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 1))
}
