package maze

// Distances records the distance
// of each cell from the root cell
type Distances struct {
	root  *Cell
	cells map[*Cell]int
}

// NewDistances takes a Cell and return a new Distances
// initialized with the root cell to the given cell
func NewDistances(cell *Cell) Distances {
	return Distances{
		cells: map[*Cell]int{cell: 0},
		root:  cell,
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

// PathTo finds the path between to cells
func (d Distances) PathTo(goal *Cell) Distances {
	current := goal
	breadcrumbs := NewDistances(d.root)
	breadcrumbs.SetDistance(current, d.GetDistance(current))

	for {
		if current == d.root {
			break
		}

		for _, neighbor := range current.Links() {
			if d.GetDistance(neighbor) < d.GetDistance(current) {
				breadcrumbs.SetDistance(neighbor, d.GetDistance(neighbor))
				current = neighbor
				break
			}
		}
	}

	return breadcrumbs
}

// Max finds the cell that has the longest distance from the root cell
// returns the cell and the distance
func (d Distances) Max() (maxCell *Cell, maxDistance int) {
	maxDistance = 0
	maxCell = d.root

	for cell, distance := range d.cells {
		if distance > maxDistance {
			maxCell = cell
			maxDistance = distance
		}
	}
	return
}
