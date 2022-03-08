package ast

import (
	"testing"
)

func TestParseSizeAttr(t *testing.T) {
	tests := []struct {
		in  string
		out AttrSize
	}{
		{
			in:  "fill",
			out: SizeFill{},
		},
		{
			in:  "px:234",
			out: SizePx(234),
		},
		{
			in:  "portion:8",
			out: SizePortion(8),
		},
	}

	for i, tt := range tests {
		result, err := ParseSizeAttr(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if result.Get() != tt.out.Get() {
			t.Fatalf("[%d] failed got '%v' exp '%v'", i, result.Get(), tt.out.Get())
		}
		if result.Type() != tt.out.Type() {
			t.Fatalf("[%d] failed got '%v' exp '%v'", i, result.Type(), tt.out.Type())
		}
	}
}
