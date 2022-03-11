package generator

import (
	"app/ast"
	"app/pkg/html"
	"fmt"
	"io"
)

func Generate(w io.Writer, el ast.El) error {
	gen := generator{
		css: map[string]struct{}{},
	}
	body, err := gen.generate([]ast.El{el})
	if err != nil {
		return err
	}
	head, err := generateHead(gen.css)
	if err != nil {
		return err
	}
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

func (g *generator) generate(el []ast.El) ([]html.Tag, error) {
	out := make([]html.Tag, 0, len(el))
	for _, item := range el {
		var err error
		var child html.Tag
		switch item.Type {
		case ast.TypeElLayout:
			child, err = g.generateLayout(item)
		case ast.TypeElColumn:
			child, err = g.generateColumn(item)
		case ast.TypeElRow:
			child, err = g.generateRow(item)
		case ast.TypeElEl:
			child, err = g.generateEl(item)
		case ast.TypeElImage:
			child, err = g.generateImage(item)
		case ast.TypeElButton:
			child, err = g.generateButton(item)
		case ast.TypeElText:
			child, err = g.generateText(item.Content)
		case ast.TypeElForm:
			child, err = g.generateForm(item)
		default:
			return nil, fmt.Errorf("cannot generate '%s' : unknown type", item.Type)
		}
		if err != nil {
			return nil, err
		}
		out = append(out, child)
	}
	return out, nil
}

func (g *generator) generateLayout(el ast.El) (html.Tag, error) {
	children, err := g.generate(el.Children)
	return html.Div(
		html.Attributes{html.AttributeClass: "bg-255-255-255-0 fc-0-0-0-255 font-size-20 font-open-sanshelveticaverdanasans-serif s e ui s e"},
		children...,
	), err
}

func (g *generator) generateColumn(el ast.El) (html.Tag, error) {
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s c ct cl"))
	if err != nil {
		return html.Tag{}, err
	}
	g.mode = modeColumn
	children, err := g.generate(el.Children)
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		children...,
	), err
}

func (g *generator) generateRow(el ast.El) (html.Tag, error) {
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s r cl ccy"))
	if err != nil {
		return html.Tag{}, err
	}

	g.mode = modeColumn
	children, err := g.generate(el.Children)
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		children...,
	), err
}

func (g *generator) generateEl(el ast.El) (html.Tag, error) {
	if len(el.Children) > 0 && el.Children[0].Type == ast.TypeElText {
		return g.generateElText(el)
	}
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s e"))
	if err != nil {
		return html.Tag{}, err
	}
	g.mode = modeNormal
	children, err := g.generate(el.Children)
	tag := html.Div(
		html.Attributes{html.AttributeClass: classes},
		children...,
	)
	if len(el.Children) == 0 {
		return html.Inline(tag), err
	}
	return tag, err
}

func (g *generator) generateElText(el ast.El) (html.Tag, error) {
	if len(el.Attr) == 0 {
		return g.generateText(el.Children[0].Content)
	}
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s e"))
	if err != nil {
		return html.Tag{}, err
	}
	g.mode = modeNormal
	children, err := g.generateText(el.Children[0].Content)
	return html.Div(
		html.Attributes{html.AttributeClass: classes},
		children,
	), err
}

func (g *generator) generateImage(el ast.El) (html.Tag, error) {
	srcAttr, _ := el.GetAttr(ast.TypeAttrSrc)
	altAttr, _ := el.GetAttr(ast.TypeAttrAlt)
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s e ic"))
	if err != nil {
		return html.Tag{}, err
	}
	divClasses := UniqueClassesFrom(classes)
	divClasses.Remove("hc")
	divClasses.Remove("wc")
	if divClasses.Has("he") {
		divClasses.Remove("ic")
		divClasses.Add("i")
	}
	classes, err = g.parseAttribute(el.Attr, UniqueClassesFrom("s e"))
	if err != nil {
		return html.Tag{}, err
	}
	imgClasses := UniqueClassesFrom(classes)
	imgClasses.Remove("hc")
	imgClasses.Remove("wc")
	return html.Div(
		html.Attributes{html.AttributeClass: divClasses.String()},
		html.Img(
			html.Attributes{
				html.AttributeClass: imgClasses.String(),
				html.AttributeSrc:   srcAttr.Value,
				html.AttributeAlt:   altAttr.Value,
			},
		),
	), nil
}

func (g *generator) generateText(txt string) (html.Tag, error) {
	class := "s t wf hf"
	if g.mode == modeColumn {
		class = "s t wc hc"
	}
	return html.Inline(html.Div(html.Attributes{html.AttributeClass: class},
		html.Text(txt),
	)), nil
}

func (g *generator) generateButton(el ast.El) (html.Tag, error) {
	nameAttr, _ := el.GetAttr(ast.TypeAttrName)
	valueAttr, _ := el.GetAttr(ast.TypeAttrValue)
	classes, err := g.parseAttribute(el.Attr, UniqueClassesFrom("s e ccx ccy cptr hc notxt sbt"))
	if err != nil {
		return html.Tag{}, err
	}
	children, err := g.generate(el.Children)
	return html.Button(
		html.Attributes{
			html.AttributeClass: classes,
			html.AttributeName:  nameAttr.Value,
			html.AttributeValue: valueAttr.Value,
		},
		children...,
	), err
}

func (g *generator) generateForm(el ast.El) (html.Tag, error) {
	nameAttr, _ := el.GetAttr(ast.TypeAttrName)
	actionAttr, _ := el.GetAttr(ast.TypeAttrAction)
	methodAttr, _ := el.GetAttr(ast.TypeAttrMethod)
	children, err := g.generate(el.Children)
	return html.Form(
		html.Attributes{
			html.AttributeAction: actionAttr.Value,
			html.AttributeName:   nameAttr.Value,
			html.AttributeMethod: methodAttr.Value,
		},
		children...,
	), err
}
