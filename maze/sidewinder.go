package maze

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// SideWinderOn applies the SideWinder
// algorithm on the given grid
func SideWinderOn(grid *Grid) {
	for row := range grid.EachRow() {
		var run []*Cell

		for _, cell := range row {
			run = append(run, cell)

			atEasternBoundary := (cell.east == nil)
			atNorthernBoundary := (cell.north == nil)

			if shouldCloseOut := atEasternBoundary || (!atNorthernBoundary && rand.Intn(2) == 0); shouldCloseOut {
				index := rand.Intn(len(run))
				member := run[index]

				if member.north != nil {
					member.Link(member.north, true)
					run = nil
				}
			} else {
				cell.Link(cell.east, true)
			}
		}
	}
}
