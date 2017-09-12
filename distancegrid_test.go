package main

import (
	"reflect"
	"testing"
)

func TestNewDistanceGrid(t *testing.T) {
	type args struct {
		rootCell *Cell
	}

	root := NewCell(0, 0)

	tests := []struct {
		name string
		args args
		want DistanceGrid
	}{
		{
			name: "returns an initialized DistanceGrid",
			args: args{rootCell: root},
			want: DistanceGrid{rootCell: root},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDistanceGrid(tt.args.rootCell); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDistanceGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceGrid_Render(t *testing.T) {
	type fields struct {
		rootCell *Cell
	}
	type args struct {
		currentcell *Cell
	}
	cell := NewCell(0, 0)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "renders the cell's distance from root",
			fields: fields{rootCell: cell},
			args:   args{currentcell: cell},
			want:   " 0 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := DistanceGrid{
				rootCell: tt.fields.rootCell,
			}
			if got := g.Render(tt.args.currentcell); got != tt.want {
				t.Errorf("DistanceGrid.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
