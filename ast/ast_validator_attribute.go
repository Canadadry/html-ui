package ast

import (
	"fmt"
	"sort"
)

func validateAttribute(el El) error {
	validationCtx := map[ElType]map[AttrType]struct{}{
		TypeElLayout:     map[AttrType]struct{}{},
		TypeElDefinition: map[AttrType]struct{}{},
		TypeElFont: map[AttrType]struct{}{
			TypeAttrSrc:  {},
			TypeAttrName: {},
		},
		TypeElImage: map[AttrType]struct{}{
			TypeAttrWidth:  {},
			TypeAttrHeight: {},
			TypeAttrSrc:    {},
			TypeAttrAlt:    {},
		},
		TypeElColumn: map[AttrType]struct{}{
			TypeAttrWidth:         {},
			TypeAttrHeight:        {},
			TypeAttrAlign:         {},
			TypeAttrSpacing:       {},
			TypeAttrPadding:       {},
			TypeAttrBgColor:       {},
			TypeAttrFontColor:     {},
			TypeAttrFontSize:      {},
			TypeAttrBorderRounded: {},
			TypeAttrBorderColor:   {},
			TypeAttrBorderWidth:   {},
			TypeAttrFontFamily:    {},
			TypeAttrFontWeigth:    {},
			TypeAttrShadow:        {},
		},
		TypeElRow: map[AttrType]struct{}{
			TypeAttrWidth:         {},
			TypeAttrHeight:        {},
			TypeAttrAlign:         {},
			TypeAttrSpacing:       {},
			TypeAttrPadding:       {},
			TypeAttrBgColor:       {},
			TypeAttrFontColor:     {},
			TypeAttrFontSize:      {},
			TypeAttrBorderRounded: {},
			TypeAttrBorderColor:   {},
			TypeAttrBorderWidth:   {},
			TypeAttrFontFamily:    {},
			TypeAttrFontWeigth:    {},
			TypeAttrShadow:        {},
		},
		TypeElText: map[AttrType]struct{}{},
		TypeElButton: map[AttrType]struct{}{
			TypeAttrWidth:         {},
			TypeAttrHeight:        {},
			TypeAttrAlign:         {},
			TypeAttrSpacing:       {},
			TypeAttrPadding:       {},
			TypeAttrBgColor:       {},
			TypeAttrFocusBgColor:  {},
			TypeAttrFontColor:     {},
			TypeAttrFontSize:      {},
			TypeAttrBorderRounded: {},
			TypeAttrBorderColor:   {},
			TypeAttrBorderWidth:   {},
			TypeAttrName:          {},
			TypeAttrValue:         {},
			TypeAttrFontFamily:    {},
			TypeAttrFontWeigth:    {},
			TypeAttrShadow:        {},
		},
		TypeElEl: map[AttrType]struct{}{
			TypeAttrWidth:         {},
			TypeAttrHeight:        {},
			TypeAttrAlign:         {},
			TypeAttrSpacing:       {},
			TypeAttrPadding:       {},
			TypeAttrBgColor:       {},
			TypeAttrFontColor:     {},
			TypeAttrFontSize:      {},
			TypeAttrBorderRounded: {},
			TypeAttrBorderColor:   {},
			TypeAttrBorderWidth:   {},
			TypeAttrFontFamily:    {},
			TypeAttrFontWeigth:    {},
			TypeAttrShadow:        {},
		},
		TypeElForm: map[AttrType]struct{}{
			TypeAttrMethod: {},
			TypeAttrAction: {},
			TypeAttrName:   {},
		},
		TypeElInput: map[AttrType]struct{}{
			TypeAttrName:  {},
			TypeAttrValue: {},
			TypeAttrType:  {},
		},
		TypeElLabel: map[AttrType]struct{}{
			TypeAttrPosition: {},
		},
		TypeElPlaceholder: map[AttrType]struct{}{},
	}

	ctx, ok := validationCtx[el.Type]
	if !ok {
		return fmt.Errorf("cannot validate attribute type of unknown type '%s'", el.Type)
	}
	return validateAttributeField(el.Type, el.Attr, ctx)

}

func validateAttributeField(elType ElType, attrs []Attribute, legal map[AttrType]struct{}) error {
	for _, attr := range attrs {
		_, ok := legal[attr.Type]
		if !ok {
			return fmt.Errorf("%s cannot have attribute '%s' possibilities are %v", elType, attr.Type, mapAttrTypeToString(legal))
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
