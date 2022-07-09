package video

import (
	"fmt"
	"os/exec"
)

func AddSound(video, audio string) {
	cmd := exec.Command("ffmpeg", "-i", video, "-i", audio, "-c", "copy", "-map", "0:v:0", "-map", "1:a:0", "data/sound.mov")
	o, _ := cmd.CombinedOutput()
	fmt.Println(string(o))
}
