package ast

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
	TypeAttrFontSize               = "font-size"
	TypeAttrBorderRounded          = "border-rounded"
)

var ValidAttrType = map[AttrType]struct{}{
	TypeAttrWidth:         {},
	TypeAttrAlign:         {},
	TypeAttrSpacing:       {},
	TypeAttrPadding:       {},
	TypeAttrBgColor:       {},
	TypeAttrFontColor:     {},
	TypeAttrFontSize:      {},
	TypeAttrBorderRounded: {},
}

type Attribute struct {
	Type  AttrType
	Value string
}

func (att Attribute) Xml(w io.Writer) error {
	_, err := fmt.Fprintf(w, " %v=\"%s\"", att.Type, att.Value)
	return err
}
