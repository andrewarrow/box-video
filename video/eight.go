package video

import (
	"fmt"
	"image/color"
	"math"
	"sort"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
	"golang.org/x/image/math/fixed"
)

type EightData struct {
	P          gg.Point
	UpsideDown bool
	Side       string
}

var eightChannel chan EightData = make(chan EightData, 1)
var upsideDownLeft bool
var upsideDownRight bool
var leftDotx float64
var leftDoty float64
var rightDotx float64
var rightDoty float64

func EightLoop(upsides ...bool) {
	x := 746.0
	y := 246.0
	MakeArcDotGoing(upsides[3], x, y, 200, 0, 2.3, false)
	x = 746.0
	y = 246.0
	MakeArcDotGoing(upsides[0], 60+x, 140+y, -200, 0, 2.3, true)
	x = 400.0
	y = 400.0
	MakeArcDotGoing(upsides[1], x, y, 200, 0, 2.3, true)
	MakeArcDotGoing(upsides[2], 60+x, 140+y, -200, 0, 2.3, false)
}

func ReadEightChannelData() {
	for thing := range eightChannel {
		renderEightFrame(thing)
	}
}

func MakeEight() {
	RmRfBang()

	go ReadEightChannelData()

	go func() {
		EightLoop(false, false, false, false)
		EightLoop(true, true, true, true)
	}()

	x := 1300.0
	y := 400.0
	//gold := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
	//red := color.RGBA{R: 255, G: 0, B: 0, A: 0xff}
	//white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	//black := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}

	MakeDotGoing(x, y, x+200, y+400, true, false)
	MakeDotGoing(x+200, y+400, x, y, false, true)
	MakeDotGoing(x, y, x-200, y+400, true, true)
	MakeDotGoing(x-200, y+400, x, y, false, false)

	MakeDotGoing(x, y, x+200, y+400, true, false)
	MakeDotGoing(x+200, y+400, x, y, false, true)
	MakeDotGoing(x, y, x-200, y+400, true, true)
	MakeDotGoing(x-200, y+400, x, y, false, false)

	//dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
	ffmpeg("9")
}

func EightContext() *gg.Context {
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 200, 200)
	dc.Clear()
	dc.SetLineWidth(6)

	if upsideDownLeft {
		ColorDot(dc, leftDotx, leftDoty)
	}
	if upsideDownRight {
		ColorDot(dc, rightDotx, rightDoty)
	}

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

	dc.DrawLine(x, y, x+200, y+400)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	dc.DrawLine(x, y, x-200, y+400)
	dc.Stroke()

	if !upsideDownLeft {
		ColorDot(dc, leftDotx, leftDoty)
	}
	if !upsideDownRight {
		ColorDot(dc, rightDotx, rightDoty)
	}

	dc.SetRGB(255, 255, 255)
	dc.LoadFontFace("arialbd.ttf", 96)
	nine := "+9"
	if upsideDownRight {
		nine = "-9"
	}
	dc.DrawString(nine, x-60, y-40)
	dc.LoadFontFace("arialbd.ttf", 36)
	three := "+3"
	if upsideDownRight {
		three = "-3"
	}
	dc.DrawString(three, x-240, y+460)
	six := "+6"
	if upsideDownRight {
		six = "-6"
	}
	dc.DrawString(six, x+200, y+460)

	return dc
}

func MakeArcDotGoing(upsideDown bool, x, y, r, angle1, angle2 float64, sortBool bool) {
	var p raster.Path
	p.Start(Fixed(x, y))
	fmt.Println(p)
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

		if i == 0 {
			p.Add1(Fixed(x0, y0))
		}
		p.Add2(Fixed(cx, cy), Fixed(x2, y2))
	}
	fmt.Println(p)
	//fmt.Println("flattenPath", flattenPath(p))
	//fmt.Println("rasterPath", rasterPath(flattenPath(p)))
	ep := &EightPainter{}
	ep.Points = []gg.Point{}
	ras := raster.NewRasterizer(1920, 1080)
	ras.UseNonZeroWinding = true
	ras.Clear()
	ras.AddStroke(p, fix(1), raster.SquareCapper, raster.RoundJoiner)
	ras.Rasterize(ep)
	ep.SortPoints(sortBool)
	fmt.Println(ep.Points)
	for i := 0; i < len(ep.Points); i++ {
		if i%60 != 0 {
			continue
		}
		p := ep.Points[i]
		eightChannel <- EightData{p, upsideDown, "left"}
	}
}

func MakeDotGoing(x1, y1, x2, y2 float64,
	appendAtEnd bool, upsideDown bool) {

	var p raster.Path
	p.Start(Fixed(x1, y1))
	fmt.Println(p)
	p.Add1(Fixed(x2, y2))
	fmt.Println(p)

	ep := &EightPainter{}
	ep.Points = []gg.Point{}
	ep.AppendAtEnd = appendAtEnd

	r := raster.NewRasterizer(1920, 1080)
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(p, fix(0.1), raster.SquareCapper, raster.RoundJoiner)
	r.Rasterize(ep)

	if appendAtEnd {
		for i := 0; i < len(ep.Points); i++ {
			if i%40 != 0 {
				continue
			}
			p := ep.Points[i]
			eightChannel <- EightData{p, upsideDown, "right"}
		}
	} else {
		for i := len(ep.Points) - 1; i > 0; i-- {
			if i%40 != 0 {
				continue
			}
			p := ep.Points[i]
			eightChannel <- EightData{p, upsideDown, "right"}
		}
	}
}

func renderEightFrame(ed EightData) {
	var c *gg.Context
	fmt.Println(frameCount)
	if ed.Side == "left" {
		upsideDownLeft = ed.UpsideDown
		leftDotx = ed.P.X
		leftDoty = ed.P.Y
	} else {
		upsideDownRight = ed.UpsideDown
		rightDotx = ed.P.X
		rightDoty = ed.P.Y
	}
	c = EightContext()
	c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
	frameCount++
}

func Fixed(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{fix(x), fix(y)}
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func ColorDot(dc *gg.Context, x, y float64) {
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	pattern := gg.NewSolidPattern(white)
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, 24)
	dc.Fill()
}

type EightPainter struct {
	Points      []gg.Point
	AppendAtEnd bool
}

func (ep *EightPainter) Paint(ss []raster.Span, done bool) {
	for _, s := range ss {
		np := gg.Point{float64(s.X0), float64(s.Y)}
		ep.Points = append(ep.Points, np)
		np = gg.Point{float64(s.X1), float64(s.Y)}
		ep.Points = append(ep.Points, np)
	}
}
func (ep *EightPainter) SortPoints(b bool) {
	if b == true {
		sort.SliceStable(ep.Points, func(i, j int) bool {
			return ep.Points[i].X > ep.Points[j].X
		})
	} else {
		sort.SliceStable(ep.Points, func(i, j int) bool {
			return ep.Points[i].X < ep.Points[j].X
		})
	}
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
