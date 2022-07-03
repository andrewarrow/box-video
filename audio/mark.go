package audio

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"golang.org/x/term"
)

func PlayForMark(filename string) {
	f, _ := os.Open(filename)
	var streamer beep.StreamSeekCloser
	streamer, globalFormat, _ = mp3.Decode(f)
	defer streamer.Close()

	speaker.Init(globalFormat.SampleRate, globalFormat.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	speaker.Play(ctrl)
	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		if c == 3 || c == 32 {
			speaker.Lock()
			pos := streamer.Position()
			fmt.Println(pos, PositionAsSeconds(pos))
			term.Restore(int(os.Stdin.Fd()), oldState)
			speaker.Unlock()
			break
		}
	}
}
