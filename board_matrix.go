package main

type BoardCell struct {
	Color    Color
	Empty    bool
	Embedded bool
}

type BoardMatrix [10][20]BoardCell

func (matrix *BoardMatrix) ResetEmptyCells() {

	for x, cells := range matrix {
		for y, cell := range cells {
			if cell.Embedded == false {
				matrix.ResetCell(x, y)
			}
		}
	}
}

func (matrix *BoardMatrix) ResetCell(x, y int) {

	matrix[x][y].Empty = true
	matrix[x][y].Embedded = false
	matrix[x][y].Color = ColorBlack
}
