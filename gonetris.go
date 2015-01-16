package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/nsf/termbox-go"
	"os"
)

var opts struct {
	Name    string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players int    `short:"p" long:"players" description:"Number of players" required:"true"`
}

var (
	quit chan bool = make(chan bool)
)

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

}

func main() {

	termbox.Init()
	defer termbox.Close()

	go HandleTerminal()
	go HandleBoard()
	go HandleKeys()

	<-quit

}
