package generator

import (
	"app/ast"
	"app/pkg/html"
	"io"
)

func Generate(el ast.El, w io.Writer) error {
	d := html.Document{
		Head: []html.Tag{
			html.Link(html.Attributes{
				html.AttributeHref: "base.css",
				html.AttributeType: "text/css",
				html.AttributeRel:  "stylesheet",
			}),
		},
		Body: generate([]ast.El{el}),
	}
	return d.Render(w)
}

func generate(el []ast.El) []html.Tag {
	out := make([]html.Tag, 0, len(el))
	for _, item := range el {
		switch item.Type {
		case ast.TypeElLayout:
			out = append(out, generateLayout(item))
		case ast.TypeElText:
			out = append(out, generateText(item.Content))
		}
	}
	return out
}

func generateLayout(el ast.El) html.Tag {
	return html.Div(
		html.Attributes{html.AttributeClass: "bg-255-255-255-0 fc-0-0-0-255 font-size-20 font-open-sanshelveticaverdanasans-serif s e ui s e"},
		generate(el.Children)...,
	)
}

func generateText(txt string) html.Tag {
	return html.Inline(html.Div(html.Attributes{html.AttributeClass: "s t wf hf"},
		html.Text(txt),
	))
}
