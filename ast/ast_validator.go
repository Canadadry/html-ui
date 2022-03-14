package ast

import (
	"fmt"
)

type childType map[ElType]struct{}

func (ct childType) merge(t ...ElType) childType {
	newChildTypes := childType{}
	for types := range ct {
		newChildTypes[types] = struct{}{}
	}
	for _, types := range t {
		newChildTypes[types] = struct{}{}
	}
	return newChildTypes
}

type AllowedChild map[ElType]childType

type MaxChildrenLenByType map[ElType]int

var (
	errInvalidRootType               = fmt.Errorf("root should be of type layout")
	errInvalideTypeFound             = fmt.Errorf("element with an invalid type found")
	errInvalidChildType              = fmt.Errorf("invalid child found")
	errInvalidChildrenLen            = fmt.Errorf("invalid children len")
	errInvalidChildrenDuplicatedType = fmt.Errorf("can have only one element of each type")

	basicEl = childType{
		TypeElEl:     {},
		TypeElRow:    {},
		TypeElColumn: {},
		TypeElImage:  {},
		TypeElButton: {},
		TypeElForm:   {},
		TypeElInput:  {},
	}
	allowedChild = AllowedChild{
		TypeElLayout: basicEl.merge(TypeElText, TypeElDefinition),
		TypeElEl:     basicEl.merge(TypeElText),
		TypeElRow:    basicEl,
		TypeElColumn: basicEl,
		TypeElImage:  {},
		TypeElText:   {},
		TypeElButton: basicEl,
		TypeElForm:   basicEl,
		TypeElInput: {
			TypeElLabel:       {},
			TypeElPlaceholder: {},
		},
		TypeElLabel: {
			TypeElText:   {},
			TypeElEl:     {},
			TypeElRow:    {},
			TypeElColumn: {},
		},
		TypeElPlaceholder: {
			TypeElText: {},
		},
		TypeElDefinition: {
			TypeElFont: {},
		},
		TypeElFont: {},
	}
	maxChildrenLenByType = MaxChildrenLenByType{
		TypeElLayout:      2,
		TypeElText:        0,
		TypeElImage:       0,
		TypeElFont:        0,
		TypeElInput:       2,
		TypeElLabel:       1,
		TypeElPlaceholder: 1,
		TypeElEl:          -1,
		TypeElRow:         -1,
		TypeElColumn:      -1,
		TypeElButton:      -1,
		TypeElForm:        -1,
		TypeElDefinition:  -1,
	}
)

func Validate(el El) error {
	return validateRoot(el, allowedChild, maxChildrenLenByType)
}

func validateRoot(el El, a AllowedChild, m MaxChildrenLenByType) error {
	if el.Type != TypeElLayout {
		return errInvalidRootType
	}
	if err := validateAttribute(el); err != nil {
		return err
	}
	if err := validateChildrenMaxLen(el, m[TypeElLayout]); err != nil {
		return err
	}
	if err := validateChildrenPossibleType(el, a[TypeElLayout]); err != nil {
		return err
	}
	if err := restrictChildToOneOfEachTypes(el); err != nil {
		return err
	}
	return validateNode(el, a, m)
}

func validateNode(el El, a AllowedChild, m MaxChildrenLenByType) error {
	err := validateAttribute(el)
	if err != nil {
		return err
	}
	for _, c := range el.Children {
		max, ok := m[c.Type]
		if !ok {
			return fmt.Errorf("%w : '%s' when trying to validate children len", errInvalideTypeFound, c.Type)
		}
		if err := validateChildrenMaxLen(c, max); err != nil {
			return err
		}
		types, ok := a[c.Type]
		if !ok {
			return fmt.Errorf("%w : '%s' when trying to validate child type", errInvalideTypeFound, c.Type)
		}
		if err := validateChildrenPossibleType(c, types); err != nil {
			return err
		}
		if err := validateNode(c, a, m); err != nil {
			return err
		}
	}
	return nil
}

func validateChildrenMaxLen(el El, max int) error {
	if max < 0 {
		return nil
	}
	if len(el.Children) > max {
		return fmt.Errorf("%w : max %d got %d for %s", errInvalidChildrenLen, max, len(el.Children), el.Type)
	}
	return nil
}

func validateChildrenPossibleType(el El, types map[ElType]struct{}) error {
	for _, c := range el.Children {
		if _, ok := types[c.Type]; !ok {
			return fmt.Errorf("%w : %s cannot have child of type %s", errInvalidChildType, el.Type, c.Type)
		}
	}
	return nil
}

func restrictChildToOneOfEachTypes(el El) error {
	count := map[ElType]int{}
	for _, c := range el.Children {
		if _, ok := count[c.Type]; !ok {
			count[c.Type] = 1
		} else {
			return fmt.Errorf("%w : found at least two of %s", errInvalidChildrenDuplicatedType, c.Type)
		}
	}
	return nil
}
