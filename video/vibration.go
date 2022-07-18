package video

import (
	"fmt"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/raster"
)

const HD_W = 1920
const HD_H = 1080

func MakeVibration() {
	RmRfBang()
	dc := gg.NewContext(HD_W, HD_H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetLineWidth(1)

	x := HD_W / 2.0
	y := HD_H / 2.0

	p := PathFromTo(x, y, x+300, y+300)
	fmt.Println(p)
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

	dc.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
}

func PathFromTo(x1, y1, x2, y2 float64) []gg.Point {
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

func SetNiceBlue(dc *gg.Context) {
	dc.SetRGB(0, 15, 55)
}
