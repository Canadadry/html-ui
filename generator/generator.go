package generator

import (
	"app/ast"
	"app/pkg/html"
	"io"
)

func Generate(el ast.El, w io.Writer) error {
	d := html.Document{
		Head: []html.Node{
			html.Link(html.Attributes{
				html.AttributeHref: "base.css",
				html.AttributeType: "text/css",
				html.AttributeRel:  "stylesheet",
			}),
		},
		Body: []html.Node{html.Div(
			html.Attributes{html.AttributeClass: "bg-255-255-255-0 fc-0-0-0-255 font-size-20 font-open-sanshelveticaverdanasans-serif s e ui s e"},
			html.Div(html.Attributes{html.AttributeClass: "s t wf hf"},
				html.Text("Hello world!"),
			),
		)},
	}
	return d.Render(w)
}
