package video

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
)

const HD_W = 1280 //1920
const HD_H = 720  //1080

var DotSize = 6.0
var lastPoints []gg.Point
var lastDots []gg.Point
var dotColor color.RGBA

func MakeVibration() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	i := 0
	for {
		if i%2 == 0 {
			dotColor = color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
		} else {
			dotColor = color.RGBA{R: 255, G: 0, B: 255, A: 0xff}
		}
		lastPoints = []gg.Point{}
		lastDots = []gg.Point{}
		DotSize = 6.0
		points := PointsFromTo(x, y, x+300, y+300) //    from zero to ONE
		FramePoints(points, true, 1)
		points = PointsFromTo(x+300, y+300, x+600, y) // from ONE to TWO
		FramePoints(points, false, 2)
		points = PointsFromTo(x+600, y, x+300, y-300) // from TWO to FOUR
		FramePoints(points, false, 4)
		points = PointsFromTo(x+300, y-300, x, y) //     from FOUR thru zero
		FramePoints(points, true, 8)
		points = PointsFromTo(x, y, x-300, y+300) //     thru zero to EIGHT
		FramePoints(points, true, 8)
		points = PointsFromTo(x-300, y+300, x-600, y) // from EIGHT to SEVEN
		FramePoints(points, false, 16)
		points = PointsFromTo(x-600, y, x-300, y-300) // from SEVEN to FIVE
		FramePoints(points, false, 32)
		points = PointsFromTo(x-300, y-300, x, y) //     from FIVE thru zero to ONE
		FramePoints(points, true, 1)
		i++
		if i > 1 {
			break
		}
	}

	ffmpeg("96")
	//SetNiceBlue(dc)
	//dc.MoveTo(x, y)
	//dc.LineTo(x+300, y+300)
	//dc.LineTo(x+600, y)
	//dc.LineTo(x+300, y-300)
	//dc.LineTo(x-300, y+300)
	//dc.LineTo(x-600, y)
	//dc.LineTo(x-300, y-300)
	//dc.LineTo(x, y)
	//dc.Stroke()
	//dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
}

func FramePoints(points []gg.Point, dir bool, size int) {
	len64 := float64(len(points))
	var lastPoint gg.Point
	if dir {
		for i := 0; i < len(points); i++ {
			if i%80 != 0 {
				continue
			}
			DrawVibrationFrame(float64(i)/len64, size, points[i])
		}
		lastPoint = points[len(points)-1]
	} else {
		for i := len(points) - 1; i > 0; i-- {
			if i%80 != 0 {
				continue
			}
			DrawVibrationFrame(float64(i)/len64, size, points[i])
		}
		lastPoint = points[0]
	}
	lastDots = append(lastDots, lastPoint)
}

func DrawVibrationFrame(per float64, size int, p gg.Point) {
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	//ColorSizeDot(dc, p.X, p.Y, 6)
	DotSize += 0.06
	for _, lp := range lastPoints {
		ColorSizeDot(dc, lp.X, lp.Y, 1)
	}

	for _, ld := range lastDots {
		ColorSizeDot(dc, ld.X, ld.Y, 10)
	}

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
	frameCount++
	fmt.Println(frameCount)
	lastPoints = append(lastPoints, p)
}

func PointsFromTo(x1, y1, x2, y2 float64) []gg.Point {
	var p raster.Path
	p.Start(Fixed(x1, y1))
	fmt.Println(p)
	p.Add1(Fixed(x2, y2))
	fmt.Println(p)
	ep := &EightPainter{}
	ep.Points = []gg.Point{}

	r := raster.NewRasterizer(HD_W, HD_H)
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(p, fix(0.1), raster.SquareCapper, raster.RoundJoiner)
	r.Rasterize(ep)

	return ep.Points
}

func ColorSizeDot(dc *gg.Context, x, y, size float64) {
	pattern := gg.NewSolidPattern(dotColor)
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, size)
	dc.Fill()
}

func SetNiceBlue(dc *gg.Context) {
	dc.SetRGB(0, 15, 55)
}
