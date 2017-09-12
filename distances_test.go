package main

import (
	"reflect"
	"testing"
)

func TestNewDistances(t *testing.T) {
	type args struct {
		cell *Cell
	}
	cell := NewCell(0, 0)
	tests := []struct {
		name string
		args args
		want Distances
	}{
		{
			name: "returns a new Distance with the root cell initialized",
			args: args{cell: cell},
			want: Distances{cells: map[*Cell]int{cell: 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDistances(tt.args.cell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistances_GetDistance(t *testing.T) {
	type fields struct {
		cells map[*Cell]int
	}
	type args struct {
		cell *Cell
	}

	root := NewCell(0, 0)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "returns the distance of the given cell from the root cell",
			fields: fields{cells: map[*Cell]int{root: 0}},
			args:   args{cell: root},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Distances{
				cells: tt.fields.cells,
			}
			if got := d.GetDistance(tt.args.cell); got != tt.want {
				t.Errorf("Distances.GetDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistances_SetDistance(t *testing.T) {
	type fields struct {
		cells map[*Cell]int
	}
	type args struct {
		cell     *Cell
		distance int
	}

	cell := NewCell(0, 0)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "set the distance of the given cell from the root cell",
			fields: fields{map[*Cell]int{cell: 0}},
			args:   args{cell: cell, distance: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Distances{
				cells: tt.fields.cells,
			}
			d.SetDistance(tt.args.cell, tt.args.distance)
		})
	}
}

func TestDistances_Contains(t *testing.T) {
	type fields struct {
		cells map[*Cell]int
	}
	type args struct {
		cell *Cell
	}

	cellA := NewCell(0, 0)
	cellB := NewCell(0, 1)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "returns true if the cell exists in the distances",
			fields: fields{map[*Cell]int{cellA: 0}},
			args:   args{cell: cellA},
			want:   true,
		},
		{
			name:   "returns false if the cell doesn't exists in the distances",
			fields: fields{map[*Cell]int{cellA: 0}},
			args:   args{cell: cellB},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Distances{
				cells: tt.fields.cells,
			}
			if got := d.Contains(tt.args.cell); got != tt.want {
				t.Errorf("Distances.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistances_Cells(t *testing.T) {
	type fields struct {
		cells map[*Cell]int
	}

	cell := NewCell(0, 0)

	tests := []struct {
		name   string
		fields fields
		want   []*Cell
	}{
		{
			name:   "returns all the cells",
			fields: fields{cells: map[*Cell]int{cell: 0}},
			want:   []*Cell{cell},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Distances{
				cells: tt.fields.cells,
			}
			if got := d.Cells(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distances.Cells() = %v, want %v", got, tt.want)
			}
		})
	}
}
