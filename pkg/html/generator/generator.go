package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func Generate(in io.Reader, out io.Writer) error {
	h, err := html.Parse(in)
	if err != nil {
		return fmt.Errorf("cannot parse input html : %w", err)
	}
	ast, _ := parseHtmlNode(h)
	if ast == nil {
		return fmt.Errorf("return empty ast")
	}
	ast.Root = true
	buf := bytes.Buffer{}
	err = generate(&buf, *ast)
	if err != nil {
		return fmt.Errorf("cannot generate code : %w", err)
	}
	source := buf.String()
	err = gofmt(source, out)
	if err != nil {
		return fmt.Errorf("cannot format code : %w\n %s", err, source)
	}
	return nil
}

func generate(w io.Writer, a Ast) error {
	_, err := fmt.Fprint(w, `package generated

import (
	"app/pkg/html"
)

func Html() html.Node {
	return `)
	if err != nil {
		return err
	}
	err = a.Print(w)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, `
}`)
	return err
}

func parseHtmlNode(n *html.Node) (*Ast, *html.Node) {
	n = filterDocumentNode(n)
	a := Ast{}
	if n.Type == html.TextNode {
		trimmed := trimWhiteLine(n.Data)
		if trimmed != "" {
			a.Text = trimmed
		}
	} else {
		a.Type = n.Data
		for _, attr := range n.Attr {
			a.Attr = append(a.Attr, Attr{Key: attr.Key, Val: attr.Val})
		}
	}
	if n.FirstChild != nil {
		sibling := n.FirstChild
		var c *Ast
		for sibling != nil {
			c, sibling = parseHtmlNode(sibling)
			if c != nil {
				a.Children = append(a.Children, *c)
			}
		}
	}
	if a.Text == "" && a.Type == "" {
		return nil, n.NextSibling
	}
	return &a, n.NextSibling
}

func filterDocumentNode(n *html.Node) *html.Node {
	if n.Type == html.DocumentNode {
		return filterDocumentNode(n.FirstChild)
	}
	if n.Data == "html" {
		return filterDocumentNode(n.FirstChild)
	}
	if n.Data == "head" {
		return filterDocumentNode(n.NextSibling)
	}
	if n.Data == "body" {
		return filterDocumentNode(n.FirstChild)
	}
	return n
}

type Ast struct {
	Text     string
	Type     string
	Attr     []Attr
	Children []Ast
	Root     bool
}

type Attr struct {
	Key string
	Val string
}

func (a Ast) Print(w io.Writer) error {
	if a.Text != "" {
		_, err := fmt.Fprintf(w, "html.Text(\"%s\"),\n", a.Text)
		if err != nil {
			return err
		}
		return nil
	}
	_, err := fmt.Fprintf(w, "html.%s(Attributes{\n", strings.Title(a.Type))
	if err != nil {
		return err
	}
	for _, attr := range a.Attr {
		_, err = fmt.Fprintf(w, "html.Attribute%s:\"%s\",\n", formatAttributeName(attr.Key), formatAttributeValue(attr.Val))
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, "},\n")
	if err != nil {
		return err
	}
	for _, child := range a.Children {
		err = child.Print(w)
		if err != nil {
			return err
		}
	}
	if !a.Root {
		_, err = fmt.Fprintf(w, "),\n")
		if err != nil {
			return err
		}
	} else {
		_, err = fmt.Fprintf(w, ")")
		if err != nil {
			return err
		}
	}
	return nil
}

func trimWhiteLine(in string) string {
	lines := strings.Split(in, "\n")
	for i := range lines {
		lines[i] = strings.Trim(lines[i], " \t")
	}
	first := 0
	last := len(lines) - 1
	for lines[first] == "" && first < last {
		first++
	}
	for lines[last] == "" && first < last {
		last--
	}
	joined := strings.Join(lines[first:last+1], "\n")
	return strings.ReplaceAll(joined, "\n", "\\n")
}

func formatAttributeName(name string) string {
	parts := strings.Split(name, "-")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	return strings.Join(parts, "")
}
func formatAttributeValue(value string) string {
	return strings.ReplaceAll(value, "\n", "\\n")
}

func gofmt(in string, w io.Writer) error {
	buf, err := format.Source([]byte(in))
	if err != nil {
		return err
	}
	_, err = w.Write(buf)
	return err
}
