package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// BTreeOn applies the BTree algorithm on the given Grid
func BTreeOn(grid *Grid) {
	for cell := range grid.EachCell() {
		var neighbors []*Cell

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
}
