package main

import (
	"github.com/jessevdk/go-flags"
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

func main() {

	go HandleBoard()
	time.Sleep(time.Duration(8) * time.Second)

}
