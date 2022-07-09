package video

import (
	"fmt"

	"github.com/fogleman/gg"
)

func MakeWords() {
	RmRfBang()
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.LoadFontFace("arialbd.ttf", 96)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored("Words", 1920/2, 1000, 0.5, 0.5)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 1))
}
