package video

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/fogleman/gg"
)

type RiverDot struct {
	X     int
	Y     int
	C     color.RGBA
	Label string
}

var leftEdge map[int]int
var rightEdge map[int]int
var riverDots []*RiverDot
var r1, g1, b1, r2, g2, b2 byte

func MakeRiver() {
	RmRfBang()

	x := HD_W / 2.0
	//y := HD_H / 2.0

	x = x + (x / 2.0)
	xi := int(x) - 60
	yi := int(0)
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	dotColor = white
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	leftEdge = DrawRiverLine(dc, xi, yi)
	rightEdge = DrawRiverLine(dc, int(HD_W), 0)

	fmt.Println(len(leftEdge), len(rightEdge))

	//dotColor = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
	r1 = byte(0)
	g1 = byte(255)
	b1 = byte(255)
	r2 = byte(255)
	g2 = byte(0)
	b2 = byte(255)
	AddRiverDots()
	MoveDotsDownRiver(dc)

	ffmpeg("18")
}

func ChangeColors() {
	r1 = byte(rand.Intn(256))
	g1 = byte(rand.Intn(256))
	b1 = byte(rand.Intn(256))
	r2 = byte(rand.Intn(256))
	g2 = byte(rand.Intn(256))
	b2 = byte(rand.Intn(256))
}

func AddRiverDots() {
	i := 0
	for {
		rd := RiverDot{}
		//fmt.Println(rightEdge[0]-leftEdge[0], rightEdge[0], leftEdge[0])
		rd.X = leftEdge[0] + rand.Intn(rightEdge[0]-leftEdge[0])
		rd.Y = 0
		rd.C = color.RGBA{R: r1, G: g1, B: b1, A: 0xff}
		if rand.Intn(2) == 0 {
			rd.C = color.RGBA{R: r2, G: g2, B: b2, A: 0xff}
		}
		riverDots = append(riverDots, &rd)
		if i > 60 {
			break
		}
		i++
	}
}

func MoveDotsDownRiver(dc *gg.Context) {
	for {
		c := gg.NewContextForImage(dc.Image())
		for _, dot := range riverDots {
			dotColor = dot.C
			ColorSizeDot(c, float64(dot.X), float64(dot.Y), 6)

			//fmt.Println("mddr", x, y, leftEdge[y], rightEdge[y])

			xr := rand.Intn(10)
			if rand.Intn(2) == 0 {
				xr = xr * -1
			}
			yr := rand.Intn(10)
			dot.X += xr
			dot.Y += yr
			if dot.X < leftEdge[dot.Y] || dot.X > rightEdge[dot.Y] {
				delta := rightEdge[dot.Y] - leftEdge[dot.Y]
				if delta > 0 {
					dot.X = leftEdge[dot.Y] + rand.Intn(delta)
				}
			}
		}
		if rand.Intn(40) == 0 {
			ChangeColors()
		}
		if rand.Intn(20) == 0 {
			AddRiverDots()
		}
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		fmt.Println(frameCount)
		if frameCount > 600 {
			break
		}
	}
}

func DrawRiverLine(dc *gg.Context, x, y int) map[int]int {
	// for this Y what is x?
	m := map[int]int{}
	var lastY int
	for {
		ColorSizeDot(dc, float64(x), float64(y), 1)
		m[y] = x
		if lastY > 0 {
			// lastY was 1, now we on 10, 2,3,4,5,6,7,8,9 set to x
			sub := lastY + 1
			for {
				m[sub] = x
				sub++
				if sub >= y {
					break
				}
			}
		}

		xr := rand.Intn(13) * -1
		yr := rand.Intn(10)
		x += xr
		lastY = y
		y += yr
		if y >= int(HD_H) {
			break
		}
		//dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		//frameCount++
		//fmt.Println(frameCount)
	}
	sub := lastY + 1
	for {
		m[sub] = x
		sub++
		if sub >= y {
			break
		}
	}
	return m
}

/*
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
*/
