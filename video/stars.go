package video

import (
	"fmt"
	"image/color"
	"math/rand"
	"os/exec"

	"github.com/fogleman/gg"
)

type Star struct {
	X  int
	Y  int
	On bool
}

var list = []*Star{}

func MakeStars() {
	i := 0
	for {
		x := rand.Intn(1920)
		y := rand.Intn(1080)
		s := Star{x, y, true}
		list = append(list, &s)
		if i > 50000 {
			break
		}
		i++
	}

	i = 0
	for {
		dc := gg.NewContext(1920, 1080)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
		pattern := gg.NewSolidPattern(color.White)
		dc.SetFillStyle(pattern)

		count := 0
		for _, star := range list {
			if !star.On {
				continue
			}
			count++
			dc.DrawCircle(float64(star.X), float64(star.Y), 1)
			dc.Fill()
		}
		fmt.Println(count, i)
		makeLineGoingUp(dc, i)

		if true {
			break
		}
	}
	ffmpeg()
}

func makeLineGoingUp(dc *gg.Context, i int) {
	j := 1069
	for {
		color := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
		pattern := gg.NewSolidPattern(color)
		c := gg.NewContextForImage(dc.Image())
		c.SetFillStyle(pattern)
		c.DrawRectangle(0, float64(j), 1920, 10)
		c.Fill()
		c.SavePNG(fmt.Sprintf("data/img%07d.png", j))
		j--
		if j < 0 {
			break
		}
	}
}

func MakeStars2() {

	i := 0
	for {
		x := rand.Intn(1920)
		y := rand.Intn(1080)
		s := Star{x, y, true}
		list = append(list, &s)
		if i > 50000 {
			break
		}
		i++
	}

	i = 0
	for {
		dc := gg.NewContext(1920, 1080)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
		pattern := gg.NewSolidPattern(color.White)
		dc.SetFillStyle(pattern)

		count := 0
		for _, star := range list {
			if !star.On {
				continue
			}
			count++
			dc.DrawCircle(float64(star.X), float64(star.Y), 1)
			dc.Fill()
		}
		fmt.Println(count, i)

		// save 300 with line going up
		// remove N stars
		// save 300 with line going down

		j := 0
		for {
			dc.SavePNG(fmt.Sprintf("data/img%07d.png", (i*30)+j))
			j++
			if j == 30 {
				break
			}
		}

		for _, star := range list {
			chance := rand.Intn(100)
			if chance <= 2 {
				star.On = false
			}
		}

		if i > 130 {
			break
		}
		i++
	}
	ffmpeg()
}

func ffmpeg() {
	fmt.Println("ffmpeg")
	exec.Command("/usr/local/bin/ffmpeg", "-framerate", "29.97", "-pattern_type", "glob", "-i", "data/*.png", "-c:v", "libx264",
		"-pix_fmt", "yuv420p", "data/temp.mov").CombinedOutput()
	fmt.Println("ffmpeg")
}
