package game

import (
	"fmt"
	"os"
	"sync"

	"github.com/jessevdk/go-flags"

	"github.com/grzesiek/gonetris/multiplayer"
	"github.com/grzesiek/gonetris/terminal"
	"github.com/grzesiek/gonetris/tick"
)

type opts struct {
	Nickname string `short:"n" long:"nick" description:"Your nickname in game" required:"true"`
	Players  int    `short:"p" long:"players" description:"Number of players" required:"true"`
	Interval int    `short:"i" long:"interval" description:"Step-down interval in miliseconds" required:"false" default:"400"`
}

type game struct {
	Wg         *sync.WaitGroup
	Opts       opts
	CloseEvent chan bool
}

func NewGame() *game {

	var wg sync.WaitGroup
	g := game{Wg: &wg, CloseEvent: make(chan bool)}

	_, err := flags.Parse(&g.Opts)
	if err != nil {
		fmt.Println("Invalid flags !")
		os.Exit(1)
	}

	return &g
}

func (game *game) Play() {

	game.Wg.Add(5)

	multiplayer := multiplayer.New(game.Opts.Players)
	go multiplayer.Handle(game.Wg)
	player := multiplayer.AddPlayer(game.Opts.Nickname, "", 5, 5)

	terminal := terminal.New()
	go terminal.Handle(game.Wg)
	go terminal.HandleKeys(game.Wg, game.CloseEvent, player.Board.BrickOperationEvent)

	tick := tick.New(game.Opts.Interval)
	go tick.Handle(game.Wg)

	go player.Board.Handle(game.Wg, tick, terminal)

	<-game.CloseEvent
	player.Board.CloseEvent <- true
	terminal.CloseEvent <- true
	tick.CloseEvent <- true

	game.Wg.Wait()

}
