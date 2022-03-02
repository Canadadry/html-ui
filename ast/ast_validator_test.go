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
						Type: TypeElColumn,
						Children: []El{
							El{
								Type: TypeElImage,
							},
							El{
								Type: TypeElText,
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
