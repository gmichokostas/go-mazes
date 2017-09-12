package main

import "bytes"

// GridPrinter contains the Grid for printing
// the type of the grid
type GridPrinter struct {
	Grid     *Grid
	GridType CellRenderer
}

// NewGridPrinter creates a pointer to a GridPrinter
func NewGridPrinter(grid *Grid, gridType CellRenderer) *GridPrinter {
	return &GridPrinter{Grid: grid, GridType: gridType}
}

// PrintGrid prints the grid
func (gp *GridPrinter) PrintGrid() string {
	var output bytes.Buffer
	for column := 0; column < gp.Grid.columns; column++ {
		output.WriteString("+---")
	}
	output.WriteString("+\n")

	for row := 0; row < gp.Grid.rows; row++ {
		var bottom bytes.Buffer
		bottom.WriteString("+")

		for column := 0; column < gp.Grid.columns; column++ {
			cell := gp.Grid.Cell(row, column)

			if column == 0 {
				if cell.IsLinked(cell.west) == true {
					output.WriteString(" ")
				} else {
					output.WriteString("|")
				}
			}

			output.WriteString(gp.GridType.Render(cell))
			if cell.IsLinked(cell.east) == true {
				output.WriteString(" ")
			} else {
				output.WriteString("|")
			}

			if cell.IsLinked(cell.south) == true {
				bottom.WriteString("   ")
			} else {
				bottom.WriteString("---")
			}
			bottom.WriteString("+")
		}

		output.WriteString("\n")
		output.WriteString(bottom.String())
		output.WriteString("\n")
	}

	return output.String()
}
