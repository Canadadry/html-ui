package generator

import (
	"app/ast"
	"app/pkg/html"
	"fmt"
	"io"
)

func Generate(el ast.El, w io.Writer) error {
	gen := generator{
		css: map[string]struct{}{},
	}
	body := gen.generate([]ast.El{el})
	head := generateHead(gen.css)
	d := html.Document{
		Head: head,
		Body: body,
	}
	return d.Render(w)
}

type mode string

const (
	modeNormal mode = ""
	modeColumn      = "column"
)

type generator struct {
	css  UniqueClasses
	mode mode
}

func (g *generator) generate(el []ast.El) []html.Tag {
	out := make([]html.Tag, 0, len(el))
	for _, item := range el {
		switch item.Type {
		case ast.TypeElLayout:
			out = append(out, g.generateLayout(item))
		case ast.TypeElColumn:
			out = append(out, g.generateColumn(item))
		case ast.TypeElRow:
			out = append(out, g.generateRow(item))
		case ast.TypeElEl:
			out = append(out, g.generateEl(item))
		case ast.TypeElImage:
			out = append(out, g.generateImage(item))
		case ast.TypeElText:
			out = append(out, g.generateText(item.Content))
		}
	}
	return out
}

func (g *generator) generateLayout(el ast.El) html.Tag {
	return html.Div(
		html.Attributes{html.AttributeClass: "bg-255-255-255-0 fc-0-0-0-255 font-size-20 font-open-sanshelveticaverdanasans-serif s e ui s e"},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateColumn(el ast.El) html.Tag {
	classes := g.parseAttribute(el.Attr, UniqueClassesFrom("s c ct cl"))
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateRow(el ast.El) html.Tag {
	classes := g.parseAttribute(el.Attr, UniqueClassesFrom("s r cl ccy"))
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
}

func (g *generator) parseAttribute(attrs []ast.Attribute, base UniqueClasses) string {
	hasWidthAttr := false
	hasHeightAttr := false
	for _, attr := range attrs {
		class := ""
		switch attr.Type {
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
			continue
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
	return base.String()
}

func (g *generator) generateEl(el ast.El) html.Tag {
	if len(el.Children) > 0 && el.Children[0].Type == ast.TypeElText {
		return g.generateElText(el)
	}
	classes := g.parseAttribute(el.Attr, UniqueClassesFrom("s e"))
	g.mode = modeNormal
	tag := html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
	if len(el.Children) == 0 {
		return html.Inline(tag)
	}
	return tag
}

func (g *generator) generateElText(el ast.El) html.Tag {
	if len(el.Attr) == 0 {
		return g.generateText(el.Children[0].Content)
	}
	classes := g.parseAttribute(el.Attr, UniqueClassesFrom("s e"))
	g.mode = modeNormal
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generateText(el.Children[0].Content),
	)
}

func (g *generator) generateImage(el ast.El) html.Tag {
	src := ""
	alt := ""
	for _, attr := range el.Attr {
		if attr.Type == ast.TypeAttrSrc {
			src = attr.Value
		}
		if attr.Type == ast.TypeAttrAlt {
			alt = attr.Value
		}
	}
	divClasses := UniqueClassesFrom(g.parseAttribute(el.Attr, UniqueClassesFrom("s e ic")))
	divClasses.Remove("hc")
	divClasses.Remove("wc")
	if divClasses.Has("he") {
		divClasses.Remove("ic")
		divClasses.Add("i")
	}
	imgClasses := UniqueClassesFrom(g.parseAttribute(el.Attr, UniqueClassesFrom("s e")))
	imgClasses.Remove("hc")
	imgClasses.Remove("wc")
	return html.Div(
		html.Attributes{html.AttributeClass: divClasses.String()},
		html.Img(
			html.Attributes{
				html.AttributeClass: imgClasses.String(),
				html.AttributeSrc:   src,
				html.AttributeAlt:   alt,
			},
		),
	)
}

func (g *generator) generateText(txt string) html.Tag {
	class := "s t wf hf"
	if g.mode == modeColumn {
		class = "s t wc hc"
	}
	return html.Inline(html.Div(html.Attributes{html.AttributeClass: class},
		html.Text(txt),
	))
}
