package video

import (
	"fmt"

	"github.com/fogleman/gg"
)

const HD_W = 1920
const HD_H = 1080

func MakeVibration() {
	RmRfBang()
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetLineWidth(1)

	x := HD_W / 2.0
	y := HD_H / 2.0
	dc.SetRGB(0, 15, 55)
	dc.MoveTo(x, y)
	dc.LineTo(x+300, y+300)
	dc.LineTo(x+600, y)
	dc.LineTo(x+300, y-300)
	dc.LineTo(x, y)
	dc.LineTo(x-300, y+300)
	dc.LineTo(x-600, y)
	dc.LineTo(x-300, y-300)
	dc.LineTo(x, y)
	dc.Stroke()

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
}
