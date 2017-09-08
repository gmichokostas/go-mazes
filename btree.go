package main

import (
	"math/rand"
	"time"
)

// BTree contains the BTree algorithm
type BTree struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// On applies the BTree algorithm on the given Grid
func (bt BTree) On(grid *Grid) *Grid {

	for cell := range grid.EachCell() {
		neighbors := make([]*Cell, 0, 10)

		if cell.north != nil {
			neighbors = append(neighbors, cell.north)
		}

		if cell.east != nil {
			neighbors = append(neighbors, cell.east)
		}

		if len(neighbors) == 0 {
			continue
		}

		index := rand.Intn(len(neighbors))
		neighbor := neighbors[index]

		if neighbor != nil {
			cell.Link(neighbor, true)
		}
	}

	return grid
}
