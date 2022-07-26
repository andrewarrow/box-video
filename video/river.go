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
	C color.RGBA
}

func MakeRiver() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	x = x + (x / 2.0)
	y = y - (y / 2.0)

	riverDots := []*RiverDot{}
	rd := RiverDot{}
	rd.X = x
	rd.Y = y
	rd.C = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
	riverDots = append(riverDots, &rd)

	rd2 := RiverDot{}
	rd2.X = x - 39
	rd2.Y = y - 56
	rd2.C = color.RGBA{R: 255, G: 0, B: 255, A: 0xff}
	riverDots = append(riverDots, &rd2)

	for {
		dc := gg.NewContext(HD_W, HD_H)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		for _, rd := range riverDots {
			dotColor = rd.C
			ColorSizeDot(dc, rd.X, rd.Y, 10)

			xr := rand.Intn(10)
			if rand.Intn(1) == 0 {
				xr = xr * -1
			}
			yr := rand.Intn(10)
			rd.X += float64(xr)
			rd.Y += float64(yr)
		}
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++

		if frameCount > 100 {
			break
		}
	}

	ffmpeg("9")

}
