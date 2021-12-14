package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/eiannone/keyboard"
)

type state int

const (
	Working state = iota
	Break
	LongBreak
	Pause
)

var (
	currentState    = Working
	prevState       = Working
	currentDuration = 25 * 60
	pomodoros       = 1
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer cancel()
	go timer(ctx)
	go keys(cancel)
	<-ctx.Done()
}

func keys(cancel func()) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		switch string(char) {
		case "p":
			if currentState != Pause {
				prevState = currentState
				currentState = Pause
			} else {
				currentState = prevState
			}
		case "n":
			currentDuration = 0
		case "q":
			cancel()
			return
		}
	}
}

func timer(ctx context.Context) {
	needClear := false
	for {
		if currentState != Pause {
			currentDuration--
		}
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
		if needClear {
			clear()
		}
		needClear = true
		displayTimer()

		select {
		case <-time.After(1 * time.Second):
			// Do nothing
		case <-ctx.Done():
			return
		}
	}
}

func displayTimer() {
	stateText := ""
	if currentState == Pause {
		stateText = fmt.Sprintf("[Paused] %s", getStateText(prevState))
	} else {
		stateText = getStateText(currentState)
	}
	fmt.Printf("%s – %s\n(keys: p - pause/resume timer, n - next phase, q - quit)", stateText, secondsToMinutes(currentDuration))
}

func getStateText(state state) string {
	switch state {
	case Working:
		return fmt.Sprintf("🍅 #%d Working", pomodoros)
	case Break:
		return fmt.Sprintf("☕️ #%d Break", pomodoros)
	case LongBreak:
		return fmt.Sprintf("☕️ #%d Long break", pomodoros)
	}
	return ""
}

func bell() {
	fmt.Print("\u0007")
}

func clear() {
	fmt.Print("\u001B[2K\u001B[F\u001B[2K\r")
}

func secondsToMinutes(inSeconds int) string {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	return str
}
