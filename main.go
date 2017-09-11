package main

import (
	"fmt"
)

func main() {
	grid, err := NewGrid(4, 4)
	if err != nil {
		fmt.Printf("Error create new grid: %v", err)
	}

	// BTreeOn(grid)
	SideWinderOn(grid)
	ToImage(grid, "out")

	fmt.Println(grid)
}
