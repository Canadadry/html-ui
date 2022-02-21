package ast

import (
	"bytes"
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
