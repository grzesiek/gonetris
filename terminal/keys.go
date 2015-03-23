package terminal

func (terminal *terminal) HandleKeys(gameCloseEvent chan bool, brickOperationEvent chan string) {

	defer Wg.Done()

	for {
		if event := termbox.PollEvent(); event.Type == termbox.EventKey {

			switch event.Ch {
			case 'p': /*	Pause  					 */
			case 'q': /*	Quit						 */
				gameCloseEvent <- true
				return
			case 'j': /*	Move brick left */
				brickOperationEvent <- "BrickMoveLeft"
			case 'l': /*	Move brick right */
				brickOperationEvent <- "BrickMoveRight"
			case 'k': /*  Rotate brick */
				brickOperationEvent <- "BrickRotate"
			case 'm': /*  Move down brick */
				brickOperationEvent <- "BrickMoveDown"
			}

			switch event.Key {
			case termbox.KeySpace: /*  Drop brick */
				brickOperationEvent <- "BrickDrop"
			}

		}
	}
}
