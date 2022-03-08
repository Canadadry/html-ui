package ast

import (
	"app/pkg/colors"
	"image/color"
)

func ParseColorAttr(v string) (color.RGBA, error) {
	return colors.FromString(v)
}
