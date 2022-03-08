package ast

import (
	"testing"
)

func TestParseAlignAttr(t *testing.T) {
	tests := []struct {
		in   string
		outX AlignXType
		outY AlignYType
	}{
		{
			in:   "left,right,centerX,top,bottom,centerY",
			outX: AlignCenterX,
			outY: AlignCenterY,
		},
	}

	for i, tt := range tests {
		x, y, err := ParseAlignAttr(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if x != tt.outX {
			t.Fatalf("[%d] x failed got '%v' exp '%v'", i, x, tt.outX)
		}
		if y != tt.outY {
			t.Fatalf("[%d] y failed got '%v' exp '%v'", i, y, tt.outY)
		}
	}
}
