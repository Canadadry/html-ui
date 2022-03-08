package ast

import (
	"image/color"
	"testing"
)

func TestParseColorAttr(t *testing.T) {
	tests := []struct {
		in  string
		out color.RGBA
	}{
		{
			in:  "rgb(128,255,0)",
			out: color.RGBA{128, 255, 0, 0},
		},
	}

	for i, tt := range tests {
		result, err := ParseColorAttr(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if result != tt.out {
			t.Fatalf("[%d] failed got '%v' exp '%v'", i, result, tt.out)
		}
	}
}
