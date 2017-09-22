package maze

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

const goRoutines = 15

var (
	bgColor  = color.White
	cellSize = 10
)

// ImageableGrid image representation of a grid
type ImageableGrid struct {
	Grid      *Grid
	Distances *Distances
}

type point struct {
	x, y int
}

type imgCell struct {
	p1, p2 point
	clr    color.Color
	cell   *Cell
	img    *image.RGBA
}

// ToImage creates a colorful representation of the grid
func (imageable ImageableGrid) ToImage(filename string) {
	var wg sync.WaitGroup
	img := image.NewRGBA(image.Rect(0, 0, (cellSize*imageable.Grid.columns + 1), (cellSize*imageable.Grid.rows + 1)))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	tasks := make(chan imgCell, imageable.Grid.Size())

	wg.Add(goRoutines)
	for id := 1; id <= goRoutines; id++ {
		go worker(tasks, &wg)
	}

	for cell := range imageable.Grid.EachCell() {
		p1 := point{x: cell.col * cellSize, y: cell.row * cellSize}
		p2 := point{x: (cell.col + 1) * cellSize, y: (cell.row + 1) * cellSize}
		clr := imageable.ColorForCell(cell)
		imgCl := imgCell{p1: p1, p2: p2, cell: cell, clr: clr, img: img}

		tasks <- imgCl
	}

	close(tasks)
	wg.Wait()

	save(img, filename)
}

func worker(tasks chan imgCell, wg *sync.WaitGroup) {
	defer wg.Done()
	for imgCl := range tasks {
		drawBackground(imgCl)
	}
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

// drawBackground draws the background
func drawBackground(imgCl imgCell) {
	rect := image.NewRGBA(image.Rect(imgCl.p1.x, imgCl.p1.y, imgCl.p2.x, imgCl.p2.y))
	draw.Draw(imgCl.img, rect.Bounds(), &image.Uniform{imgCl.clr}, image.ZP, draw.Src)
}

// drawWalls draws the walls of the maze
func drawWalls(imgCl imgCell) {
	if imgCl.cell.north == nil {
		// top horizontal line
		HorLine(imgCl.p1.x, imgCl.p2.x, imgCl.p1.y, imgCl.img)
	}

	if imgCl.cell.west == nil {
		// right vertical line
		VerLine(imgCl.p1.y, imgCl.p2.y, imgCl.p1.x, imgCl.img)
	}

	if imgCl.cell.IsLinked(imgCl.cell.east) != true {
		// left vertical line
		VerLine(imgCl.p1.y, imgCl.p2.y, imgCl.p2.x, imgCl.img)
	}

	if imgCl.cell.IsLinked(imgCl.cell.south) != true {
		// bottom horizontal line
		HorLine(imgCl.p1.x, imgCl.p2.x, imgCl.p2.y, imgCl.img)
	}
}

func save(img *image.RGBA, filename string) {
	f, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	png.Encode(f, img)
}
