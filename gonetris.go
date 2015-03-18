package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"sync"
)

var Opts struct {
	Name     string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players  int    `short:"p" long:"players" description:"Number of players" required:"true"`
	Interval int    `short:"i" long:"interval" description:"Step-down interval in miliseconds" required:"false" default:"400"`
}

var (
	GameClose = make(chan bool)
	Wg        sync.WaitGroup
)

func init() {

	_, err := flags.Parse(&Opts)
	if err != nil {
		fmt.Println("Invalid flags !")
		os.Exit(1)
	}
}

func main() {

	Wg.Add(5)

	go HandleTerminal()
	go HandleKeys()
	go HandlePlayers()
	go HandleBoard()
	go HandleTick()

	<-GameClose
	TerminalClose <- true
	BoardClose <- true
	TickClose <- true

	Wg.Wait()
}
