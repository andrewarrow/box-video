package video

import (
	"fmt"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
)

var pattern gg.Pattern = gg.NewSolidPattern(color.White)

func MakeEight() {
	RmRfBang()
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()
	dc.SetLineWidth(6)

	x := 400.0
	y := 400.0

	dc.SetRGB(0, 40, 255)
	dc.DrawArc(x, y, 200, 0, 2.3)
	dc.Stroke()
	dc.SetRGB(40, 0, 255)
	dc.DrawArc(60+x, 140+y, -200, 0, 2.3)
	dc.Stroke()

	x = 746.0
	y = 246.0

	dc.SetRGB(40, 0, 255)
	dc.DrawArc(x, y, 200, 0, 2.3)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	dc.DrawArc(60+x, 140+y, -200, 0, 2.3)
	dc.Stroke()

	x = 1300.0
	y = 400.0

	dc.SetRGB(40, 0, 255)
	var p raster.Path
	p.Start(Fixed(x, y))
	fmt.Println(p)
	p.Add1(Fixed(x+200, y+400))
	fmt.Println(p)

	var painter raster.Painter
	painter = EightPainter{}
	r := raster.NewRasterizer(1920, 1080)
	r.UseNonZeroWinding = true
	r.Clear()
	fp := flattenPath(p)
	rp := rasterPath(fp)
	fmt.Println(rp)
	r.AddStroke(rp, fix(24), raster.RoundCapper, raster.RoundJoiner)
	r.Rasterize(painter)

	dc.DrawLine(x, y, x+200, y+400)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	dc.DrawLine(x, y, x-200, y+400)
	dc.Stroke()

	x, y = MakeDotGoingDown(dc, x, y)
	MakeDotGoingUp(dc, x, y)

	//dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
	ffmpeg("9")
}

func Fixed(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{fix(x), fix(y)}
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func WhiteDot(dc *gg.Context, x, y float64) {
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, 24)
	dc.Fill()
}

type EightPainter struct{}

func (ep EightPainter) Paint(ss []raster.Span, done bool) {
	fmt.Println(" ")
	lasty := ss[0].Y
	last := ss[0]
	for _, s := range ss {
		if s.Y != lasty {
			fmt.Println(last.X0, last.Y, done)
		}
		lasty = s.Y
		last = s
	}
	fmt.Println(last.X0, last.Y, done)
}

func MakeDotGoingDown(dc *gg.Context, x, y float64) (float64, float64) {
	myx := x
	myy := y
	var c *gg.Context
	for {
		fmt.Println(frameCount)
		c = gg.NewContextForImage(dc.Image())
		WhiteDot(c, myx, myy)
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		myy += 32
		myx -= 32
		frameCount++
		if myy > y+400 {
			break
		}
	}
	return myx, myy
}
func MakeDotGoingUp(dc *gg.Context, x, y float64) (float64, float64) {
	myx := x
	myy := y
	var c *gg.Context
	for {
		fmt.Println(frameCount)
		c = gg.NewContextForImage(dc.Image())
		WhiteDot(c, myx, myy)
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		myy -= 32
		fmt.Println(myy, y)
		myx += 32
		frameCount++
		if myy < y-420 {
			break
		}
	}
	return myx, myy
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
