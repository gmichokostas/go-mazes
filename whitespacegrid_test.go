package main

import "testing"

func TestWhiteSpaceGrid_Render(t *testing.T) {
	type args struct {
		currentcell *Cell
	}

	cell := NewCell(0, 0)

	tests := []struct {
		name string
		g    WhiteSpaceGrid
		args args
		want string
	}{
		{
			name: "returns whitespaces as the content of the cell",
			g:    WhiteSpaceGrid{},
			args: args{currentcell: cell},
			want: "   ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := WhiteSpaceGrid{}
			if got := g.Render(tt.args.currentcell); got != tt.want {
				t.Errorf("WhiteSpaceGrid.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
