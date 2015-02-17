package main

import (
	"github.com/nsf/termbox-go"
)

var ()

func HandleKeys() {

	defer Wg.Done()

	for {
		if event := termbox.PollEvent(); event.Type == termbox.EventKey {

			switch event.Ch {
			case 'p': /*	Pause  					 */
			case 'q': /*	Quit						 */
				GameClose <- true
				return
			case 'j': /*	Move brick left */
				BrickOperation <- "MoveLeft"
			case 'l': /*	Move brick right */
				BrickOperation <- "MoveRight"
			case 'k': /*  Rotate brick */
				BrickOperation <- "Rotate"
			}
		}
	}
}
