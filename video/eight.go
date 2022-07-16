package video

import (
	"fmt"
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

var pattern gg.Pattern = gg.NewSolidPattern(color.White)

func MakeEight() {
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()
	dc.SetLineWidth(6)

	x := 400.0
	y := 400.0

	dc.SetRGB(0, 40, 255)
	ArcWithDot(dc, x, y, 200, 0, 2.3)
	dc.Stroke()
	dc.SetRGB(40, 0, 255)
	ArcWithDot(dc, 60+x, 140+y, -200, 0, 2.3)
	dc.Stroke()

	x = 746.0
	y = 246.0

	dc.SetRGB(40, 0, 255)
	ArcWithDot(dc, x, y, 200, 0, 2.3)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	ArcWithDot(dc, 60+x, 140+y, -200, 0, 2.3)
	dc.Stroke()

	x = 1300.0
	y = 400.0

	dc.SetRGB(40, 0, 255)
	dc.DrawLine(x, y, x+200, y+400)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	dc.DrawLine(x, y, x-200, y+400)
	dc.Stroke()

	WhiteDot(dc, x, y)

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
}

func WhiteDot(dc *gg.Context, x, y float64) {
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, 6)
	dc.Fill()
}

func ArcWithDot(dc *gg.Context, x, y, r, angle1, angle2 float64) {
	const n = 16
	for i := 0; i < n; i++ {
		p1 := float64(i+0) / n
		p2 := float64(i+1) / n
		a1 := angle1 + (angle2-angle1)*p1
		a2 := angle1 + (angle2-angle1)*p2
		x0 := x + r*math.Cos(a1)
		y0 := y + r*math.Sin(a1)
		x1 := x + r*math.Cos((a1+a2)/2)
		y1 := y + r*math.Sin((a1+a2)/2)
		x2 := x + r*math.Cos(a2)
		y2 := y + r*math.Sin(a2)
		cx := 2*x1 - x0/2 - x2/2
		cy := 2*y1 - y0/2 - y2/2

		_, hasCurrent := dc.GetCurrentPoint()

		if i == 0 {
			if hasCurrent {
				dc.LineTo(x0, y0)
			} else {
				dc.MoveTo(x0, y0)
			}
		}
		dc.QuadraticTo(cx, cy, x2, y2)
	}
}

func QuadraticTo(x1, y1, x2, y2 float64) {
	/*
		if !dc.hasCurrent {
			dc.MoveTo(x1, y1)
		}
		x1, y1 = dc.TransformPoint(x1, y1)
		x2, y2 = dc.TransformPoint(x2, y2)
		p1 := Point{x1, y1}
		p2 := Point{x2, y2}
		dc.strokePath.Add2(p1.Fixed(), p2.Fixed())
		dc.fillPath.Add2(p1.Fixed(), p2.Fixed())
		dc.current = p2*/
}

func LineTo(x, y float64) {
	/*
		if !dc.hasCurrent {
			dc.MoveTo(x, y)
		} else {
			x, y = dc.TransformPoint(x, y)
			p := Point{x, y}
			dc.strokePath.Add1(p.Fixed())
			dc.fillPath.Add1(p.Fixed())
			dc.current = p
		}*/
}
