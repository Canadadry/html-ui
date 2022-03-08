package ast

import (
	"testing"
)

func TestParseNumberAttr(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "234",
			out: 234,
		},
	}

	for i, tt := range tests {
		result, err := ParseNumberAttr(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if result != tt.out {
			t.Fatalf("[%d] failed got '%v' exp '%v'", i, result, tt.out)
		}
	}
}
