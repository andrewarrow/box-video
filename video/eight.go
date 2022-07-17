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

	dc.DrawLine(x, y, x+200, y+400)
	dc.Stroke()
	dc.SetRGB(0, 40, 255)
	dc.DrawLine(x, y, x-200, y+400)
	dc.Stroke()

	gold := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
	red := color.RGBA{R: 255, G: 0, B: 0, A: 0xff}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 0xff}
	black := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	MakeDotGoing(dc, x, y, x+200, y+400, true, gold)
	MakeDotGoing(dc, x+200, y+400, x, y, false, red)
	MakeDotGoing(dc, x, y, x-200, y+400, false, white)
	MakeDotGoing(dc, x-200, y+400, x, y, true, black)

	//dc.SavePNG(fmt.Sprintf("data/img%07d.png", 0))
	ffmpeg("9")
}

func Fixed(x, y float64) fixed.Point26_6 {
	return fixed.Point26_6{fix(x), fix(y)}
}

func fix(x float64) fixed.Int26_6 {
	return fixed.Int26_6(math.Round(x * 64))
}

func ColorDot(dc *gg.Context, x, y float64, c color.RGBA) {
	pattern := gg.NewSolidPattern(c)
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, 24)
	dc.Fill()
}

type EightPainter struct {
	Points          []gg.Point
	AppendAtEnd     bool
	TheYs           map[int][]int
	TheYsUniqSorted map[int][]int
	AllXs           []int
	AllYs           []int
	YforX           map[int]int
}

/*
{799 1089 1113 65535}
{799 1113 1114 41994}
{800 1088 1089 62879}
{800 1089 1112 65535}
{800 1112 1113 64767}
{800 1113 1114 10002}
{801 1088 1089 57358}
*/
func (ep *EightPainter) Paint(ss []raster.Span, done bool) {
	for _, s := range ss {
		ep.TheYs[s.Y] = append(ep.TheYs[s.Y], s.X0, s.X1)
	}
}

func dedupAndSort(v []int) []int {
	m := map[int]bool{}
	for _, vv := range v {
		m[vv] = true
	}
	list := []int{}
	for kk, _ := range m {
		list = append(list, kk)
	}
	sort.Ints(list)
	return list
}

func (ep *EightPainter) FindSmallYForX(x int) float64 {
	return float64(ep.YforX[x])
}

// 390: [1300, 1301, 1302]
// 391: [1300, 1301, 1302]
// 392: [1301, 1302, 1303]

func (ep *EightPainter) DedupAndSortYs() {
	xs := []int{}
	ys := []int{}
	for k, v := range ep.TheYs {
		for _, x := range v {
			if k < ep.YforX[x] || ep.YforX[x] == 0 {
				ep.YforX[x] = k
			}
		}
		ep.TheYsUniqSorted[k] = dedupAndSort(v)
		xs = append(xs, v...)
		ys = append(ys, k)
	}
	ep.AllXs = dedupAndSort(xs)
	ep.AllYs = dedupAndSort(ys)
}

func (ep *EightPainter) OldPaint(ss []raster.Span, done bool) {
	lasty := ss[0].Y
	last := ss[0]
	for _, s := range ss {
		if s.Y != lasty {
			//fmt.Println(last.X1, last.Y, done)
			np := gg.Point{float64(last.X1), float64(last.Y)}
			if ep.AppendAtEnd {
				ep.Points = append(ep.Points, np)
			} else {
				ep.Points = append([]gg.Point{np}, ep.Points...)
			}
		}
		lasty = s.Y
		last = s
	}
	//fmt.Println(last.X0, last.Y, done)
	np := gg.Point{float64(last.X1), float64(last.Y)}
	if ep.AppendAtEnd {
		ep.Points = append(ep.Points, np)
	} else {
		ep.Points = append([]gg.Point{np}, ep.Points...)
	}
}

func MakeDotGoing(dc *gg.Context, x1, y1, x2, y2 float64,
	appendAtEnd bool, color color.RGBA) {

	var p raster.Path
	p.Start(Fixed(x1, y1))
	fmt.Println(p)
	p.Add1(Fixed(x2, y2))
	fmt.Println(p)

	ep := &EightPainter{}
	ep.Points = []gg.Point{}
	ep.AppendAtEnd = appendAtEnd
	ep.TheYs = map[int][]int{}
	ep.TheYsUniqSorted = map[int][]int{}
	ep.YforX = map[int]int{}

	r := raster.NewRasterizer(1920, 1080)
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(p, fix(24), raster.RoundCapper, raster.RoundJoiner)
	r.Rasterize(ep)
	ep.DedupAndSortYs()
	fmt.Println(len(ep.AllXs))
	fmt.Println(len(ep.AllYs))

	if appendAtEnd {
		for i := 0; i < len(ep.AllXs); i++ {
			if i%20 != 0 {
				continue
			}
			x := ep.AllXs[i]
			renderEightFrame(dc, float64(x), ep.FindSmallYForX(x), color)
		}
	} else {
		for i := len(ep.AllXs) - 1; i > 0; i-- {
			if i%20 != 0 {
				continue
			}
			x := ep.AllXs[i]
			renderEightFrame(dc, float64(x), ep.FindSmallYForX(x), color)
		}
	}
}

func renderEightFrame(dc *gg.Context, x, y float64, color color.RGBA) {
	var c *gg.Context
	fmt.Println(frameCount)
	c = gg.NewContextForImage(dc.Image())
	ColorDot(c, x, y, color)
	c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
	frameCount++
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
