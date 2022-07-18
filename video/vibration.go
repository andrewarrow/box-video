package video

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
)

const HD_W = 1280 //1920
const HD_H = 720  //1080

func MakeVibration() {
	RmRfBang()

	x := HD_W / 2.0
	y := HD_H / 2.0

	points := PointsFromTo(x, y, x+300, y+300)
	FramePoints(points, true)
	points = PointsFromTo(x+300, y+300, x+600, y)
	FramePoints(points, false)
	points = PointsFromTo(x+600, y, x+300, y-300)
	FramePoints(points, false)
	points = PointsFromTo(x+300, y-300, x-300, y+300)
	FramePoints(points, true)
	points = PointsFromTo(x-300, y+300, x-600, y)
	FramePoints(points, false)
	points = PointsFromTo(x-600, y, x-300, y-300)
	FramePoints(points, false)
	points = PointsFromTo(x-300, y-300, x, y)
	FramePoints(points, true)
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

func FramePoints(points []gg.Point, dir bool) {
	if dir {
		for i := 0; i < len(points); i++ {
			if i%40 != 0 {
				continue
			}
			DrawVibrationFrame(i, points[i])
		}
	} else {
		for i := len(points) - 1; i > 0; i-- {
			if i%40 != 0 {
				continue
			}
			DrawVibrationFrame(i, points[i])
		}
	}
}

func DrawVibrationFrame(i int, p gg.Point) {
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	ColorSizeDot(dc, p.X, p.Y, float64(i)/300.0)
	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
	frameCount++
	fmt.Println(frameCount)
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
	white := color.RGBA{R: 255, G: 0, B: 0, A: 0xff}
	pattern := gg.NewSolidPattern(white)
	dc.SetFillStyle(pattern)
	dc.DrawCircle(x, y, size)
	dc.Fill()
}

func SetNiceBlue(dc *gg.Context) {
	dc.SetRGB(0, 15, 55)
}
