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
				matrix.resetCell(x, y)
			}
		}
	}
}

func (matrix *BoardMatrix) RemoveFullLines() int {

	var lineFull bool
	var removedLines []int

	for by := 0; by < len(matrix[0]); by++ {
		lineFull = true
		for bx := 0; bx < len(matrix); bx++ {
			lineFull = lineFull && matrix[bx][by].Embedded
		}

		if lineFull {
			for bx := 0; bx < len(matrix); bx++ {
				matrix.resetCell(bx, by)
			}
			removedLines = append(removedLines, by)
		}
	}

	if len(removedLines) > 0 {
		for _, y := range removedLines {
			for by := y - 1; by > 0; by-- {
				for bx := 0; bx < len(matrix); bx++ {
					matrix[bx][by+1] = matrix[bx][by]
					matrix.resetCell(bx, by)
				}
			}
		}
		matrix.ResetEmptyCells()
	}

	return len(removedLines)
}

func (matrix *BoardMatrix) resetCell(x, y int) {

	matrix[x][y].Empty = true
	matrix[x][y].Embedded = false
	matrix[x][y].Color = ColorBlack
}
