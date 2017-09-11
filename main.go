package main

import (
	"flag"
	"fmt"
)

var (
	rows    = flag.Int("rows", 10, "rows of the grid")
	columns = flag.Int("columns", 10, "columns of the grid")
)

func main() {
	flag.Parse()

	grid, err := NewGrid(*rows, *columns)
	if err != nil {
		fmt.Printf("Error create new grid: %v", err)
	}

	// BTreeOn(grid)
	SideWinderOn(grid)
	ToImage(grid, "out")

	fmt.Println(grid)
}
