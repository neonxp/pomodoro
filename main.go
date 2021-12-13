package main

import (
	"fmt"
	"time"
)

type state int

const (
	Working state = iota
	Break
	LongBreak
)

var (
	currentState    = Working
	currentDuration = 25 * 60
	pomodoros       = 1
)

func main() {
	for {
		clearline()
		switch currentState {
		case Working:
			fmt.Printf("🍅 #%d Working (%s)", pomodoros, secondsToMinutes(currentDuration))
		case Break:
			fmt.Printf("☕️ #%d Break (%s)", pomodoros, secondsToMinutes(currentDuration))
		case LongBreak:
			fmt.Printf("☕️ #%d Long break (%s)", pomodoros, secondsToMinutes(currentDuration))
		}
		currentDuration--
		if currentDuration < 0 {
			switch currentState {
			case Working:
				if pomodoros%4 == 0 {
					currentState = LongBreak
					currentDuration = 15 * 60
				} else {
					currentState = Break
					currentDuration = 5 * 60
				}
			case Break, LongBreak:
				currentState = Working
				currentDuration = 25 * 60
				pomodoros++
			}
			bell()
		}
		<-time.After(1 * time.Second)
	}
}

func bell() {
	fmt.Print("\a")
}

func clearline() {
	fmt.Print("\x1b[2K\r")
}

func secondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}
