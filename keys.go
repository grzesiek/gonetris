package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func HandleKeys() {

	for {

		if event := termbox.PollEvent(); event.Type == termbox.EventKey {

			switch event.Ch {
			case 'q': /* Quit */
				fmt.Println("Closing ...")
				quit <- true
			}

		}

		time.Sleep(10 * time.Millisecond)

	}

}
