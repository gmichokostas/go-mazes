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

	rootCell := grid.Cell(0, 0)
	distanceGrid := NewDistanceGrid(rootCell)
	distancePrinter := NewGridPrinter(grid, distanceGrid)
	fmt.Println(distancePrinter.PrintGrid())

	whitespacePrinter := NewGridPrinter(grid, WhiteSpaceGrid{})
	fmt.Println(whitespacePrinter.PrintGrid())

	ToImage(grid, "out")

}
