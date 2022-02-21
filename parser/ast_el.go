package parser

import (
	"fmt"
	"io"
)

type ElType string

const (
	indent = "\t"

	TypeElText   ElType = "text"
	TypeElEl            = "el"
	TypeElRow           = "row"
	TypeElLayout        = "layout"
)

var ValidElType = map[ElType]struct{}{
	TypeElText:   struct{}{},
	TypeElEl:     struct{}{},
	TypeElRow:    struct{}{},
	TypeElLayout: struct{}{},
}

type El struct {
	Type     ElType
	Attr     []Attribute
	Children []El
	Content  string
}

func (el *El) Xml(w io.Writer, prefix string) error {
	if el.Type == TypeElText {
		_, err := fmt.Fprintf(w, "%s%s", prefix, el.Content)
		return err
	}
	_, err := fmt.Fprintf(w, "%s<%s", prefix, el.Type)
	if err != nil {
		return err
	}
	for _, attr := range el.Attr {
		err = attr.Xml(w)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, ">\n")
	if err != nil {
		return err
	}
	for _, child := range el.Children {
		err = child.Xml(w, prefix+indent)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintf(w, "\n")
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, "%s</%s>", prefix, el.Type)
	if err != nil {
		return err
	}
	return nil
}
