package video

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func MakeEight() {
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()

	//w := 100
	//h := 100
	color := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
	pattern := gg.NewSolidPattern(color)
	dc.SetFillStyle(pattern)
	dc.DrawArc(200, 800, 200, 0, 2)
	dc.Fill()
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
}
