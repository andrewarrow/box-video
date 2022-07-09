package video

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func MakeFrames(filename, name string) {
	exec.Command("rm", "-rf", name).CombinedOutput()
	os.Mkdir(name, 0755)

	tokens := strings.Split(filename, "/")
	file := tokens[len(tokens)-1]

	cmd := exec.Command("ffmpeg", "-i", filename, "-filter:v", "fps=fps=9", name+"/fps_"+file)
	o, _ := cmd.CombinedOutput()
	fmt.Println(string(o))

	cmd = exec.Command("ffmpeg", "-i", name+"/fps_"+file, "-vf", "fps=9", name+"/img%07d.png")
	o, _ = cmd.CombinedOutput()
	fmt.Println(string(o))
}

// ffmpeg -i perfect.mp4 -filter:v fps=fps=9 perfect9.mp4
