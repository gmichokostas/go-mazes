package main

import (
	"fmt"
)

func main() {
	var btree BTree

	grid, err := NewGrid(5, 5)
	if err != nil {
		fmt.Printf("Error create new grid: %v", err)
	}
	btree.On(grid)
	ToImage(grid, "out")

	fmt.Println(grid)
}
