package main

import (
	"github.com/nsf/termbox-go"
)

func HandleKeys() {

	defer Wg.Done()

	for Running {

		if event := termbox.PollEvent(); event.Type == termbox.EventKey {

			switch event.Ch {
			case 'p': /*	Pause  					 */
				Paused = true
			case 'q': /*	Quit						 */
				RunningChan <- false
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
