package main

import (
	"fmt"
)

func main() {
	var btree BTree

	grid, err := NewGrid(4, 4)
	if err != nil {
		fmt.Printf("Error create new grid: %v", err)
	}

	btree.On(grid)
}
