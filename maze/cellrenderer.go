package maze

// CellRenderer is an interface that supports
// the rendering of a given cell
type CellRenderer interface {
	// Render renders the content of the currentcell
	Render(currentcell *Cell) string
}
