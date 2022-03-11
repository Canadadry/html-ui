package ast

import (
	"fmt"
	"io"
)

type ElType string

const (
	indent = "\t"

	TypeElText        ElType = "text"
	TypeElEl                 = "el"
	TypeElRow                = "row"
	TypeElColumn             = "column"
	TypeElLayout             = "layout"
	TypeElImage              = "img"
	TypeElButton             = "button"
	TypeElForm               = "form"
	TypeElInput              = "input"
	TypeElLabel              = "label"
	TypeElPlaceholder        = "placeholder"
)

var ValidElType = map[ElType]struct{}{
	TypeElText:        {},
	TypeElEl:          {},
	TypeElRow:         {},
	TypeElColumn:      {},
	TypeElLayout:      {},
	TypeElImage:       {},
	TypeElButton:      {},
	TypeElForm:        {},
	TypeElInput:       {},
	TypeElLabel:       {},
	TypeElPlaceholder: {},
}

type El struct {
	Type     ElType
	Attr     []Attribute
	Children []El
	Content  string
}

func (el *El) GetAttr(t AttrType) (Attribute, bool) {
	for _, attr := range el.Attr {
		if attr.Type == t {
			return attr, true
		}
	}
	return Attribute{}, false
}

func (el *El) GetChild(t ElType) (El, bool) {
	for _, c := range el.Children {
		if c.Type == t {
			return c, true
		}
	}
	return El{}, false
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
