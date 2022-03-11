package ast

import (
	"fmt"
)

const (
	PositonLeft  = "left"
	PositonRight = "right"
	PositonAbove = "above"
	PositonBelow = "below"
)

var (
	ErrInvalidPosition = fmt.Errorf("invalid position")
)

func ParsePositionAttr(v string) error {
	var err error
	switch v {
	case PositonLeft:
	case PositonRight:
	case PositonAbove:
	case PositonBelow:
	default:
		err = fmt.Errorf("%w : %s", ErrInvalidPosition, v)
	}
	return err
}
