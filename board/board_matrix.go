package board

type cell struct {
	Color    Color
	Empty    bool
	Embedded bool
}

type matrix [10][20]cell

func newMatrix() matrix {

	var m matrix

	for x, cells := range m {
		for y := range cells {
			m.resetCell(x, y)
		}
	}

	return m
}

func (matrix *matrix) resetEmptyCells() {

	for x, cells := range matrix {
		for y, cell := range cells {
			if cell.Embedded == false {
				matrix.resetCell(x, y)
			}
		}
	}
}

func (matrix *matrix) removeFullLines() int {

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
		matrix.resetEmptyCells()
	}

	return len(removedLines)
}

func (matrix *matrix) resetCell(x, y int) {

	matrix[x][y].Empty = true
	matrix[x][y].Embedded = false
	matrix[x][y].Color = ColorBlack
}
