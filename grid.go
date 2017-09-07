package main

import (
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
	if err := grid.prepare(); err != nil {
		return nil, errors.New("Error while preparing the grid: " + err.Error())
	}
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

// prepare creates a 2D slice of Cells
func (g *Grid) prepare() error {
	g.structure = make([][]*Cell, g.rows)
	for row := range g.structure {
		g.structure[row] = make([]*Cell, g.columns)
	}

	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.columns; j++ {

			cell, err := NewCell(i, j)
			if err != nil {
				return err
			}
			g.structure[i][j] = cell
		}
	}
	return nil
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

// eachRow returns a chan with eachRow of the grid
func (g *Grid) eachRow() chan []*Cell {
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
		for row := range g.eachRow() {
			for _, cell := range row {
				c <- cell
			}
		}
		close(c)
	}()

	return c
}
