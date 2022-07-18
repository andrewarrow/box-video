package video

import (
	"fmt"

	"github.com/fogleman/gg"
)

func MakeVibration() {
	RmRfBang()
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetLineWidth(6)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
}
