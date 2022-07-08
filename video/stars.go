package video

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"os/exec"

	"github.com/fogleman/gg"
)

type Star struct {
	X  int
	Y  int
	On bool
}

var list = []*Star{}
var frameCount = 0

func CountActive() int {
	count := 0
	for _, star := range list {
		if !star.On {
			continue
		}
		count++
	}
	return count
}

func MakeStars() {
	exec.Command("rm", "-rf", "data").CombinedOutput()
	os.Mkdir("data", 0755)

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
	pattern := gg.NewSolidPattern(color.White)
	for {
		dc := gg.NewContext(1920, 1080)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
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
		fmt.Println(count, i, frameCount)
		makeLineGoingUp(dc, i)

		for _, star := range list {
			if !star.On {
				continue
			}
			chance := rand.Intn(100)
			if chance <= 2 {
				star.On = false
			}
		}

		dc = gg.NewContext(1920, 1080)
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		dc.SetRGB(1, 1, 1)
		dc.SetFillStyle(pattern)

		count = 0
		for _, star := range list {
			if !star.On {
				continue
			}
			count++
			dc.DrawCircle(float64(star.X), float64(star.Y), 1)
			dc.Fill()
		}
		fmt.Println(count, i, frameCount)
		makeLineGoingDown(dc, i)

		for _, star := range list {
			if !star.On {
				continue
			}
			chance := rand.Intn(100)
			if chance <= 2 {
				star.On = false
			}
		}

		i++
		if CountActive() == 0 {
			break
		}
	}
	dc := gg.NewContext(1920, 1080)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	makeLineGoingDown(dc, i)
	ffmpeg()
}

func makeLineGoingUp(dc *gg.Context, i int) {
	j := 1069
	var c *gg.Context
	for {
		color := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
		pattern := gg.NewSolidPattern(color)
		c = gg.NewContextForImage(dc.Image())
		c.SetFillStyle(pattern)
		c.DrawRectangle(0, float64(j), 1920, 10)
		c.Fill()
		fmt.Println(frameCount)
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		j -= 4
		frameCount++
		if j < 0 {
			break
		}
	}
}

func makeLineGoingDown(dc *gg.Context, i int) {
	var c *gg.Context
	j := 0
	for {
		color := color.RGBA{R: 255, G: 215, B: 0, A: 0xff}
		pattern := gg.NewSolidPattern(color)
		c = gg.NewContextForImage(dc.Image())
		c.SetFillStyle(pattern)
		c.DrawRectangle(0, float64(j), 1920, 10)
		c.Fill()
		fmt.Println(frameCount)
		c.SavePNG(fmt.Sprintf("data/img%07d.png", frameCount))
		j += 4
		frameCount++
		if j > 1069 {
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
	//ffmpeg -framerate 29.97 -pattern_type glob -i data/*.png -c:v libx264 -pix_fmt yuv420p data/temp.mov
	exec.Command("/usr/local/bin/ffmpeg", "-framerate", "29.97", "-pattern_type", "glob", "-i", "data/*.png", "-c:v", "libx264",
		"-pix_fmt", "yuv420p", "data/temp.mov").CombinedOutput()
	fmt.Println("ffmpeg")
}
