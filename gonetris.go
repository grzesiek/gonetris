package main

import (
	"github.com/jessevdk/go-flags"
	"sync"
)

var opts struct {
	Name    string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players int    `short:"p" long:"players" description:"Number of players" required:"true"`
}

var (
	RunningChan = make(chan bool)
	Running     = true
	Paused      = false
	Wg          sync.WaitGroup
)

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		panic("Invalid flags !")
	}
}

func main() {

	Wg.Add(5)

	go HandleTerminal()
	go HandleKeys()
	go HandlePlayers()
	go HandleBoard()
	go HandleTick()

	Running = <-RunningChan
	Wg.Wait()
}
