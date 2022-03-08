package ast

import (
	"fmt"
	"image/color"
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
	Type   AttrType
	Value  string
	Number int
	Color  color.RGBA
	Size   AttrSize
	AlignX AlignXType
	AlignY AlignYType
}

func (att *Attribute) Parse() error {
	var err error
	switch att.Type {
	case TypeAttrWidth:
		att.Size, err = ParseSizeAttr(att.Value)
	case TypeAttrHeight:
		att.Size, err = ParseSizeAttr(att.Value)
	case TypeAttrAlign:
		att.AlignX, att.AlignY, err = ParseAlignAttr(att.Value)
	case TypeAttrSpacing:
		att.Number, err = ParseNumberAttr(att.Value)
	case TypeAttrPadding:
		att.Number, err = ParseNumberAttr(att.Value)
	case TypeAttrBgColor:
		att.Color, err = ParseColorAttr(att.Value)
	case TypeAttrFontColor:
		att.Color, err = ParseColorAttr(att.Value)
	case TypeAttrFontSize:
		att.Number, err = ParseNumberAttr(att.Value)
	case TypeAttrBorderRounded:
		att.Number, err = ParseNumberAttr(att.Value)
	case TypeAttrBorderColor:
		att.Color, err = ParseColorAttr(att.Value)
	case TypeAttrBorderWidth:
		att.Number, err = ParseNumberAttr(att.Value)
	case TypeAttrSrc:
	case TypeAttrAlt:
	default:
		err = fmt.Errorf("unhandled case type %s", att.Type)
	}
	return err
}

func (att Attribute) Xml(w io.Writer) error {
	_, err := fmt.Fprintf(w, " %v=\"%s\"", att.Type, att.Value)
	return err
}
