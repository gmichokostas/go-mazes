package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	rows    = flag.Int("rows", 10, "rows of the grid")
	columns = flag.Int("columns", 10, "columns of the grid")
)

func main() {
	flag.Parse()

	grid, err := NewGrid(*rows, *columns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error create new grid: %v\n", err)
		os.Exit(2)
	}

	// BTreeOn(grid)
	SideWinderOn(grid)

	distances := grid.Cell(0, 0).Distances()
	distanceGrid := NewDistanceGrid(distances)

	ds := distances.PathTo(grid.Cell(grid.rows-1, 0))
	distanced := NewDistanceGrid(ds)

	distancePrinter := NewGridPrinter(grid, distanceGrid)
	fmt.Println(distancePrinter.PrintGrid())

	distancedPrinter := NewGridPrinter(grid, distanced)
	fmt.Println(distancedPrinter.PrintGrid())

	ToImage(grid, "out")

}
