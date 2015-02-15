package main

import (
	"github.com/jessevdk/go-flags"
	"sync"
	"time"
)

var opts struct {
	Name    string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players int    `short:"p" long:"players" description:"Number of players" required:"true"`
}

var (
	Running = true
	Paused  = false
	Wg      sync.WaitGroup
)

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		panic("Invalid flags !")
	}

	Tick = 200 * time.Millisecond
}

func main() {

	Wg.Add(5)

	go HandleTerminal()
	go HandleKeys()
	go HandlePlayers()
	go HandleBrick()
	go HandleBoard()
	go HandleTick()

	Wg.Wait()
}

func Quit() {
	Running = false
	close(TerminalEvent)
}
