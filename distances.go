package main

// Distances records the distance
// of each cell from the root cell
type Distances struct {
	cells map[*Cell]int
}

// NewDistances takes a Cell and return a new Distances
// initialized with the root cell to the given cell
func NewDistances(cell *Cell) Distances {
	return Distances{
		cells: map[*Cell]int{cell: 0},
	}
}

// GetDistance of the given cell
func (d Distances) GetDistance(cell *Cell) int {
	return d.cells[cell]
}

// SetDistance of the given cell
func (d Distances) SetDistance(cell *Cell, distance int) {
	d.cells[cell] = distance
}

// Contains checks if the given cell exists
func (d Distances) Contains(cell *Cell) bool {
	if _, ok := d.cells[cell]; ok {
		return true
	}
	return false
}

// Cells returns all the cells
func (d Distances) Cells() []*Cell {
	keys := make([]*Cell, len(d.cells))

	i := 0
	for key := range d.cells {
		keys[i] = key
		i++
	}

	return keys
}
