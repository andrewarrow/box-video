package video

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func MakeRiver() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	x = x + (x / 2.0)
	y = y - (y / 2.0)

	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dotColor = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
	ColorSizeDot(dc, x, y, 10)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
}
