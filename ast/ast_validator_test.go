package ast

import (
	"testing"
)

func TestValidateSuccess(t *testing.T) {
	tests := []struct {
		in El
	}{
		{
			in: El{
				Type: TypeElLayout,
				Children: []El{
					{
						Type: TypeElText,
					},
				},
			},
		},
		{
			in: El{
				Type: TypeElLayout,
				Children: []El{
					{
						Type: TypeElColumn,
						Children: []El{
							{
								Type: TypeElImage,
							},
							{
								Type: TypeElEl,
								Children: []El{
									{
										Type: TypeElText,
									},
								},
							},
						},
					},
				},
			},
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
		exp string
	}{
		0: {
			in:  El{},
			exp: "root should be of type layout",
		},
		1: {
			in:  El{Type: TypeElLayout},
			exp: "layout has wrong number of children to render (expected only one)",
		},
		2: {
			in:  El{Type: TypeElLayout, Children: []El{El{}, El{}}},
			exp: "layout has wrong number of children to render (expected only one)",
		},
		3: {
			in:  El{Type: TypeElLayout, Children: []El{El{}}},
			exp: "el with an invalid type found : ''",
		},
		31: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: ElType("fake1")},
			}},
			exp: "el with an invalid type found : 'fake1'",
		},
		32: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElColumn, Children: []El{
					El{Type: ElType("fake2")},
				}},
			}},
			exp: "el with an invalid type found : 'fake2'",
		},
		4: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElImage, Children: []El{El{}}},
			}},
			exp: "invalid image found : should not have children",
		},
		5: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElColumn, Children: []El{
					El{Type: TypeElImage, Children: []El{El{}}},
				}},
			}},
			exp: "invalid image found : should not have children",
		},
		51: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElColumn, Children: []El{
					El{Type: TypeElText, Children: []El{El{}}},
				}},
			}},
			exp: "invalid child found : text should be placed in el",
		},
		6: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElEl, Children: []El{
					El{Type: TypeElText, Children: []El{El{}}},
				}},
			}},
			exp: "invalid text found : should not have children",
		},
		7: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElColumn, Children: []El{
					El{Type: TypeElEl, Children: []El{
						El{Type: TypeElText, Children: []El{El{}}},
					}},
				}},
			}},
			exp: "invalid text found : should not have children",
		},
		8: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElRow, Children: []El{
					El{Type: TypeElEl, Children: []El{
						El{Type: TypeElText, Children: []El{El{}}},
					}},
				}},
			}},
			exp: "invalid text found : should not have children",
		},
		9: {
			in: El{Type: TypeElLayout, Children: []El{
				El{Type: TypeElRow, Children: []El{
					El{Type: TypeElRow, Children: []El{
						El{Type: TypeElEl, Children: []El{
							El{Type: TypeElText, Children: []El{El{}}},
						}},
					}},
				}},
			}},
			exp: "invalid text found : should not have children",
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
