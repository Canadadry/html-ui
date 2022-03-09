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
	errInvalidChildText          = fmt.Errorf("invalid child found : text should be placed in el")
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
	err := validateAttribute(el)
	if err != nil {
		return err
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
	case TypeElForm:
		return validate(el.Children[0])
	}
	return fmt.Errorf("%w : '%s'", errInvalideTypeFound, el.Children[0].Type)
}

func validate(el El) error {
	err := validateAttribute(el)
	if err != nil {
		return err
	}
	for _, c := range el.Children {
		var err error
		switch c.Type {
		case TypeElImage:
			err = validateImage(c)
		case TypeElColumn:
			err = validate(c)
		case TypeElRow:
			err = validate(c)
		case TypeElButton:
			err = validate(c)
		case TypeElText:
			if el.Type != TypeElEl {
				return errInvalidChildText
			}
			err = validateText(c)
		case TypeElEl:
			err = validate(c)
		case TypeElForm:
			err = validate(c)
		default:
			err = fmt.Errorf("%w : '%s'", errInvalideTypeFound, c.Type)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func validateImage(el El) error {
	err := validateAttribute(el)
	if err != nil {
		return err
	}
	if len(el.Children) != 0 {
		return fmt.Errorf("%v : %v", errInvalidImageType, errShouldNotHaveChildren)
	}
	return nil
}

func validateText(el El) error {
	err := validateAttribute(el)
	if err != nil {
		return err
	}
	if len(el.Children) != 0 {
		return fmt.Errorf("%v : %v", errInvalidTextType, errShouldNotHaveChildren)
	}
	return nil
}
