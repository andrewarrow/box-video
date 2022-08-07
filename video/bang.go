package video

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

// 1280 x 720 = 921,600

// / 900 = 1024

func MakeBang() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	colors := []color.RGBA{}
	for i := 0; i < 1024; i++ {
		ChangeColors()
		c := color.RGBA{R: r1, G: g1, B: b1, A: 0xff}
		colors = append(colors, c)
	}

	i := 0
	for {
		rd := RiverDot{}
		rd.X = int(x)
		rd.Y = int(y)
		rd.C = colors[rand.Intn(1024)]
		riverDots = append(riverDots, &rd)
		if i > 921600 {
			break
		}
		i++
	}
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	MoveBangDots(dc)
	ffmpeg("18")

}

func MoveBangDots(dc *gg.Context) {
	for {
		c := gg.NewContextForImage(dc.Image())
		for _, dot := range riverDots {
			dotColor = dot.C
			ColorSizeDot(c, float64(dot.X), float64(dot.Y), 1)

			xr := rand.Intn(10)
			if rand.Intn(2) == 0 {
				xr = xr * -1
			}
			yr := rand.Intn(10)
			if rand.Intn(2) == 0 {
				yr = yr * -1
			}
			dot.X += xr
			dot.Y += yr
		}
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		fmt.Println(frameCount)
		if frameCount > 10 {
			break
		}
	}
}
