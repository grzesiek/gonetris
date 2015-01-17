package main

import (
	"github.com/nsf/termbox-go"
)

func HandleKeys() {

	defer Wg.Done()

	for Running {

		if event := termbox.PollEvent(); event.Type == termbox.EventKey {

			switch event.Ch {
			case 'q': /*		Quit		*/
				Running = false
			}
		}
	}

}
