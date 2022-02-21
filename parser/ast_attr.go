package parser

import (
	"fmt"
	"io"
)

type AttrType string

const (
	TypeAttrWidth         AttrType = "width"
	TypeAttrAlign                  = "align"
	TypeAttrSpacing                = "spacing"
	TypeAttrPadding                = "padding"
	TypeAttrBgColor                = "bg-color"
	TypeAttrFontColor              = "font-color"
	TypeAttrBorderRounded          = "border-rounded"
)

var ValidAttrType = map[AttrType]struct{}{
	TypeAttrWidth:         struct{}{},
	TypeAttrAlign:         struct{}{},
	TypeAttrSpacing:       struct{}{},
	TypeAttrPadding:       struct{}{},
	TypeAttrBgColor:       struct{}{},
	TypeAttrFontColor:     struct{}{},
	TypeAttrBorderRounded: struct{}{},
}

type Attribute struct {
	Type  AttrType
	Value string
}

func (att Attribute) Xml(w io.Writer) error {
	_, err := fmt.Fprintf(w, " %v=\"%s\"", att.Type, att.Value)
	return err
}
