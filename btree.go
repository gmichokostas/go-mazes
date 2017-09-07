package main

import "math/rand"

// BTree contains the BTree algorithm
type BTree struct{}

// On applies the BTree algorithm on the given Grid
func (bt BTree) On(grid *Grid) *Grid {
	var neighbors []*Cell

	for cell := range grid.EachCell() {
		if cell.north != nil {
			neighbors = append(neighbors, cell.north)
		}

		if cell.east != nil {
			neighbors = append(neighbors, cell.east)
		}

		index := rand.Intn(len(neighbors))
		neighbor := neighbors[index]

		if neighbor != nil {
			cell.Link(neighbor, true)
		}
	}

	return grid
}
