package main

import (
	"time"
)

var (
	GameTicker *time.Ticker
)

func init() {
	GameTicker = time.NewTicker(1 * time.Second)
}

func HandleGame() {

	defer Wg.Done()

	// Add first, default player
	Players = append(Players, NewPlayer())
	PrintStatus("Game started ...")

	//	for t := range GameTicker.C {

	//		PrintStatus(fmt.Sprintf("%v", t))
	//	}

}
