package main

import (
	"bytes"
	"errors"
	"math/rand"
)

// Grid is a container of cells
// keeps track of the number of rows and columns
type Grid struct {
	rows, columns int
	structure     [][]*Cell
}

// NewGrid returns a Grid with the desired dimensions
func NewGrid(rows, columns int) (*Grid, error) {
	if rows <= 0 {
		return nil, errors.New("rows must be greater than zero")
	}

	if columns <= 0 {
		return nil, errors.New("columns must be greater that zero")
	}

	grid := &Grid{rows: rows, columns: columns}
	grid.prepare()
	grid.configureCells()

	return grid, nil
}

// Size returns the size of the grid
func (g *Grid) Size() int {
	return g.rows * g.columns
}

// RandomCell returns a random cell from the grid
func (g *Grid) RandomCell() *Cell {
	row := rand.Intn(g.rows)
	column := rand.Intn(len(g.structure[row]))

	return g.cell(row, column)
}

// String representation of the Grid
func (g *Grid) String() string {
	var output bytes.Buffer

	for column := 0; column < g.columns; column++ {
		output.WriteString("+---")
	}
	output.WriteString("+\n")

	for row := 0; row < g.rows; row++ {
		var bottom bytes.Buffer
		bottom.WriteString("+")

		for column := 0; column < g.columns; column++ {
			cell := g.cell(row, column)

			if column == 0 {
				if cell.IsLinked(cell.west) == true {
					output.WriteString(" ")
				} else {
					output.WriteString("|")
				}
			}

			output.WriteString("   ")
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

// prepare creates a 2D slice of Cells
func (g *Grid) prepare() {
	g.structure = make([][]*Cell, g.rows)
	for row := range g.structure {
		g.structure[row] = make([]*Cell, g.columns)
	}

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {

			cell := NewCell(i, j)
			g.structure[i][j] = cell
		}
	}
}

// configureCells configs each cell of the grid
func (g *Grid) configureCells() {
	for cell := range g.EachCell() {
		row, col := cell.row, cell.col

		cell.north = g.cell(row-1, col)
		cell.south = g.cell(row+1, col)
		cell.east = g.cell(row, col+1)
		cell.west = g.cell(row, col-1)
	}
}

// cell returns the cell at the row - col
func (g *Grid) cell(row, col int) *Cell {
	if row < 0 || row > (g.rows-1) {
		return nil
	}

	if col < 0 || col > (len(g.structure[row])-1) {
		return nil
	}
	return g.structure[row][col]
}

// EachRow returns a chan with EachRow of the grid
func (g *Grid) EachRow() chan []*Cell {
	c := make(chan []*Cell)

	go func() {
		for _, row := range g.structure {
			c <- row
		}
		close(c)
	}()

	return c
}

// EachCell yields each cell of the grid
func (g *Grid) EachCell() chan *Cell {
	c := make(chan *Cell)

	go func() {
		for row := range g.EachRow() {
			for _, cell := range row {
				c <- cell
			}
		}
		close(c)
	}()

	return c
}
