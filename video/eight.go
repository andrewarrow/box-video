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
	dc.SetLineWidth(6)

	//w := 100
	//h := 100
	color1 := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
	dc.SetRGB(0, 40, 255)
	pattern := gg.NewSolidPattern(color1)
	dc.SetFillStyle(pattern)
	dc.DrawArc(200, 400, 200, 0, 2.3)
	dc.Stroke()
	color2 := color.RGBA{R: 215, G: 255, B: 0, A: 0xff}
	dc.SetRGB(40, 0, 255)
	pattern = gg.NewSolidPattern(color2)
	dc.SetFillStyle(pattern)
	dc.DrawArc(260, 540, -200, 0, 2.3)
	dc.Stroke()
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
}
