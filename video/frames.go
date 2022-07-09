package video

import (
	"fmt"
	"os"
	"os/exec"
)

// 59.96
func MakeFrames(filename, name, fps string) {
	exec.Command("rm", "-rf", name).CombinedOutput()
	os.Mkdir(name, 0755)
	cmd := exec.Command("ffmpeg", "-i", filename, "-vf", fmt.Sprintf("fps=%s", fps), name+"/img%07d.png")
	o, _ := cmd.CombinedOutput()
	fmt.Println(string(o))
}

// ffmpeg -i perfect.mp4 -filter:v fps=fps=9 perfect9.mp4
