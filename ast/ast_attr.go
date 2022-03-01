package ast

import (
	"fmt"
	"io"
)

type AttrType string

const (
	TypeAttrWidth         AttrType = "width"
	TypeAttrHeight                 = "height"
	TypeAttrAlign                  = "align"
	TypeAttrSpacing                = "spacing"
	TypeAttrPadding                = "padding"
	TypeAttrBgColor                = "bg-color"
	TypeAttrFontColor              = "font-color"
	TypeAttrFontSize               = "font-size"
	TypeAttrBorderRounded          = "border-rounded"
	TypeAttrBorderColor            = "border-color"
	TypeAttrBorderWidth            = "border-width"
	TypeAttrSrc                    = "src"
	TypeAttrAlt                    = "alt"
)

var ValidAttrType = map[AttrType]struct{}{
	TypeAttrWidth:         {},
	TypeAttrHeight:        {},
	TypeAttrAlign:         {},
	TypeAttrSpacing:       {},
	TypeAttrPadding:       {},
	TypeAttrBgColor:       {},
	TypeAttrFontColor:     {},
	TypeAttrFontSize:      {},
	TypeAttrBorderWidth:   {},
	TypeAttrBorderColor:   {},
	TypeAttrBorderRounded: {},
	TypeAttrSrc:           {},
	TypeAttrAlt:           {},
}

type Attribute struct {
	Type  AttrType
	Value string
}

func (att Attribute) Xml(w io.Writer) error {
	_, err := fmt.Fprintf(w, " %v=\"%s\"", att.Type, att.Value)
	return err
}
