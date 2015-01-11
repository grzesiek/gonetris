package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/nsf/termbox-go"
	"os"
	"time"
)

var opts struct {
	Name    string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players int    `short:"p" long:"players" description:"Number of players" required:"true"`
}

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

}

var (
	Running bool = true
	Ticker  *time.Ticker
)

func main() {

	termbox.Init()
	defer termbox.Close()

	Ticker = time.NewTicker(100 * time.Millisecond)

	go HandleTerminal()
	go HandleBoard()
	go HandleKeys()

	for Running {
		select {
		case <-Ticker.C:
		}
	}

}
