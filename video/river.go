package video

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type RiverDot struct {
	X     float64
	Y     float64
	C     color.RGBA
	Label string
}

func MakeRiver() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	x = x + (x / 2.0)
	x = x - 60
	y = 9
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	dotColor = white
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	leftEdge := DrawRiverLine(dc, x, y)
	rightEdge := DrawRiverLine(dc, HD_W, 9)

	fmt.Println(len(leftEdge), len(rightEdge))

	ffmpeg("18")
}

func DrawRiverLine(dc *gg.Context, x, y float64) []gg.Point {
	items := []gg.Point{}
	for {
		ColorSizeDot(dc, x, y, 1)
		items = append(items, gg.Point{x, y})

		xr := rand.Intn(13) * -1
		yr := rand.Intn(10)
		x += float64(xr)
		y += float64(yr)
		if y >= HD_H {
			break
		}
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		fmt.Println(frameCount)
	}
	return items
}

func MakeRiver2() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	x = x + (x / 2.0)
	y = 0

	riverDots := []*RiverDot{}
	for {
		rd := RiverDot{}
		rd.X = x
		rd.Y = y
		xr := rand.Intn(300) * -1
		yr := rand.Intn(HD_H)
		rd.X += float64(xr)
		rd.Y += float64(yr)
		if rd.Y >= HD_H {
			rd.X = x
			rd.Y = y
		}
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

			xr := rand.Intn(10) * -1
			yr := rand.Intn(10)
			rd.X += float64(xr)
			rd.Y += float64(yr)
			if rd.Y >= HD_H {
				rd.X = x
				rd.Y = y
			}
		}
		frameCount++

		if frameCount > 200 {
			break
		}
	}

	frameCount = 0
	r1 := byte(rand.Intn(256))
	g1 := byte(rand.Intn(256))
	b1 := byte(rand.Intn(256))
	r2 := byte(rand.Intn(256))
	g2 := byte(rand.Intn(256))
	b2 := byte(rand.Intn(256))

	for {
		dc := gg.NewContext(HD_W, HD_H)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		for _, rd := range riverDots {
			dotColor = rd.C
			ColorSizeDot(dc, rd.X, rd.Y, 1)

			xr := rand.Intn(10) * -1
			yr := rand.Intn(10)
			rd.X += float64(xr)
			rd.Y += float64(yr)
			if rd.Y >= HD_H {
				rd.X = x
				rd.Y = y
				if rd.Label == "" {
					rd.C = color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
					if rand.Intn(2) == 0 {
						rd.C = color.RGBA{R: 255, G: 0, B: 255, A: 0xff}
					}
					rd.Label = "white"
				} else if rd.Label == "white" {
					rd.C = color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
					if rand.Intn(2) == 0 {
						rd.C = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
					}
					rd.Label = "gold"
					r1 = byte(rand.Intn(256))
					g1 = byte(rand.Intn(256))
					b1 = byte(rand.Intn(256))
					r2 = byte(rand.Intn(256))
					g2 = byte(rand.Intn(256))
					b2 = byte(rand.Intn(256))
				} else if rd.Label == "gold" {
					rd.C = color.RGBA{R: 215, G: 0, B: 200, A: 0xff}
					if rand.Intn(2) == 0 {
						rd.C = color.RGBA{R: 100, G: 100, B: 200, A: 0xff}
					}
					rd.Label = "rand"
				} else {
					rd.C = color.RGBA{R: r1, G: g1, B: b1, A: 0xff}
					if rand.Intn(2) == 0 {
						rd.C = color.RGBA{R: r2, G: g2, B: b2, A: 0xff}
					}
					rd.Label = ""
				}
			}
		}
		dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		fmt.Println(frameCount)

		if frameCount > 9800 {
			break
		}
	}

	ffmpeg("9")

}
