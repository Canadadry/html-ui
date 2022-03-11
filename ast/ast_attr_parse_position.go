package ast

import (
	"fmt"
)

const (
	positonLeft  = "left"
	positonRight = "right"
	positonAbove = "above"
	positonBelow = "below"
)

var (
	ErrInvalidPosition = fmt.Errorf("invalid position")
)

func ParsePositionAttr(v string) error {
	var err error
	switch v {
	case positonLeft:
	case positonRight:
	case positonAbove:
	case positonBelow:
	default:
		err = fmt.Errorf("%w : %s", ErrInvalidPosition, v)
	}
	return err
}
