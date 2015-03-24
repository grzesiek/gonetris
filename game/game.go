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

func (game *game) Play() {

	game.Wg.Add(5)

	tick := tick.New(game.Opts.Interval)
	go tick.Handle(game.Wg)

	multiplayer := multiplayer.New(game.Opts.Players)
	go multiplayer.Handle(game.Wg)

	player := multiplayer.AddPlayer(game.Opts.Nickname, "", 5, 5)
	board := player.Board

	terminal := terminal.New()
	go terminal.Handle(game.Wg)
	go terminal.HandleKeys(game.Wg, game.CloseEvent, board.BrickOperationEvent)

	go board.Handle(game.Wg, tick, terminal)

	<-game.CloseEvent
	terminal.CloseEvent <- true
	board.CloseEvent <- true
	tick.CloseEvent <- true

	game.Wg.Wait()

}
