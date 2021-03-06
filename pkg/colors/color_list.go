package colors

import (
	"image/color"
)

var colorDict map[string]color.RGBA = map[string]color.RGBA{
	"aliceblue":            color.RGBA{240, 248, 255, 0},
	"antiquewhite":         color.RGBA{250, 235, 215, 0},
	"aqua":                 color.RGBA{0, 255, 255, 0},
	"aquamarine":           color.RGBA{127, 255, 212, 0},
	"azure":                color.RGBA{240, 255, 255, 0},
	"beige":                color.RGBA{245, 245, 220, 0},
	"bisque":               color.RGBA{255, 228, 196, 0},
	"black":                color.RGBA{0, 0, 0, 0},
	"blanchedalmond":       color.RGBA{255, 235, 205, 0},
	"blue":                 color.RGBA{0, 0, 255, 0},
	"blueviolet":           color.RGBA{138, 43, 226, 0},
	"brown":                color.RGBA{165, 42, 42, 0},
	"burlywood":            color.RGBA{222, 184, 135, 0},
	"cadetblue":            color.RGBA{95, 158, 160, 0},
	"chartreuse":           color.RGBA{127, 255, 0, 0},
	"chocolate":            color.RGBA{210, 105, 30, 0},
	"coral":                color.RGBA{255, 127, 80, 0},
	"cornflowerblue":       color.RGBA{100, 149, 237, 0},
	"cornsilk":             color.RGBA{255, 248, 220, 0},
	"crimson":              color.RGBA{220, 20, 60, 0},
	"cyan":                 color.RGBA{0, 255, 255, 0},
	"darkblue":             color.RGBA{0, 0, 139, 0},
	"darkcyan":             color.RGBA{0, 139, 139, 0},
	"darkgoldenrod":        color.RGBA{184, 134, 11, 0},
	"darkgray":             color.RGBA{169, 169, 169, 0},
	"darkgreen":            color.RGBA{0, 100, 0, 0},
	"darkgrey":             color.RGBA{169, 169, 169, 0},
	"darkkhaki":            color.RGBA{189, 183, 107, 0},
	"darkmagenta":          color.RGBA{139, 0, 139, 0},
	"darkolivegreen":       color.RGBA{85, 107, 47, 0},
	"darkorange":           color.RGBA{255, 140, 0, 0},
	"darkorchid":           color.RGBA{153, 50, 204, 0},
	"darkred":              color.RGBA{139, 0, 0, 0},
	"darksalmon":           color.RGBA{233, 150, 122, 0},
	"darkseagreen":         color.RGBA{143, 188, 143, 0},
	"darkslateblue":        color.RGBA{72, 61, 139, 0},
	"darkslategray":        color.RGBA{47, 79, 79, 0},
	"darkslategrey":        color.RGBA{47, 79, 79, 0},
	"darkturquoise":        color.RGBA{0, 206, 209, 0},
	"darkviolet":           color.RGBA{148, 0, 211, 0},
	"deeppink":             color.RGBA{255, 20, 147, 0},
	"deepskyblue":          color.RGBA{0, 191, 255, 0},
	"dimgray":              color.RGBA{105, 105, 105, 0},
	"dimgrey":              color.RGBA{105, 105, 105, 0},
	"dodgerblue":           color.RGBA{30, 144, 255, 0},
	"firebrick":            color.RGBA{178, 34, 34, 0},
	"floralwhite":          color.RGBA{255, 250, 240, 0},
	"forestgreen":          color.RGBA{34, 139, 34, 0},
	"fuchsia":              color.RGBA{255, 0, 255, 0},
	"gainsboro":            color.RGBA{220, 220, 220, 0},
	"ghostwhite":           color.RGBA{248, 248, 255, 0},
	"gold":                 color.RGBA{255, 215, 0, 0},
	"goldenrod":            color.RGBA{218, 165, 32, 0},
	"gray":                 color.RGBA{128, 128, 128, 0},
	"grey":                 color.RGBA{128, 128, 128, 0},
	"green":                color.RGBA{0, 128, 0, 0},
	"greenyellow":          color.RGBA{173, 255, 47, 0},
	"honeydew":             color.RGBA{240, 255, 240, 0},
	"hotpink":              color.RGBA{255, 105, 180, 0},
	"indianred":            color.RGBA{205, 92, 92, 0},
	"indigo":               color.RGBA{75, 0, 130, 0},
	"ivory":                color.RGBA{255, 255, 240, 0},
	"khaki":                color.RGBA{240, 230, 140, 0},
	"lavender":             color.RGBA{230, 230, 250, 0},
	"lavenderblush":        color.RGBA{255, 240, 245, 0},
	"lawngreen":            color.RGBA{124, 252, 0, 0},
	"lemonchiffon":         color.RGBA{255, 250, 205, 0},
	"lightblue":            color.RGBA{173, 216, 230, 0},
	"lightcoral":           color.RGBA{240, 128, 128, 0},
	"lightcyan":            color.RGBA{224, 255, 255, 0},
	"lightgoldenrodyellow": color.RGBA{250, 250, 210, 0},
	"lightgray":            color.RGBA{211, 211, 211, 0},
	"lightgreen":           color.RGBA{144, 238, 144, 0},
	"lightgrey":            color.RGBA{211, 211, 211, 0},
	"lightpink":            color.RGBA{255, 182, 193, 0},
	"lightsalmon":          color.RGBA{255, 160, 122, 0},
	"lightseagreen":        color.RGBA{32, 178, 170, 0},
	"lightskyblue":         color.RGBA{135, 206, 250, 0},
	"lightslategray":       color.RGBA{119, 136, 153, 0},
	"lightslategrey":       color.RGBA{119, 136, 153, 0},
	"lightsteelblue":       color.RGBA{176, 196, 222, 0},
	"lightyellow":          color.RGBA{255, 255, 224, 0},
	"lime":                 color.RGBA{0, 255, 0, 0},
	"limegreen":            color.RGBA{50, 205, 50, 0},
	"linen":                color.RGBA{250, 240, 230, 0},
	"magenta":              color.RGBA{255, 0, 255, 0},
	"maroon":               color.RGBA{128, 0, 0, 0},
	"mediumaquamarine":     color.RGBA{102, 205, 170, 0},
	"mediumblue":           color.RGBA{0, 0, 205, 0},
	"mediumorchid":         color.RGBA{186, 85, 211, 0},
	"mediumpurple":         color.RGBA{147, 112, 219, 0},
	"mediumseagreen":       color.RGBA{60, 179, 113, 0},
	"mediumslateblue":      color.RGBA{123, 104, 238, 0},
	"mediumspringgreen":    color.RGBA{0, 250, 154, 0},
	"mediumturquoise":      color.RGBA{72, 209, 204, 0},
	"mediumvioletred":      color.RGBA{199, 21, 133, 0},
	"midnightblue":         color.RGBA{25, 25, 112, 0},
	"mintcream":            color.RGBA{245, 255, 250, 0},
	"mistyrose":            color.RGBA{255, 228, 225, 0},
	"moccasin":             color.RGBA{255, 228, 181, 0},
	"navajowhite":          color.RGBA{255, 222, 173, 0},
	"navy":                 color.RGBA{0, 0, 128, 0},
	"oldlace":              color.RGBA{253, 245, 230, 0},
	"olive":                color.RGBA{128, 128, 0, 0},
	"olivedrab":            color.RGBA{107, 142, 35, 0},
	"orange":               color.RGBA{255, 165, 0, 0},
	"orangered":            color.RGBA{255, 69, 0, 0},
	"orchid":               color.RGBA{218, 112, 214, 0},
	"palegoldenrod":        color.RGBA{238, 232, 170, 0},
	"palegreen":            color.RGBA{152, 251, 152, 0},
	"paleturquoise":        color.RGBA{175, 238, 238, 0},
	"palevioletred":        color.RGBA{219, 112, 147, 0},
	"papayawhip":           color.RGBA{255, 239, 213, 0},
	"peachpuff":            color.RGBA{255, 218, 185, 0},
	"peru":                 color.RGBA{205, 133, 63, 0},
	"pink":                 color.RGBA{255, 192, 203, 0},
	"plum":                 color.RGBA{221, 160, 221, 0},
	"powderblue":           color.RGBA{176, 224, 230, 0},
	"purple":               color.RGBA{128, 0, 128, 0},
	"red":                  color.RGBA{255, 0, 0, 0},
	"rosybrown":            color.RGBA{188, 143, 143, 0},
	"royalblue":            color.RGBA{65, 105, 225, 0},
	"saddlebrown":          color.RGBA{139, 69, 19, 0},
	"salmon":               color.RGBA{250, 128, 114, 0},
	"sandybrown":           color.RGBA{244, 164, 96, 0},
	"seagreen":             color.RGBA{46, 139, 87, 0},
	"seashell":             color.RGBA{255, 245, 238, 0},
	"sienna":               color.RGBA{160, 82, 45, 0},
	"silver":               color.RGBA{192, 192, 192, 0},
	"skyblue":              color.RGBA{135, 206, 235, 0},
	"slateblue":            color.RGBA{106, 90, 205, 0},
	"slategray":            color.RGBA{112, 128, 144, 0},
	"slategrey":            color.RGBA{112, 128, 144, 0},
	"snow":                 color.RGBA{255, 250, 250, 0},
	"springgreen":          color.RGBA{0, 255, 127, 0},
	"steelblue":            color.RGBA{70, 130, 180, 0},
	"tan":                  color.RGBA{210, 180, 140, 0},
	"teal":                 color.RGBA{0, 128, 128, 0},
	"thistle":              color.RGBA{216, 191, 216, 0},
	"tomato":               color.RGBA{255, 99, 71, 0},
	"turquoise":            color.RGBA{64, 224, 208, 0},
	"violet":               color.RGBA{238, 130, 238, 0},
	"wheat":                color.RGBA{245, 222, 179, 0},
	"white":                color.RGBA{255, 255, 255, 0},
	"whitesmoke":           color.RGBA{245, 245, 245, 0},
	"yellow":               color.RGBA{255, 255, 0, 0},
	"yellowgreen":          color.RGBA{154, 205, 50, 0},
}
