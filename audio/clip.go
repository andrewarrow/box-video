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

func PlayForClip(filename string) {
	f, err := os.Open(filename)
	fmt.Println(err)
	var streamer beep.StreamSeekCloser
	streamer, globalFormat, _ = mp3.Decode(f)
	globalMaxLength = streamer.Len()
	fmt.Println(globalMaxLength, globalFormat.SampleRate)
	defer streamer.Close()

	speaker.Init(globalFormat.SampleRate, globalFormat.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	speaker.Play(ctrl)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		if c == 3 {
			speaker.Lock()
			term.Restore(int(os.Stdin.Fd()), oldState)
			break
		} else if c == 108 { // L
			speaker.Lock()
			streamer.Seek(streamer.Position() + 100000)
			globalFrom = streamer.Position()
			percentDone := float64(globalFrom) / float64(globalMaxLength)
			fmt.Printf("%0.2f\n", percentDone*100)
			speaker.Unlock()
		} else if c == 106 { // J
			speaker.Lock()
			streamer.Seek(streamer.Position() - 100000)
			globalFrom = streamer.Position()
			percentDone := float64(globalFrom) / float64(globalMaxLength)
			fmt.Printf("%0.2f\n", percentDone*100)
			speaker.Unlock()
		} else if c == 107 || c == 32 { // K or space
			speaker.Lock()
			if ctrl.Paused == false {
				globalTo = streamer.Position()
				globalPauseOn = true
				ctrl.Paused = true
			} else {
				globalFrom = streamer.Position()
				globalPauseOff = true
				ctrl.Paused = false
			}
			speaker.Unlock()
		}
	}

}
