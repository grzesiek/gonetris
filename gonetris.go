package main

import (
	"github.com/jessevdk/go-flags"
	"sync"
)

var opts struct {
	Name     string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players  int    `short:"p" long:"players" description:"Number of players" required:"true"`
        Interval int    `short:"i" long:"interval" description:"Step-down interval in miliseconds" required:"false" default:"400"`
}

var (
	GameClose = make(chan bool)
	Wg        sync.WaitGroup
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
	go HandleTick(opts.Interval)

	<-GameClose
	TerminalClose <- true
	BoardClose <- true
	TickClose <- true

	Wg.Wait()
}
