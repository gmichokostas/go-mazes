package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	backgrounds = iota
	walls
)

var (
	bgColor   = color.White
	lineColor = color.Black
	cellSize  = 10
)

// ImageableGrid image representation of a grid
type ImageableGrid struct {
	Grid      *Grid
	Distances *Distances
}

// ToImage creates a colorful representation of the grid
func (imageable ImageableGrid) ToImage(filename string) {
	img := image.NewRGBA(image.Rect(0, 0, (cellSize*imageable.Grid.columns + 1), (cellSize*imageable.Grid.rows + 1)))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	for _, mode := range []int{backgrounds, walls} {
		for cell := range imageable.Grid.EachCell() {
			x1 := cell.col * cellSize
			y1 := cell.row * cellSize
			x2 := (cell.col + 1) * cellSize
			y2 := (cell.row + 1) * cellSize

			if mode == backgrounds {
				rect := image.NewRGBA(image.Rect(x1, y1, x2, y2))
				color := imageable.ColorForCell(cell)
				draw.Draw(img, rect.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
			} else {
				if cell.north == nil {
					// top horizontal line
					HorLine(x1, x2, y1, img)
				}

				if cell.west == nil {
					// right vertical line
					VerLine(y1, y2, x1, img)
				}

				if cell.IsLinked(cell.east) != true {
					// left vertical line
					VerLine(y1, y2, x2, img)
				}

				if cell.IsLinked(cell.south) != true {
					// bottom horizontal line
					HorLine(x1, x2, y2, img)
				}
			}
		}
	}

	f, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

// ColorForCell return a color based on the cell's distance
func (imageable ImageableGrid) ColorForCell(cell *Cell) color.Color {
	_, max := imageable.Distances.Max()

	if imageable.Distances.Contains(cell) == true {
		distance := imageable.Distances.GetDistance(cell)
		intensity := float64(max-distance) / float64(max)
		dark := uint8(255 * intensity)
		bright := uint8(128 + (127 * intensity))
		return color.RGBA{dark, bright, dark, 255}
	}

	return color.White
}
