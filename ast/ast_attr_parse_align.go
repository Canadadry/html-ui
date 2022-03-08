package ast

import (
	"fmt"
	"strings"
)

type AlignXType string
type AlignYType string

const (
	AlignLeft    AlignXType = "left"
	AlignRight              = "right"
	AlignCenterX            = "centerX"

	AlignTop     AlignYType = "top"
	AlignBottom             = "bottom"
	AlignCenterY            = "centerY"
)

func ParseAlignAttr(v string) (AlignXType, AlignYType, error) {
	aligns := strings.Split(v, ",")
	x := AlignLeft
	y := AlignTop
	for _, a := range aligns {
		switch a {
		case "left":
			x = AlignLeft
		case "right":
			x = AlignRight
		case "centerX":
			x = AlignCenterX
		case "top":
			y = AlignTop
		case "bottom":
			y = AlignBottom
		case "centerY":
			y = AlignCenterY
		default:
			return x, y, fmt.Errorf("invalid align value")
		}
	}
	return x, y, nil
}
