package audio

import (
	"fmt"
	"os"
	"os/exec"
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
var globalCount int64
var globalCountLast int64
var globalFormat beep.Format
var globalPlayCount int
var globalPauseCount int
var globalMaxLength int

var globalLogFile *os.File
var globalCutFile *os.File
var globalListFile *os.File

func PlayTest() {
	f, err := os.Open("test.mp3")
	fmt.Println(err)
	var streamer beep.StreamSeekCloser
	streamer, globalFormat, _ = mp3.Decode(f)
	globalMaxLength = streamer.Len()
	fmt.Println(globalMaxLength, globalFormat.SampleRate)
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
			streamer.Seek(streamer.Position() + 100000)
			globalFrom = streamer.Position()
			percentDone := float64(globalFrom) / float64(globalMaxLength)
			fmt.Printf("%0.2f\n", percentDone*100)
			speaker.Unlock()
		} else if c == 106 { // J
			speaker.Lock()
			streamer.Seek(streamer.Position() - 100000)
			globalFrom = streamer.Position()
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

func WritePlayDuration() {
	playDuration := float64(globalCount-globalCountLast) / 1000.0
	fs := PositionAsSeconds(globalFrom)
	ts := PositionAsSeconds(globalTo)
	globalLogFile.WriteString(fmt.Sprintf("played for %f, from %s to %s\n", playDuration, fs, ts))
	cut := fmt.Sprintf("ffmpeg -i ../test.mp3 -ss %s -to %s play%d.mp3\n", fs, ts, globalPlayCount)
	globalCutFile.WriteString(cut)
	globalListFile.WriteString(fmt.Sprintf("file 'play%d.mp3'\n", globalPlayCount))
	globalPlayCount++
	globalCountLast = globalCount
}

// youtube-dl --output "%(id)s.%(ext)s"

// ffmpeg -i input.mp3 -ss 5.5 -to 10.1 output.mp3
// ffmpeg -f lavfi -i anullsrc=channel_layout=5.1:sample_rate=44100 -t 1.9 silence.mp3

//file 'output.mp3'
//file 'silence.mp3'
//file 'output.mp3'
// ffmpeg -f concat -i list.txt -codec copy final.mp3

// ffmpeg -y -i audio1.mp3 -i audio2.mp3 -filter_complex "[0:0]volume=0.09[a];[1:0]volume=1.8[b];[a][b]amix=inputs=2:duration=longest" -c:a libmp3lame output.mp3

//  ffmpeg -i andrew.mp3 -af "volumedetect" -vn -sn -dn -f null /dev/null
// ffmpeg -i ../final.mp3 -filter:a "volume=4.0" final.mp3
// https://superuser.com/questions/323119/how-can-i-normalize-audio-using-ffmpeg

func RecordEverything() {
	exec.Command("rm", "-rf", "data").CombinedOutput()
	os.Mkdir("data", 0755)
	globalLogFile, _ = os.OpenFile("data/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	globalCutFile, _ = os.OpenFile("data/cut.sh", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	globalListFile, _ = os.OpenFile("data/list.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer globalLogFile.Close()
	defer globalCutFile.Close()
	defer globalListFile.Close()
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
			globalLogFile.WriteString(fmt.Sprintf("paused for %f\n", pauseDuration))
			cut := fmt.Sprintf("ffmpeg -f lavfi -i anullsrc=channel_layout=5.1:sample_rate=%d -t %f silence%d.mp3\n", globalFormat.SampleRate, pauseDuration+0.125, globalPauseCount)
			globalCutFile.WriteString(cut)
			globalListFile.WriteString(fmt.Sprintf("file 'silence%d.mp3'\n", globalPauseCount))
			globalPauseCount++
			globalPauseOff = false
			globalCountLast = globalCount
		}

		time.Sleep(time.Millisecond)
	}
}

func PositionAsSeconds(pos int) string {
	p := globalFormat.SampleRate.D(pos)
	f := float64(int64(p)) / 1000000000.0
	return fmt.Sprintf("%f", f)
}
