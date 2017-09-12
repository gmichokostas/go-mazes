package main

// WhiteSpaceGrid cells are rendered as whitespaces
type WhiteSpaceGrid struct{}

// Render the given cell as whitespaces
func (g WhiteSpaceGrid) Render(currentcell *Cell) string {
	return "   "
}
