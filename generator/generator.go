package generator

import (
	"app/ast"
	"app/pkg/html"
	"io"
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
	g.css["spacing-10-10"] = struct{}{}
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: "hc spacing-10-10 s c wc ct cl"},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateRow(el ast.El) html.Tag {
	g.css["spacing-10-10"] = struct{}{}
	g.mode = modeColumn
	return html.Div(
		html.Attributes{html.AttributeClass: "hc spacing-10-10 s r wc cl ccy"},
		g.generate(el.Children)...,
	)
}

func (g *generator) generateEl(el ast.El) html.Tag {
	if len(el.Children) == 0 {
		return html.Tag{}
	}
	if el.Children[0].Type != ast.TypeElText {
		return html.Tag{}
	}
	return g.generateText(el.Children[0].Content)
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
			html.AttributeHref: "base.css",
			html.AttributeType: "text/css",
			html.AttributeRel:  "stylesheet",
		}),
	}

	for _ = range g.css {
		out = append(out, html.Style(`.spacing-10-10.r > .s + .s{
  margin-left: 10px;
}.spacing-10-10.wrp.r > .s{
  margin: 5px 5px;
}.spacing-10-10.c > .s + .s{
  margin-top: 10px;
}.spacing-10-10.pg > .s + .s{
  margin-top: 10px;
}.spacing-10-10.pg > .al{
  margin-right: 10px;
}.spacing-10-10.pg > .ar{
  margin-left: 10px;
}.spacing-10-10.p{
  line-height: calc(1em + 10px);
}textarea.s.spacing-10-10{
  line-height: calc(1em + 10px);
  height: calc(100% + 10px);
}.spacing-10-10.p > .al{
  margin-right: 10px;
}.spacing-10-10.p > .ar{
  margin-left: 10px;
}.spacing-10-10.p::after{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-top: -5px;
}.spacing-10-10.p::before{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-bottom: -5px;
}`))
	}
	return out
}
