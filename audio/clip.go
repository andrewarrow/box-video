package audio

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"golang.org/x/term"
)

func PlayForClip(filename string) {
	f, _ := os.Open(filename)
	var streamer beep.StreamSeekCloser
	streamer, globalFormat, _ = mp3.Decode(f)
	globalMaxLength = streamer.Len()
	//fmt.Println(globalMaxLength, globalFormat.SampleRate)
	defer streamer.Close()

	speaker.Init(globalFormat.SampleRate, globalFormat.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, volume)
	speedy.SetRatio(speedy.Ratio() - 0.5)

	speaker.Play(speedy)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	ego := Word{"Ego", 1000}
	is := Word{"is", 1000}
	the := Word{"the", 1000}
	minds := Word{"mind's", 1000}
	war := Word{"war", 1000}
	against := Word{"against", 1000}
	words = []*Word{&ego, &is, &the, &minds, &war, &against}
	go DisplayWords()

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		if c == 3 {
			speaker.Lock()
			term.Restore(int(os.Stdin.Fd()), oldState)
			fmt.Println("")
			break
		} else if c == 108 { // L
			speaker.Lock()
			streamer.Seek(streamer.Position() + 100000)
			globalFrom = streamer.Position()
			//percentDone := float64(globalFrom) / float64(globalMaxLength)
			//fmt.Printf("%s\b\b", "hi")
			speaker.Unlock()
		} else if c == 106 { // J
			speaker.Lock()
			streamer.Seek(0)
			wordReset = true
			//globalFrom = streamer.Position()
			//percentDone := float64(globalFrom) / float64(globalMaxLength)
			//fmt.Printf("%s\b\b", "ih")
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

var wordIndex = 0
var wordChars = 0
var wordReset = false
var wordMutex sync.Mutex
var words []*Word

type Word struct {
	Word string
	Time int
}

func DisplayWords() {
	i := 0
	wordChars = 0
	for {
		if wordReset {
			wordReset = false
			break
		}
		wordMutex.Lock()
		wordChars += len(words[i].Word) + 1
		fmt.Printf("%s ", words[i].Word)
		wordMutex.Unlock()
		time.Sleep(time.Millisecond * time.Duration(words[i].Time))
		i++
		if i >= len(words) {
			break
		}
	}

	for i := 0; i < wordChars; i++ {
		fmt.Printf("\b")
	}
	for i := 0; i < wordChars; i++ {
		fmt.Printf(" ")
	}
	for i := 0; i < wordChars; i++ {
		fmt.Printf("\b")
	}
	go DisplayWords()
}
