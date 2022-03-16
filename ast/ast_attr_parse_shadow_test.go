package ast

import (
	"image/color"
	"testing"
)

func TestParseShadowAttr(t *testing.T) {
	tests := []struct {
		in  string
		out Shadow
	}{
		{
			in:  "ox:1",
			out: Shadow{OffsetX: 1},
		},
		{
			in:  "oy:1",
			out: Shadow{OffsetY: 1},
		},
		{
			in:  "b:1",
			out: Shadow{Blur: 1},
		},
		{
			in:  "s:1",
			out: Shadow{Size: 1},
		},
		{
			in:  "color:white",
			out: Shadow{Color: color.RGBA{255, 255, 255, 0}},
		},
	}

	for i, tt := range tests {
		result, err := ParseShadowAttr(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if result != tt.out {
			t.Fatalf("[%d] x failed got '%v' exp '%v'", i, result, tt.out)
		}
	}
}
