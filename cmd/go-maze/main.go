package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gmichokostas/go-mazes/maze"
)

var (
	rows    = flag.Int("rows", 10, "rows of the grid")
	columns = flag.Int("columns", 10, "columns of the grid")
)

func main() {
	flag.Parse()

	grid, err := maze.NewGrid(*rows, *columns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error create new grid: %v\n", err)
		os.Exit(2)
	}

	maze.BTreeOn(grid)
	// SideWinderOn(grid)

	distances := grid.Cell(0, 0).Distances()
	img := maze.ImageableGrid{Grid: grid, Distances: &distances}
	img.ToImage("out")

	// distances := grid.Cell(0, 0).Distances()
	// distanceGrid := NewDistanceGrid(distances)

	// distancePrinter := NewGridPrinter(grid, distanceGrid)
	// fmt.Println(distancePrinter.PrintGrid())

	// ds := distances.PathTo(grid.Cell(grid.rows-1, 0))
	// distanced := NewDistanceGrid(ds)

	// distancedPrinter := NewGridPrinter(grid, distanced)
	// fmt.Println(distancedPrinter.PrintGrid())

	// start := distanceGrid.rootCell
	// newDistances := start.Distances()
	// newStart, _ := newDistances.Max()

	// newDistances = newStart.Distances()
	// goal, _ := newDistances.Max()

	// maxDis := newDistances.PathTo(goal)
	// max := NewDistanceGrid(maxDis)

	// mx := NewGridPrinter(grid, max)
	// fmt.Println(mx.PrintGrid())

	// ToImage(grid, "out")
}
