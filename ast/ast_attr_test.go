package ast

import (
	"bytes"
	"image/color"
	"testing"
)

func TestAttributesXml(t *testing.T) {
	tests := []struct {
		in  Attribute
		out string
	}{
		{
			in:  Attribute{Type: TypeAttrWidth, Value: "fill"},
			out: ` width="fill"`,
		},
	}

	for i, tt := range tests {
		buf := bytes.Buffer{}
		err := tt.in.Xml(&buf)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if buf.String() != tt.out {
			t.Fatalf("[%d] failed \ngot '%s'\nexp '%s'\n", i, buf.String(), tt.out)
		}
	}
}

func TestAttributesParse(t *testing.T) {
	tests := []struct {
		in  Attribute
		out Attribute
	}{
		{
			in:  Attribute{Type: TypeAttrWidth, Value: "fill"},
			out: Attribute{Type: TypeAttrWidth, Value: "fill", Size: SizeFill{}},
		},
		{
			in:  Attribute{Type: TypeAttrHeight, Value: "fill"},
			out: Attribute{Type: TypeAttrHeight, Value: "fill", Size: SizeFill{}},
		},
		{
			in:  Attribute{Type: TypeAttrAlign, Value: "left,centerY"},
			out: Attribute{Type: TypeAttrAlign, Value: "left,centerY", AlignX: AlignLeft, AlignY: AlignCenterY},
		},
		{
			in:  Attribute{Type: TypeAttrSpacing, Value: "12"},
			out: Attribute{Type: TypeAttrSpacing, Value: "12", Number: 12},
		},
		{
			in:  Attribute{Type: TypeAttrPadding, Value: "12"},
			out: Attribute{Type: TypeAttrPadding, Value: "12", Number: 12},
		},
		{
			in:  Attribute{Type: TypeAttrBgColor, Value: "rgb(128,255,0)"},
			out: Attribute{Type: TypeAttrBgColor, Value: "rgb(128,255,0)", Color: color.RGBA{128, 255, 0, 0}},
		},
		{
			in:  Attribute{Type: TypeAttrFontColor, Value: "rgb(128,255,0)"},
			out: Attribute{Type: TypeAttrFontColor, Value: "rgb(128,255,0)", Color: color.RGBA{128, 255, 0, 0}},
		},
		{
			in:  Attribute{Type: TypeAttrBorderColor, Value: "rgb(128,255,0)"},
			out: Attribute{Type: TypeAttrBorderColor, Value: "rgb(128,255,0)", Color: color.RGBA{128, 255, 0, 0}},
		},
		{
			in:  Attribute{Type: TypeAttrFontSize, Value: "12"},
			out: Attribute{Type: TypeAttrFontSize, Value: "12", Number: 12},
		},
		{
			in:  Attribute{Type: TypeAttrBorderRounded, Value: "12"},
			out: Attribute{Type: TypeAttrBorderRounded, Value: "12", Number: 12},
		},
		{
			in:  Attribute{Type: TypeAttrBorderWidth, Value: "12"},
			out: Attribute{Type: TypeAttrBorderWidth, Value: "12", Number: 12},
		},
		{
			in:  Attribute{Type: TypeAttrSrc, Value: "test"},
			out: Attribute{Type: TypeAttrSrc, Value: "test"},
		},
		{
			in:  Attribute{Type: TypeAttrAlt, Value: "test"},
			out: Attribute{Type: TypeAttrAlt, Value: "test"},
		},
	}

	for i, tt := range tests {
		err := tt.in.Parse()
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if tt.in != tt.out {
			t.Fatalf("[%d] failed \ngot '%#v'\nexp '%#v'\n", i, tt.in, tt.out)
		}

	}
}
