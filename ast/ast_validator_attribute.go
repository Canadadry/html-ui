package ast

import (
	"fmt"
	"sort"
)

func validateAttribute(el El) error {
	switch el.Type {
	case TypeElLayout:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	case TypeElImage:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	case TypeElColumn:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	case TypeElRow:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	case TypeElText:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	case TypeElEl:
		return validateAttributeField(el.Attr, map[AttrType]struct{}{})
	}
	return nil
}

func validateAttributeField(attrs []Attribute, legal map[AttrType]struct{}) error {
	for _, attr := range attrs {
		_, ok := legal[attr.Type]
		if !ok {
			return fmt.Errorf("layout cannot have attribute '%s' possibilities are %v", attr.Type, mapAttrTypeToString(legal))
		}
	}
	return nil
}

func mapAttrTypeToString(m map[AttrType]struct{}) []string {
	out := make([]string, 0, len(m))
	for t := range m {
		out = append(out, string(t))
	}
	sort.Strings(out)
	return out
}
