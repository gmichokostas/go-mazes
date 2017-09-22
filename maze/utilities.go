package maze

import (
	"image"
	"image/color"
)

// HorLine draws a horizontal line from minX to maxX
func HorLine(minX, maxX, y int, img *image.RGBA) {
	for ; minX <= maxX; minX++ {
		img.Set(minX, y, color.Black)
	}
}

// VerLine draws a vertical line from minY to maxY
func VerLine(minY, maxY, x int, img *image.RGBA) {
	for ; minY <= maxY; minY++ {
		img.Set(x, minY, color.Black)
	}
}
