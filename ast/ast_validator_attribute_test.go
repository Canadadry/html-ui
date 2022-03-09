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
				[]AttrType{},
				buildElWidthAttr(TypeElColumn, []AttrType{"fake"}),
			),
			exp: "column cannot have attribute 'fake' possibilities are [align bg-color border-color border-rounded border-width font-color font-size height padding spacing width]",
		},
		2: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(TypeElRow, []AttrType{"fake"}),
			),
			exp: "row cannot have attribute 'fake' possibilities are [align bg-color border-color border-rounded border-width font-color font-size height padding spacing width]",
		},
		3: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(TypeElEl, []AttrType{"fake"}),
			),
			exp: "el cannot have attribute 'fake' possibilities are [align bg-color border-color border-rounded border-width font-color font-size height padding spacing width]",
		},
		4: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(TypeElRow, []AttrType{},
					buildElWidthAttr(TypeElImage, []AttrType{"fake"}),
				),
			),
			exp: "img cannot have attribute 'fake' possibilities are [alt height src width]",
		},
		41: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(TypeElRow, []AttrType{},
					buildElWidthAttr(TypeElEl, []AttrType{}),
					buildElWidthAttr(TypeElImage, []AttrType{TypeAttrPadding}),
				),
			),
			exp: "img cannot have attribute 'padding' possibilities are [alt height src width]",
		},
		5: {
			in: buildElWidthAttr(
				TypeElLayout,
				[]AttrType{},
				buildElWidthAttr(TypeElEl, []AttrType{},
					buildElWidthAttr(TypeElText, []AttrType{"fake"}),
				),
			),
			exp: "text cannot have attribute 'fake' possibilities are []",
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
