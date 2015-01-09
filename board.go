package main

import (
	"github.com/nsf/termbox-go"
	"time"
)

func HandleBoard() {

	termbox.Init()
	defer termbox.Close()

	termbox.HideCursor()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
	termbox.SetOutputMode(termbox.OutputNormal)
	termbox.Sync()

	termbox.SetCell(0, 0, 'a', termbox.ColorRed, termbox.ColorGreen)
	termbox.SetCell(10, 10, 'b', termbox.ColorRed, termbox.ColorGreen)

	termbox.Flush()

	time.Sleep(time.Duration(5) * time.Second)

}
