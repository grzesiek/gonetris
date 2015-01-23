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
	Running = true
	Wg      sync.WaitGroup
)

func init() {

	_, err := flags.Parse(&opts)
	if err != nil {
		panic("Invalid flags !")
	}

}

func main() {

	Wg.Add(4)

	go HandleTerminal()
	go HandleBoards()
	go HandleKeys()
	go HandleGame()

	Wg.Wait()

}

func Quit() {

	close(TerminalEvent)
	close(BoardEvent)
}
