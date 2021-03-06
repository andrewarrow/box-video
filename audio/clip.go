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

func PlayForClip(filename, wordfile string) {
	ReadWords(wordfile)
	ReadTimes("times.txt")

	f, _ := os.Open(filename)
	var streamer beep.StreamSeekCloser
	streamer, globalFormat, _ = mp3.Decode(f)
	globalMaxLength = streamer.Len()
	//fmt.Println(globalMaxLength, globalFormat.SampleRate)
	defer streamer.Close()

	speaker.Init(globalFormat.SampleRate, globalFormat.SampleRate.N(time.Second/10))

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	ctrl.Paused = true
	globalPauseOn = true
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speedy := beep.ResampleRatio(4, 1, volume)
	//speedy.SetRatio(speedy.Ratio() - 0.5)

	speaker.Play(speedy)

	oldState, _ := term.MakeRaw(int(os.Stdin.Fd()))

	words = wordLines[wordLineIndex]
	go DisplayWords()
	go IncrementWordIndex()

	for {
		b := make([]byte, 1)
		os.Stdin.Read(b)
		c := b[0]
		if c == 3 {
			speaker.Lock()
			term.Restore(int(os.Stdin.Fd()), oldState)
			fmt.Println("")
			fmt.Println("Saving times...")
			SaveTimes()
			fmt.Println("")
			break
		} else if c >= 48 && c <= 57 { // 0-9
			wordIndex = int(c) - 48
			wordChange = true
		} else if (c == 67 || c == 93) && wordIndex < len(words)-1 { // -> ]
			wordIndex += 1
			wordChange = true
		} else if (c == 68 || c == 91) && wordIndex > 0 { // <- [
			wordIndex -= 1
			wordChange = true
		} else if c == 45 { // -
			words[wordIndex].Time -= 100
			wordChange = true
		} else if c == 61 { // +
			words[wordIndex].Time += 100
			wordChange = true
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
			speaker.Unlock()
			//globalFrom = streamer.Position()
			//percentDone := float64(globalFrom) / float64(globalMaxLength)
			//fmt.Printf("%s\b\b", "ih")
		} else if c == 107 || c == 32 { // K or space
			speaker.Lock()
			if ctrl.Paused == false {
				globalTo = streamer.Position()
				globalPauseOn = true
				ctrl.Paused = true
			} else {
				globalFrom = streamer.Position()
				streamer.Seek(0)
				wordLineIndex = 0
				words = wordLines[wordLineIndex]
				wordIndex = 0
				wordChange = true
				globalPauseOn = false
				ctrl.Paused = false
			}
			speaker.Unlock()
		}
	}

}

func DisplayWords() {
	wordChange = true
	for {
		if wordChange == false {
			time.Sleep(time.Nanosecond * 10)
			continue
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
		wordChars = 0
		for i, word := range words {
			txt := fmt.Sprintf("%s(%d) ", word.Word, word.Time)
			if i == wordIndex {
				txt = fmt.Sprintf("|%s(%d)| ", word.Word, word.Time)
			}
			wordChars += len(txt)
			fmt.Printf(txt)
		}

		wordChange = false
	}
}

func IncrementWordIndex() {
	wordIndex = 0
	wordChange = true
	for {
		if globalPauseOn {
			time.Sleep(time.Nanosecond * 10)
			continue
		}
		time.Sleep(time.Millisecond * time.Duration(words[wordIndex].Time))
		wordIndex++
		if wordIndex >= len(words) {
			wordLineIndex++
			words = wordLines[wordLineIndex]
			wordIndex = 0
		}
		wordChange = true
	}
}
