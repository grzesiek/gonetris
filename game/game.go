package game

import (
	"fmt"
	"github.com/grzesiek/gonetris/board"
	"github.com/grzesiek/gonetris/multiplayer"
	"github.com/jessevdk/go-flags"
	"os"
	"sync"
)

type opts struct {
	Name     string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players  int    `short:"p" long:"players" description:"Number of players" required:"true"`
	Interval int    `short:"i" long:"interval" description:"Step-down interval in miliseconds" required:"false" default:"400"`
}

type game struct {
	Wg         sync.WaitGroup
	Opts       opts
	CloseEvent chan bool
}

func NewGame() *game {

	g := game{CloseEvent: make(chan bool)}

	_, err := flags.Parse(&g.Opts)
	if err != nil {
		fmt.Println("Invalid flags !")
		os.Exit(1)
	}

	return &g
}

func (game *game) Handle() {

	game.Wg.Add(5)

	multiplayer := multiplayer.New()
	go multiplayer.Handle(game.Wg)

	board := board.New()
	go board.Handle(game.Wg)

	tick := tick.New()
	go tick.Handle(game.Wg, game.Opts.Interval)

	terminal := terminal.New()
	go terminal.Handle(game.Wg)
	go terminal.HandleKeys(game.Wg)

	<-game.CloseEvent
	terminal.CloseEvent <- true
	board.CloseEvent <- true
	tick.CloseEvent <- true

	game.Wg.Wait()

}
