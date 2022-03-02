package ast

import (
	"fmt"
)

var (
	errInvalidRootType           = fmt.Errorf("root should be of type layout")
	errInvalidNumberRootChildren = fmt.Errorf("layout has wrong number of children to render (expected only one)")
	errInvalideTypeFound         = fmt.Errorf("el with an invalid type found")
	errInvalidImageType          = fmt.Errorf("invalid image found")
	errInvalidTextType           = fmt.Errorf("invalid text found")
	errShouldNotHaveChildren     = fmt.Errorf("should not have children")
)

func Validate(el El) error {
	if el.Type != TypeElLayout {
		return errInvalidRootType
	}
	if len(el.Children) != 1 {
		return errInvalidNumberRootChildren
	}
	if el.Children[0].Type == "" {
		return fmt.Errorf("%w : '%s'", errInvalideTypeFound, el.Children[0].Type)
	}
	switch el.Children[0].Type {
	case TypeElImage:
		return validateImage(el.Children[0])
	case TypeElColumn:
		return validate(el.Children[0])
	case TypeElRow:
		return validate(el.Children[0])
	case TypeElText:
		return validateText(el.Children[0])
	case TypeElEl:
		return validate(el.Children[0])
	}
	return nil
}

func validate(el El) error {
	for _, c := range el.Children {
		var err error
		switch c.Type {
		case TypeElImage:
			err = validateImage(c)
		case TypeElColumn:
			err = validate(c)
		case TypeElRow:
			err = validate(c)
		case TypeElText:
			err = validateText(c)
		case TypeElEl:
			return validate(c)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func validateImage(el El) error {
	if len(el.Children) != 0 {
		return fmt.Errorf("%v : %v", errInvalidImageType, errShouldNotHaveChildren)
	}
	return nil
}

func validateText(el El) error {
	if len(el.Children) != 0 {
		return fmt.Errorf("%v : %v", errInvalidTextType, errShouldNotHaveChildren)
	}
	return nil
}
