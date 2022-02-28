package colors

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func FromString(attribute string) (color.RGBA, error) {

	if attribute == "" {
		return color.RGBA{}, fmt.Errorf("empty value")
	}
	if strings.HasPrefix(attribute, "rgb(") && strings.HasSuffix(attribute, ")") {
		return parseFuncRgbColor(attribute)
	}

	if attribute[0] == '#' {
		return parseHexColor(attribute)
	}

	rgb, ok := colorDict[attribute]
	if !ok {
		return color.RGBA{}, fmt.Errorf("Unknown color %s", attribute)
	}
	return rgb, nil
}

func parseHexColor(s string) (color.RGBA, error) {
	var (
		c   color.RGBA
		err error
	)
	fullSize := 7
	shortSize := 4
	shortSizeScale := uint8(17)

	switch len(s) {
	case fullSize:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case shortSize:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		c.R *= shortSizeScale
		c.G *= shortSizeScale
		c.B *= shortSizeScale
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return c, err
}

func parseFuncRgbColor(attribute string) (color.RGBA, error) {
	colorAttr := attribute[4 : len(attribute)-1]
	colorPart := strings.Split(colorAttr, ",")
	if len(colorPart) != 3 {
		return color.RGBA{}, fmt.Errorf("invalid number of arg provider expected 3 got %d in %s", len(colorPart), colorAttr)
	}
	r, err := strconv.ParseUint(colorPart[0], 10, 8)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid parsing of argument 1 (%s) : %w", colorPart[0], err)
	}
	g, err := strconv.ParseUint(colorPart[1], 10, 8)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid parsing of argument 2 (%s) : %w", colorPart[0], err)
	}
	b, err := strconv.ParseUint(colorPart[2], 10, 8)
	if err != nil {
		return color.RGBA{}, fmt.Errorf("invalid parsing of argument 3 (%s) : %w", colorPart[0], err)
	}

	return color.RGBA{uint8(r), uint8(g), uint8(b), 0}, nil
}
