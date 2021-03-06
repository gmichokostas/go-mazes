package maze

import (
	"reflect"
	"testing"
)

func TestNewCell(t *testing.T) {
	type args struct {
		row   int
		col   int
		links map[*Cell]bool
	}
	tests := []struct {
		name string
		args args
		want *Cell
	}{
		{
			name: "valid row and column",
			args: args{col: 1, row: 1, links: map[*Cell]bool{}},
			want: &Cell{row: 1, col: 1, links: map[*Cell]bool{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCell(tt.args.row, tt.args.col)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Links(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
	}

	otherCell := NewCell(1, 2)
	tests := []struct {
		name   string
		fields fields
		want   []*Cell
	}{
		{
			name:   "includes the other cell",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{otherCell: true}},
			want:   []*Cell{otherCell},
		},
		{
			name:   "does not include the other cell",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{}},
			want:   []*Cell{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
			}
			if got := c.Links(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cell.Links() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_IsLinked(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
	}
	type args struct {
		cell *Cell
	}

	otherCell := NewCell(1, 2)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "true if two cells are linked together",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{otherCell: true}},
			args:   args{cell: otherCell},
			want:   true,
		},
		{
			name:   "false if two cells are not linked together",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{}},
			args:   args{cell: otherCell},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
			}
			if got := c.IsLinked(tt.args.cell); got != tt.want {
				t.Errorf("Cell.IsLinked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Neighbors(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
		north *Cell
		south *Cell
		east  *Cell
		west  *Cell
	}

	north := NewCell(0, 1)
	south := NewCell(2, 1)
	east := NewCell(1, 2)
	west := NewCell(1, 0)

	tests := []struct {
		name   string
		fields fields
		want   []*Cell
	}{
		{
			name: "returns a slice of the adjoin cells",
			fields: fields{
				row:   1,
				col:   1,
				links: map[*Cell]bool{},
				north: north,
				south: south,
				east:  east,
				west:  west,
			},
			want: []*Cell{north, south, east, west},
		},
		{
			name: "returns a slice of the adjoin cells",
			fields: fields{
				row:   1,
				col:   1,
				links: map[*Cell]bool{},
				south: south,
				east:  east,
			},
			want: []*Cell{south, east},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
				north: tt.fields.north,
				south: tt.fields.south,
				east:  tt.fields.east,
				west:  tt.fields.west,
			}
			if got := c.Neighbors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cell.Neighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Link(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
	}
	type args struct {
		cell *Cell
		bidi bool
	}

	otherCell := NewCell(1, 2)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "create a link with the other cell",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{}},
			args:   args{cell: otherCell, bidi: true},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
			}
			c.Link(tt.args.cell, tt.args.bidi)

			if got := c.IsLinked(otherCell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("c.IsLinked(otherCell) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Unlink(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
	}
	type args struct {
		cell *Cell
		bidi bool
	}

	otherCell := NewCell(1, 2)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "deletes a link with the other cell",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{otherCell: true}},
			args:   args{cell: otherCell, bidi: true},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
			}
			c.Unlink(tt.args.cell, tt.args.bidi)

			if got := c.IsLinked(otherCell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("c.IsLinked(otherCell) = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_String(t *testing.T) {
	type fields struct {
		row   int
		col   int
		links map[*Cell]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "the string representation of the cell contains the [row, col]",
			fields: fields{row: 1, col: 1, links: map[*Cell]bool{}},
			want:   "[1 1]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				row:   tt.fields.row,
				col:   tt.fields.col,
				links: tt.fields.links,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Cell.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
