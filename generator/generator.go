package generator

import (
	"app/ast"
	"app/pkg/colors"
	"app/pkg/html"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func Generate(el ast.El, w io.Writer) error {
	gen := generator{
		css: map[string]struct{}{},
	}
	body := gen.generate([]ast.El{el})
	head := gen.generateHead()
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
	css  map[string]struct{}
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
	base := map[string]struct{}{
		"s":  {},
		"c":  {},
		"ct": {},
		"cl": {},
	}
	classes := g.parseAttribute(el.Attr, base)
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateRow(el ast.El) html.Tag {
	base := map[string]struct{}{
		"s":   {},
		"r":   {},
		"cl":  {},
		"ccy": {},
	}
	classes := g.parseAttribute(el.Attr, base)
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
}

func (g *generator) parseAttribute(attrs []ast.Attribute, base map[string]struct{}) string {
	hasWidthAttr := false
	hasHeightAttr := false
	for _, attr := range attrs {
		class := ""
		switch attr.Type {
		case ast.TypeAttrSpacing:
			class = fmt.Sprintf("spacing-%s-%s", attr.Value, attr.Value)
		case ast.TypeAttrFontSize:
			class = fmt.Sprintf("font-size-%s", attr.Value)
		case ast.TypeAttrPadding:
			class = fmt.Sprintf("p-%s", attr.Value)
		case ast.TypeAttrBgColor:
			c, err := colors.FromString(attr.Value)
			if err != nil {
				continue
			}
			class = fmt.Sprintf("bg-%d-%d-%d-255", c.R, c.G, c.B)
		case ast.TypeAttrFontColor:
			c, err := colors.FromString(attr.Value)
			if err != nil {
				continue
			}
			class = fmt.Sprintf("fc-%d-%d-%d-255", c.R, c.G, c.B)
		case ast.TypeAttrWidth:
			var err error
			var sup string
			class, sup, err = parseWidthAttr(attr.Value)
			if err != nil {
				continue
			}
			hasWidthAttr = true
			if class == "width-fill" {
				class = "wf"
			}
			if sup != "" {
				base[sup] = struct{}{}
			}
		case ast.TypeAttrHeight:
			var err error
			var sup string
			class, sup, err = parseHeightAttr(attr.Value)
			if err != nil {
				continue
			}
			hasHeightAttr = true
			if class == "height-fill" {
				class = "hf"
			}
			if sup != "" {
				base[sup] = struct{}{}
			}
		default:
			continue
		}
		base[class] = struct{}{}
		g.css[class] = struct{}{}
	}
	if !hasWidthAttr {
		base["wc"] = struct{}{}
	}
	if !hasHeightAttr {
		base["hc"] = struct{}{}
	}
	classes := make([]string, 0, len(base))
	for c := range base {
		classes = append(classes, c)
	}
	sort.Strings(classes)
	return strings.Join(classes, " ")
}

func parseWidthAttr(width string) (string, string, error) {
	if strings.HasPrefix(width, "px:") {
		num, err := strconv.ParseInt(width[3:], 10, 64)
		return fmt.Sprintf("width-px-%d", num), "we", err
	}
	if strings.HasPrefix(width, "portion:") {
		num, err := strconv.ParseInt(width[8:], 10, 64)
		return fmt.Sprintf("width-fill-%d", num), "wfp", err
	}
	return "width-fill", "", nil
}

func parseHeightAttr(width string) (string, string, error) {
	if strings.HasPrefix(width, "px:") {
		num, err := strconv.ParseInt(width[3:], 10, 64)
		return fmt.Sprintf("height-px-%d", num), "he", err
	}
	if strings.HasPrefix(width, "portion:") {
		num, err := strconv.ParseInt(width[8:], 10, 64)
		return fmt.Sprintf("height-fill-%d", num), "hfp", err
	}
	return "height-fill", "", nil
}

func (g *generator) generateEl(el ast.El) html.Tag {
	if len(el.Children) == 0 {
		return html.Tag{}
	}
	if el.Children[0].Type == ast.TypeElText {
		return g.generateElText(el)
	}
	base := map[string]struct{}{
		"s": {},
		"e": {},
	}
	classes := g.parseAttribute(el.Attr, base)
	g.mode = modeNormal
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateElText(el ast.El) html.Tag {
	if len(el.Attr) == 0 {
		return g.generateText(el.Children[0].Content)
	}
	base := map[string]struct{}{
		"s": {},
		"e": {},
	}
	classes := g.parseAttribute(el.Attr, base)
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
	divClasses := g.parseAttribute(el.Attr, map[string]struct{}{"s": {}, "e": {}, "ic": {}})
	divClasses = strings.ReplaceAll(divClasses, "hc ", "")
	divClasses = strings.ReplaceAll(divClasses, "wc ", "")
	divClasses = strings.ReplaceAll(divClasses, " wc", "")
	if strings.Contains(divClasses, "he") {
		divClasses = strings.ReplaceAll(divClasses, "ic", "i")
	}
	imgClasses := g.parseAttribute(el.Attr, map[string]struct{}{"s": {}, "e": {}})
	imgClasses = strings.ReplaceAll(imgClasses, "hc ", "")
	imgClasses = strings.ReplaceAll(imgClasses, "wc ", "")
	imgClasses = strings.ReplaceAll(imgClasses, " wc", "")
	return html.Div(
		html.Attributes{html.AttributeClass: divClasses},
		html.Img(
			html.Attributes{
				html.AttributeClass: imgClasses,
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

func (g *generator) generateHead() []html.Tag {
	out := []html.Tag{
		html.Link(html.Attributes{
			html.AttributeHref: "public/base.css",
			html.AttributeType: "text/css",
			html.AttributeRel:  "stylesheet",
		}),
	}

	css := make([]string, 0, len(g.css))
	for class := range g.css {
		css = append(css, class)
	}
	sort.Strings(css)

	style := ""

	for _, class := range css {
		switch true {
		case strings.HasPrefix(class, "spacing-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			style += generateSpacing(part[1])
		case strings.HasPrefix(class, "font-size-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			style += generateFontSize(part[2])
		case strings.HasPrefix(class, "p-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				continue
			}
			style += generatePadding(part[1])
		case strings.HasPrefix(class, "width-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				continue
			}
			style += generateWidth(kind, value)
		case strings.HasPrefix(class, "height-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				continue
			}
			style += generateHeight(kind, value)
		case strings.HasPrefix(class, "bg-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class)
			if err != nil {
				continue
			}
			style += fmt.Sprintf(`.%s{
  background-color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "fc-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class)
			if err != nil {
				continue
			}
			style += fmt.Sprintf(`.%s{
  color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		}
	}
	if style != "" {
		out = append(out, html.Style(style))
	}
	return out
}

func generateWidth(kind string, value int64) string {
	switch kind {
	case "fill":
		return fmt.Sprintf(`.s.r > .width-fill-%d{
  flex-grow: %d;
}`, value, value*100000)
	case "px":
		return fmt.Sprintf(`.width-px-%d{
  width: %dpx;
}`, value, value)
	}
	return ""
}

func generateHeight(kind string, value int64) string {
	switch kind {
	case "fill":
		return fmt.Sprintf(`.s.c > .height-fill-%d{
  flex-grow: %d;
}`, value, value*100000)
	case "px":
		return fmt.Sprintf(`.height-px-%d{
  height: %dpx;
}`, value, value)
	}
	return ""
}

func parseBgClass(attribute string) (uint64, uint64, uint64, error) {
	colorAttr := attribute[3 : len(attribute)-4]
	colorPart := strings.Split(colorAttr, "-")
	if len(colorPart) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid number of arg provider expected 3 got %d in %s", len(colorPart), colorAttr)
	}
	r, err := strconv.ParseUint(colorPart[0], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 1 (%s) : %w", colorPart[0], err)
	}
	g, err := strconv.ParseUint(colorPart[1], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 2 (%s) : %w", colorPart[0], err)
	}
	b, err := strconv.ParseUint(colorPart[2], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 3 (%s) : %w", colorPart[0], err)
	}

	return r, g, b, nil
}

func generateFontSize(strSize string) string {
	css := `.font-size-%size%{
  font-size: %size%px;
}`
	size, _ := strconv.ParseInt(strSize, 10, 64)
	if size < 33 {
		return ""
	}
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generatePadding(strSize string) string {
	css := `.p-%size%{
  padding: %size%px %size%px %size%px %size%px;
}`
	size, _ := strconv.ParseInt(strSize, 10, 64)
	if size < 25 {
		return ""
	}
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generateSpacing(strSpace string) string {
	cssSpacing := `.spacing-%spacing%-%spacing%.r > .s + .s{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.wrp.r > .s{
  margin: %spacing-half%px %spacing-half%px;
}.spacing-%spacing%-%spacing%.c > .s + .s{
  margin-top: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .s + .s{
  margin-top: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .al{
  margin-right: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .ar{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.p{
  line-height: calc(1em + %spacing%px);
}textarea.s.spacing-%spacing%-%spacing%{
  line-height: calc(1em + %spacing%px);
  height: calc(100% + %spacing%px);
}.spacing-%spacing%-%spacing%.p > .al{
  margin-right: %spacing%px;
}.spacing-%spacing%-%spacing%.p > .ar{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.p::after{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-top: -%spacing-half%px;
}.spacing-%spacing%-%spacing%.p::before{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-bottom: -%spacing-half%px;
}`
	space, _ := strconv.ParseInt(strSpace, 10, 64)
	cssSpacing = strings.ReplaceAll(cssSpacing, "%spacing%", fmt.Sprintf("%d", space))
	return strings.ReplaceAll(cssSpacing, "%spacing-half%", fmt.Sprintf("%d", space/2))
}
