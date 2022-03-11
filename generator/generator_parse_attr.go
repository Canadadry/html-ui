package generator

import (
	"app/ast"
	"fmt"
)

func (g *generator) parseAttribute(attrs []ast.Attribute, base UniqueClasses) (string, error) {
	hasWidthAttr := false
	hasHeightAttr := false
	for _, attr := range attrs {
		class := ""
		switch attr.Type {
		case ast.TypeAttrSrc:
			continue
		case ast.TypeAttrAlt:
			continue
		case ast.TypeAttrName:
			continue
		case ast.TypeAttrValue:
			continue
		case ast.TypeAttrType:
			continue
		case ast.TypeAttrSpacing:
			class = fmt.Sprintf("spacing-%s-%s", attr.Value, attr.Value)
		case ast.TypeAttrFontSize:
			class = fmt.Sprintf("font-size-%s", attr.Value)
		case ast.TypeAttrBorderWidth:
			class = fmt.Sprintf("b-%s", attr.Value)
		case ast.TypeAttrBorderRounded:
			class = fmt.Sprintf("br-%s", attr.Value)
		case ast.TypeAttrPadding:
			class = fmt.Sprintf("p-%s", attr.Value)
		case ast.TypeAttrBgColor:
			class = fmt.Sprintf("bg-%d-%d-%d-255", attr.Color.R, attr.Color.G, attr.Color.B)
		case ast.TypeAttrFocusBgColor:
			class = fmt.Sprintf("bg-%d-%d-%d-255-fs", attr.Color.R, attr.Color.G, attr.Color.B)
		case ast.TypeAttrFontColor:
			class = fmt.Sprintf("fc-%d-%d-%d-255", attr.Color.R, attr.Color.G, attr.Color.B)
		case ast.TypeAttrBorderColor:
			class = fmt.Sprintf("bc-%d-%d-%d-255", attr.Color.R, attr.Color.G, attr.Color.B)
		case ast.TypeAttrWidth:
			hasWidthAttr = true
			switch attr.Size.Type() {
			case ast.SizePxType:
				class = fmt.Sprintf("width-px-%d", attr.Size.Get())
				base.Add("we")
			case ast.SizePortionType:
				class = fmt.Sprintf("width-fill-%d", attr.Size.Get())
				base.Add("wfp")
			case ast.SizeFillType:
				class = "wf"
			case ast.SizeNoneType:
			}
		case ast.TypeAttrHeight:
			hasHeightAttr = true
			switch attr.Size.Type() {
			case ast.SizePxType:
				class = fmt.Sprintf("height-px-%d", attr.Size.Get())
				base.Add("he")
			case ast.SizePortionType:
				class = fmt.Sprintf("height-fill-%d", attr.Size.Get())
				base.Add("hfp")
			case ast.SizeFillType:
				class = "hf"
			case ast.SizeNoneType:
			}
		case ast.TypeAttrAlign:
			switch attr.AlignX {
			case "left":
				base.Add("av")
				base.Add("al")
			case "right":
				base.Add("av")
				base.Add("ar")
			case "centerX":
				base.Add("av")
				base.Add("cx")
			}
			switch attr.AlignY {
			case "top":
				base.Add("ah")
				base.Add("at")
			case "bottom":
				base.Add("ah")
				base.Add("ab")
			case "centerY":
				base.Add("ah")
				base.Add("cy")
			}
		default:
			return "", fmt.Errorf("cannot generate class for attribute '%s'", attr.Type)
		}
		if class != "" {
			base.Add(class)
			g.css.Add(class)
		}
	}
	if !hasWidthAttr {
		base.Add("wc")
	}
	if !hasHeightAttr {
		base.Add("hc")
	}
	return base.String(), nil
}
