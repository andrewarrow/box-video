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
	for {
		rd := RiverDot{}
		rd.X = x
		rd.Y = y
		rd.C = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
		if rand.Intn(2) == 0 {
			rd.C = color.RGBA{R: 255, G: 0, B: 255, A: 0xff}
		}
		riverDots = append(riverDots, &rd)
		if len(riverDots) > 6000 {
			break
		}
	}

	for {
		dc := gg.NewContext(HD_W, HD_H)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		for _, rd := range riverDots {
			dotColor = rd.C
			ColorSizeDot(dc, rd.X, rd.Y, 1)

			xr := rand.Intn(10)
			if rand.Intn(1) == 0 {
				xr = xr * -1
			}
			yr := rand.Intn(10)
			rd.X += float64(xr)
			rd.Y += float64(yr)
			if rd.X < 0 && rd.Y > HD_H {
				rd.X = x
				rd.Y = y
			}
		}
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++

		if frameCount > 200 {
			break
		}
	}

	ffmpeg("9")

}
