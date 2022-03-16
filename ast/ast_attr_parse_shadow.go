package ast

import (
	"fmt"
	"image/color"
	"strings"
)

type Shadow struct {
	OffsetX int
	OffsetY int
	Blur    int
	Size    int
	Color   color.RGBA
	Inner   bool
}

func ParseShadowAttr(v string) (Shadow, error) {
	s := Shadow{}
	params := strings.Split(v, ";")
	for _, p := range params {
		var err error
		switch true {
		case strings.HasPrefix(p, "inner"):
			s.Inner = true
		case strings.HasPrefix(p, "ox:"):
			s.OffsetX, err = ParseNumberAttr(p[3:])
		case strings.HasPrefix(p, "oy:"):
			s.OffsetY, err = ParseNumberAttr(p[3:])
		case strings.HasPrefix(p, "b:"):
			s.Blur, err = ParseNumberAttr(p[2:])
		case strings.HasPrefix(p, "s:"):
			s.Size, err = ParseNumberAttr(p[2:])
		case strings.HasPrefix(p, "color:"):
			s.Color, err = ParseColorAttr(p[6:])
		default:
			err = fmt.Errorf("cannot parse shadow attr '%s'", p)
		}
		if err != nil {
			return s, err
		}
	}
	return s, nil
}
