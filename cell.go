package mazes

import "errors"

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

// New creates a new Cell on the Grid
func New(row, col int) (*Cell, error) {
	if row < 0 {
		return nil, errors.New("The row must positive number")
	}

	if col < 0 {
		return nil, errors.New("The column must positive number")
	}

	return &Cell{
		row:   row,
		col:   col,
		links: make(map[*Cell]bool),
	}, nil
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
