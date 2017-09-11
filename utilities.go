package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var (
	bgColor   = color.White
	lineColor = color.Black
	cellSize  = 10
)

// ToImage creates a PNG image of the given Grid
func ToImage(grid *Grid, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, (cellSize*grid.columns + 1), (cellSize*grid.rows + 1)))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	// draw each cell's wall
	for cell := range grid.EachCell() {
		x1 := cell.col * cellSize
		y1 := cell.row * cellSize
		x2 := (cell.col + 1) * cellSize
		y2 := (cell.row + 1) * cellSize

		if cell.north == nil {
			// top horizontal line
			horLine(x1, x2, y1, img)
		}

		if cell.west == nil {
			// right vertical line
			verLine(y1, y2, x1, img)
		}

		if cell.IsLinked(cell.east) != true {
			// left vertical line
			verLine(y1, y2, x2, img)
		}

		if cell.IsLinked(cell.south) != true {
			// bottom horizontal line
			horLine(x1, x2, y2, img)
		}
	}

	f, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

// horLine draws a horizontal line from minX to maxX
func horLine(minX, maxX, y int, img *image.RGBA) {
	for ; minX <= maxX; minX++ {
		img.Set(minX, y, lineColor)
	}
}

// verLine draws a vertical line from minY to maxY
func verLine(minY, maxY, x int, img *image.RGBA) {
	for ; minY <= maxY; minY++ {
		img.Set(x, minY, lineColor)
	}
}
