package audio

import (
	"fmt"
	"os/exec"
)

func Merge() {
	// ffmpeg -y -i audio1.mp3 -i audio2.mp3 -filter_complex "[0:0]volume=0.09[a];[1:0]volume=1.8[b];[a][b]amix=inputs=2:duration=longest" -c:a libmp3lame output.mp3
	p := []string{}
	p = append(p, "-y")
	p = append(p, "-i")
	p = append(p, "data/andrew2.mp3")
	p = append(p, "-i")
	p = append(p, "data/final.mp3")
	p = append(p, "-filter_complex")
	p = append(p, "[0:0]volume=0.09[a];[1:0]volume=1.8[b];[a][b]amix=inputs=2:duration=longest")
	p = append(p, "-c:a")
	p = append(p, "libmp3lame")
	p = append(p, "data/merged.mp3")
	exec.Command("ffmpeg", p...).CombinedOutput()
	fmt.Println("done")

}
