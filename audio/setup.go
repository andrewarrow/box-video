package audio

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"golang.org/x/term"
)

func PlayTest() {
	f, err := os.Open("test.mp3")
	fmt.Println(err)
	streamer, format, err := mp3.Decode(f)
	fmt.Println(err)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, volume)

	speaker.Play(speedy)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		fmt.Printf("%d\n", c)
		if c == 3 {
			term.Restore(int(os.Stdin.Fd()), oldState)
			break
		}
	}

	/*
		for {
			fmt.Print("Press [ENTER] to pause/resume. ")
			var cmd string
			fmt.Scan(&cmd)
			fmt.Println(cmd)

			speaker.Lock()
			fmt.Println(streamer.Position())
			fmt.Println(format.SampleRate.D(streamer.Position()))
			//fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			if ctrl.Paused == false {
				ctrl.Paused = true
				fmt.Println("+")
				speedy.SetRatio(speedy.Ratio() + 0.5)
			} else {
				ctrl.Paused = false
				fmt.Println("-")
				speedy.SetRatio(speedy.Ratio() - 0.5)
			}
			speaker.Unlock()
		}
	*/
}

func ListenForKeys() {

}
