package ast

import (
	"errors"
	"testing"
)

func TestParsePositionAttr(t *testing.T) {
	tests := []struct {
		in  string
		err error
	}{
		{
			in:  "234",
			err: ErrInvalidPosition,
		},
		{in: positonLeft},
		{in: positonRight},
		{in: positonAbove},
		{in: positonBelow},
	}

	for i, tt := range tests {
		err := ParsePositionAttr(tt.in)
		if !errors.Is(err, tt.err) {
			t.Fatalf("[%d] failed got '%v' exp '%v'", i, err, tt.err)
		}
	}
}
