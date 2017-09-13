package main

import (
	"strconv"
)

// DistanceGrid contains the root cell
type DistanceGrid struct {
	rootCell  *Cell
	distances Distances
}

// NewDistanceGrid creates a new DistanceGrid
// counting the distances from the root cell
func NewDistanceGrid(distances Distances) DistanceGrid {
	return DistanceGrid{rootCell: distances.root, distances: distances}
}

// Render renders the distance of the given cell from the root cell
func (g DistanceGrid) Render(currentcell *Cell) string {
	distance := " "
	if g.rootCell == currentcell || g.distances.GetDistance(currentcell) != 0 {
		distance = strconv.FormatInt(int64(g.distances.GetDistance(currentcell)), 36)
	}

	return " " + distance + " "
}
