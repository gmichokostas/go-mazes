package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// ToImage creates a PNG image of the given Grid
func ToImage(grid *Grid, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, (10 * grid.columns), (10 * grid.rows)))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	for i := img.Bounds().Min.X; i < img.Bounds().Max.X; i++ {
		img.Set(i, 0, color.Black)
		img.Set(0, i, color.Black)
		img.Set(img.Bounds().Max.Y-1, i, color.Black)
		img.Set(i, img.Bounds().Max.Y-1, color.Black)
	}

	// draw each cell's wall
	for cell := range grid.EachCell() {
		x1 := cell.col * 10
		y1 := cell.row * 10
		x2 := (cell.col + 1) * 10
		y2 := (cell.row + 1) * 10

		if cell.north == nil {
			// top horizontal line
			for i := x1; i <= x2; i++ {
				img.Set(i, y1, color.Black)
			}
		}

		if cell.west == nil {
			// right vertical line
			for i := y1; i <= y2; i++ {
				img.Set(x1, i, color.Black)
			}
		}

		if cell.IsLinked(cell.east) != true {
			// left vertical line
			for i := y1; i <= y2; i++ {
				img.Set(x2, i, color.Black)
			}
		}

		if cell.IsLinked(cell.south) != true {
			// bottom horizontal line
			for i := x1; i <= x2; i++ {
				img.Set(i, y2, color.Black)
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
