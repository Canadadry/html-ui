package ast

import (
	"testing"
)

func TestValidateAttributeSuccess(t *testing.T) {
	tests := map[int]struct {
		in El
	}{
		0: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(
					TypeElColumn,
					[]AttrType{},
					buildElWidthAttr(
						TypeElRow,
						[]AttrType{},
						buildElWidthAttr(
							TypeElEl,
							[]AttrType{},
							buildElWidthAttr(
								TypeElText,
								[]AttrType{},
							),
						),
						buildElWidthAttr(
							TypeElImage,
							[]AttrType{},
						),
					),
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

func buildElWidthAttr(t ElType, attrs []AttrType, children ...El) El {
	e := El{
		Type:     t,
		Attr:     []Attribute{},
		Children: children,
	}
	for _, attr := range attrs {
		e.Attr = append(e.Attr, Attribute{Type: attr})
	}
	return e
}

func TestValidateAttributeError(t *testing.T) {
	tests := map[int]struct {
		in  El
		exp string
	}{
		0: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{TypeAttrWidth},
				buildElWidthAttr(TypeElColumn, nil),
			),
			exp: "layout cannot have attribute 'width' possibilities are []",
		},
		1: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{TypeAttrWidth},
				buildElWidthAttr(TypeElColumn, []AttrType{"fake"}),
			),
			exp: "layout cannot have attribute 'width' possibilities are []",
		},
	}

	for i, tt := range tests {
		err := Validate(tt.in)
		if err == nil {
			t.Fatalf("[%d] failed : exp '%s' got nil", i, tt.exp)
		}
		if err.Error() != tt.exp {
			t.Fatalf("[%d] failed : exp '%s', got '%v'", i, tt.exp, err)
		}
	}
}
