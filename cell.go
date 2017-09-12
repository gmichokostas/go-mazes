package main

import (
	"fmt"
)

// Cell is a member of the Grid
type Cell struct {
	row   int
	col   int
	links map[*Cell]bool
	north *Cell
	south *Cell
	east  *Cell
	west  *Cell
}

// NewCell creates a new Cell on the Grid
func NewCell(row, col int) *Cell {
	return &Cell{
		row:   row,
		col:   col,
		links: make(map[*Cell]bool),
	}
}

// Link the Cell with another Cell
func (c *Cell) Link(cell *Cell, bidi bool) {
	c.links[cell] = true
	if bidi == true {
		cell.Link(c, false)
	}
}

// Unlink the Cell from the other Cell
func (c *Cell) Unlink(cell *Cell, bidi bool) {
	delete(c.links, cell)
	if bidi == true {
		cell.Unlink(c, false)
	}
}

// Links lists all Cells connected to this Cell
func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, len(c.links))

	i := 0
	for k := range c.links {
		keys[i] = k
		i++
	}

	return keys
}

// IsLinked checks if the current Cell is linked with the other
func (c *Cell) IsLinked(cell *Cell) bool {
	if _, ok := c.links[cell]; ok {
		return true
	}

	return false
}

// Neighbors return a list of the adjoin cells
func (c *Cell) Neighbors() []*Cell {
	cells := make([]*Cell, 0, 4)

	if c.north != nil {
		cells = append(cells, c.north)
	}

	if c.south != nil {
		cells = append(cells, c.south)
	}

	if c.east != nil {
		cells = append(cells, c.east)
	}

	if c.west != nil {
		cells = append(cells, c.west)
	}

	return cells
}

// String representation of Cell
func (c *Cell) String() string {
	return fmt.Sprintf("[%d %d]", c.row, c.col)
}

// Distances applies Dijkstra's algorithm to calculate
// the distance of a cell from a reference cell
func (c *Cell) Distances() Distances {
	distances := NewDistances(c)
	frontier := []*Cell{c}

	for {
		if len(frontier) == 0 {
			break
		}

		var newFrontier []*Cell
		for _, cell := range frontier {
			for _, linked := range cell.Links() {

				if distances.Contains(linked) == true {
					continue
				}

				distances.SetDistance(linked, distances.GetDistance(cell)+1)
				newFrontier = append(newFrontier, linked)
			}
		}

		frontier = newFrontier
	}
	return distances
}
