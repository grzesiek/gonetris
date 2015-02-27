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
				BoardBrickOperation <- "BrickMoveLeft"
			case 'l': /*	Move brick right */
				BoardBrickOperation <- "BrickMoveRight"
			case 'k': /*  Rotate brick */
				BoardBrickOperation <- "BrickRotate"
			case 'm': /*  Move down brick */
				BoardBrickOperation <- "BrickMoveDown"
			}

			switch event.Key {
			case termbox.KeySpace: /*  Drop brick */
				BoardBrickOperation <- "BrickDrop"
			}

		}
	}
}
