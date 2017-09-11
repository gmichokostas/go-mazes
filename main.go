package main

import (
	"fmt"
)

func main() {
	grid, err := NewGrid(10, 10)
	if err != nil {
		fmt.Printf("Error create new grid: %v", err)
	}

	BTreeOn(grid)
	ToImage(grid, "out")

	fmt.Println(grid)
}
