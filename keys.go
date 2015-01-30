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
				Quit()
			case 'j': /*	Move brick left */
				BoardEvent <- BrickMoveLeft
			case 'l': /*	Move brick right */
				BoardEvent <- BrickMoveRight
			case 'k': /*  Rotate brick */
				BoardEvent <- BrickRotate
			}
		}
	}
}
