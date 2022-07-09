package video

import (
	"fmt"

	"github.com/fogleman/gg"
)

func MakeWords() {
	RmRfBang()
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 30, 30)
	dc.Clear()
	dc.LoadFontFace("arialbd.ttf", 96)
	x := float64(1920 / 2)
	y := float64(900)
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Words are in a nice font.", x, y, 0.5, 0.5)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored("Words are in a nice font.", x-2, y-2, 0.5, 0.5)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 1))
}
