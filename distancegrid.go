package main

import (
	"strconv"
)

// DistanceGrid contains the root cell
type DistanceGrid struct {
	rootCell *Cell
}

// NewDistanceGrid creates a new DistanceGrid
// counting the distances from the root cell
func NewDistanceGrid(rootcell *Cell) DistanceGrid {
	return DistanceGrid{rootCell: rootcell}
}

// Render renders the distance of the given cell from the root cell
func (g DistanceGrid) Render(currentcell *Cell) string {
	return " " + strconv.FormatInt(int64(currentcell.Distances().GetDistance(g.rootCell)), 36) + " "
}
