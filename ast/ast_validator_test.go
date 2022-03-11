package ast

import (
	"errors"
	"testing"
)

func buildEl(t ElType, children ...El) El {
	e := El{
		Type:     t,
		Children: children,
	}
	return e
}

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		in El
	}{
		{
			in: buildEl(TypeElLayout, buildEl(TypeElText)),
		},
		{
			in: buildEl(TypeElLayout,
				buildEl(TypeElColumn,
					buildEl(TypeElImage),
					buildEl(TypeElEl, buildEl(TypeElText)),
				),
			),
		},
	}

	for i, tt := range tests {
		err := Validate(tt.in)
		if err != nil {
			t.Fatalf("[%d] failed : %v", i, err)
		}
	}
}
func TestValidateError(t *testing.T) {
	tests := map[int]struct {
		in  El
		exp error
	}{
		0: {
			in:  El{},
			exp: errInvalidRootType,
		},
		1: {
			in:  buildEl(TypeElLayout, buildEl(TypeElEl), buildEl(TypeElEl)),
			exp: errInvalidChildrenLen,
		},
		2: {
			in:  buildEl(TypeElLayout, El{}),
			exp: errInvalidChildType,
		},
		3: {
			in:  buildEl(TypeElLayout, buildEl(TypeElEl, El{})),
			exp: errInvalidChildType,
		},
		4: {
			in:  buildEl(TypeElLayout, buildEl(TypeElEl, buildEl(TypeElImage, El{}))),
			exp: errInvalidChildrenLen,
		},
		5: {
			in:  buildEl(TypeElLayout, buildEl(TypeElColumn, buildEl(TypeElText))),
			exp: errInvalidChildType,
		},
		6: {
			in:  buildEl(TypeElLayout, buildEl(TypeElEl, buildEl(TypeElText, El{}))),
			exp: errInvalidChildrenLen,
		},
	}

	for i, tt := range tests {
		err := Validate(tt.in)
		if err == nil {
			t.Fatalf("[%d] failed : exp '%s' got nil", i, tt.exp)
		}
		if !errors.Is(err, tt.exp) {
			t.Fatalf("[%d] failed : exp '%s', got '%v'", i, tt.exp, err)
		}
	}
}
