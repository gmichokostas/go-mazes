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

	// maze.BTreeOn(grid)
	maze.SideWinderOn(grid)

	distances := grid.Cell(0, 0).Distances()
	img := maze.ImageableGrid{Grid: grid, Distances: &distances}
	img.ToImage("out")
}
