package video

import (
	"fmt"
	"image/color"
	"math/rand"
	"sort"

	"github.com/fogleman/gg"
)

// 1280 x 720 = 921,600

// / 900 = 1024
var bangEdge = false

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
		rd.Move = 3
		rd.C = colors[rand.Intn(1024)]
		rd.SingleC = int(rd.C.R) + int(rd.C.G) + int(rd.C.B)
		riverDots = append(riverDots, &rd)
		if i > 921600 {
			break
		}
		i++
	}
	sort.SliceStable(riverDots, func(i, j int) bool {
		return riverDots[i].SingleC > riverDots[j].SingleC
	})
	goalX := 0
	goalY := 0
	for _, r := range riverDots {
		r.GoalX = goalX
		r.GoalY = goalY
		goalX++
		if goalX > int(HD_W) {
			goalY++
			goalX = 0
		}
	}
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	MoveBangDots(dc)
	ffmpeg("9")

}

func MoveBangDots(dc *gg.Context) {
	for {
		c := gg.NewContextForImage(dc.Image())
		for _, dot := range riverDots {
			dotColor = dot.C
			ColorSizeDot(c, float64(dot.X), float64(dot.Y), 1)

			if bangEdge == false {
				xr := rand.Intn(dot.Move)
				if rand.Intn(2) == 0 {
					xr = xr * -1
				}
				yr := rand.Intn(dot.Move)
				if rand.Intn(2) == 0 {
					yr = yr * -1
				}
				dot.X += xr
				dot.Y += yr
				if dot.Y > int(HD_H) || dot.Y < 0 {
					bangEdge = true
				}
				dot.Move++
			} else {
				unit := rand.Intn(9)
				if dot.X > dot.GoalX {
					dot.X -= unit
				}
				if dot.Y > dot.GoalY {
					dot.Y -= unit
				}
				if dot.X < dot.GoalX {
					dot.X += unit
				}
				if dot.Y < dot.GoalY {
					dot.Y += unit
				}
			}
		}
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		frameCount++
		fmt.Println(frameCount)
		if frameCount > 400 {
			break
		}
	}
}