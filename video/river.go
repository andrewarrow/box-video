package video

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type RiverDot struct {
	X float64
	Y float64
}

func MakeRiver() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	x = x + (x / 2.0)
	y = y - (y / 2.0)

	rd := RiverDot{}
	rd.X = x
	rd.Y = y

	for {
		dc := gg.NewContext(HD_W, HD_H)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dotColor = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
		ColorSizeDot(dc, rd.X, rd.Y, 10)
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++

		xr := rand.Intn(10)
		if rand.Intn(1) == 0 {
			xr = xr * -1
		}
		yr := rand.Intn(10)
		rd.X += float64(xr)
		rd.Y += float64(yr)

		if frameCount > 100 {
			break
		}
	}

	ffmpeg("9")

}
