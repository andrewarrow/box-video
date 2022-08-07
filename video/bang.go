package video

import (
	"fmt"
	"image/color"
)

// 1280 x 720 = 921,600

// / 900 = 1024

func MakeBang() {
	colors := []color.RGBA{}
	for i := 0; i < 1024; i++ {
		ChangeColors()
		c := color.RGBA{R: r1, G: g1, B: b1, A: 0xff}
		colors = append(colors, c)
	}

	fmt.Println(colors)
}
