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

var globalPauseOn bool
var globalPauseOff bool
var globalFrom int
var globalTo int
var globalFile *os.File
var globalCount int64
var globalCountLast int64

func PlayTest() {
	f, err := os.Open("test.mp3")
	fmt.Println(err)
	streamer, format, _ := mp3.Decode(f)
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

	go RecordEverything()
	speaker.Play(speedy)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		//fmt.Printf("%d\n", c)
		if c == 3 {
			speaker.Lock()
			globalTo = streamer.Position()
			WritePlayDuration()
			term.Restore(int(os.Stdin.Fd()), oldState)
			break
		} else if c == 108 { // L
			speaker.Lock()
			globalTo = streamer.Position()
			//PrintPosition(format, streamer.Position())
			streamer.Seek(streamer.Position() + 100000)
			WritePlayDuration()
			globalFrom = streamer.Position()
			//globalFile.WriteString(fmt.Sprintf("position now %d\n", globalFrom))
			//PrintPosition(format, streamer.Position())
			speaker.Unlock()
		} else if c == 106 { // J
			speaker.Lock()
			globalTo = streamer.Position()
			//PrintPosition(format, streamer.Position())
			streamer.Seek(streamer.Position() - 100000)
			WritePlayDuration()
			globalFrom = streamer.Position()
			//globalFile.WriteString(fmt.Sprintf("position now %d\n", globalFrom))
			//PrintPosition(format, streamer.Position())
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

// played for 4.3 seconds, from 0 to 11101010 pos
// paused for 8.2 second
// played for 16.17 seconds, from 21101010 to 31101010 pos
// advanced to pos 41101010
// played for 0.01 seconds, from 41101010 to 41101011 pos
// advanced to pos 51101010
// played for 0.02 seconds, from 51101010 to 51101012 pos
// played for 13.6 seconds, from 51101012 to 61101012 pos
// paused for 9 seconds

func WritePlayDuration() {
	playDuration := float64(globalCount-globalCountLast) / 1000.0
	globalFile.WriteString(fmt.Sprintf("played for %f, from %d to %d\n", playDuration, globalFrom, globalTo))
	globalCountLast = globalCount
}

func RecordEverything() {
	os.Remove("log.txt")
	globalFile, _ = os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer globalFile.Close()
	var pauseCount int64
	for {
		globalCount++
		if globalPauseOn {
			WritePlayDuration()
			pauseCount = globalCount
			globalPauseOn = false
		}
		if globalPauseOff {
			pauseDuration := float64(globalCount-pauseCount) / 1000.0
			globalFile.WriteString(fmt.Sprintf("paused for %f\n", pauseDuration))
			globalPauseOff = false
			globalCountLast = globalCount
		}

		time.Sleep(time.Millisecond)
	}
}

func PrintPosition(format beep.Format, pos int) {
	p := format.SampleRate.D(pos)
	f := float64(int64(p)) / 1000000000.0
	fmt.Println(f)
}
