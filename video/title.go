package video

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func MakeTitle(title, path string, i int) {
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	existing, e := gg.LoadPNG(path)
	fmt.Println(e, path)

	pattern := gg.NewSurfacePattern(existing, gg.RepeatNone)
	dc.MoveTo(0, 0)
	dc.LineTo(1920, 0)
	dc.LineTo(1920, 1080)
	dc.LineTo(0, 1080)
	dc.LineTo(0, 0)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()

	dc.DrawCircle(960, 430, 400)
	pattern = gg.NewSolidPattern(color.Black)
	dc.SetFillStyle(pattern)
	dc.Fill()

	dc.LoadFontFace("arial.ttf", 96)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(title, 960, 430, 0.5, 0.5)
	j := 0
	for {
		dc.SavePNG(fmt.Sprintf("data/test%04d.png", i+j))
		j++
		if j == 30 {
			break
		}
	}
}
