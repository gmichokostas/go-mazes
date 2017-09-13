package maze

import (
	"reflect"
	"testing"
)

func TestNewGrid(t *testing.T) {
	type args struct {
		rows    int
		columns int
	}

	testGrid, _ := NewGrid(4, 4)
	tests := []struct {
		name    string
		args    args
		want    *Grid
		wantErr bool
	}{
		{
			name:    "valid rows and columns",
			args:    args{rows: 4, columns: 4},
			want:    testGrid,
			wantErr: false,
		},
		{
			name:    "invalid rows and valid columns",
			args:    args{rows: -4, columns: 4},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid rows and invalid columns",
			args:    args{rows: 4, columns: -4},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGrid(tt.args.rows, tt.args.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGrid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_Size(t *testing.T) {
	type fields struct {
		rows      int
		columns   int
		structure [][]*Cell
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "returns the size of the grid",
			fields: fields{rows: 4, columns: 4, structure: [][]*Cell{}},
			want:   16,
		},
		{
			name:   "returns zero",
			fields: fields{rows: 0, columns: 0, structure: [][]*Cell{}},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				rows:      tt.fields.rows,
				columns:   tt.fields.columns,
				structure: tt.fields.structure,
			}
			if got := g.Size(); got != tt.want {
				t.Errorf("Grid.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_RandomCell(t *testing.T) {
	type fields struct {
		rows      int
		columns   int
		structure [][]*Cell
	}

	cell := NewCell(0, 0)

	tests := []struct {
		name   string
		fields fields
		want   *Cell
	}{
		{
			name:   "returns a random cell from the grid",
			fields: fields{rows: 1, columns: 1, structure: [][]*Cell{{cell}}},
			want:   cell,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				rows:      tt.fields.rows,
				columns:   tt.fields.columns,
				structure: tt.fields.structure,
			}
			if got := g.RandomCell(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.RandomCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_cell(t *testing.T) {
	type fields struct {
		rows      int
		columns   int
		structure [][]*Cell
	}
	type args struct {
		row int
		col int
	}

	cell := NewCell(0, 0)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Cell
	}{
		{
			name:   "returns a cell from the grid at col - row",
			fields: fields{rows: 2, columns: 2, structure: [][]*Cell{{cell}}},
			want:   cell,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				rows:      tt.fields.rows,
				columns:   tt.fields.columns,
				structure: tt.fields.structure,
			}
			if got := g.Cell(tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grid.cell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_prepare(t *testing.T) {
	type fields struct {
		rows    int
		columns int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "it has the right number of rows",
			fields: fields{rows: 2, columns: 2},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grid{
				rows:    tt.fields.rows,
				columns: tt.fields.columns,
			}
			g.prepare()
			if got := len(g.structure); got != tt.want {
				t.Errorf("g.stucture length = %v, want %v", got, tt.want)
			}
		})
	}
}
