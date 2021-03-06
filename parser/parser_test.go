package parser

import (
	"bytes"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []string{`<layout>
	<row width="fill" align="centerY" spacing="30" padding="30" bg-color="rgb(0,0,245)">
		<el bg-color="rgb(240,0,245)" font-color="rgb(255,255,255)" border-rounded="3" padding="30">
			stylish!
		</el>
		<el bg-color="rgb(240,0,245)" font-color="rgb(255,255,255)" border-rounded="3" padding="30">
			stylish!
		</el>
		<el align="right">
			<el bg-color="rgb(240,0,245)" font-color="rgb(255,255,255)" border-rounded="3" padding="30">
				stylish!
			</el>
		</el>
	</row>
</layout>`, `<layout>
	<column spacing="10">
		<el>
			hello world!
		</el>
		<el>
			hello world!
		</el>
		<el>
			hello world!
		</el>
	</column>
</layout>`, `<layout>
	<font name="Poppins" src="https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600&amp;display=swap">
	</font>
</layout>`,
	}

	for i, tt := range tests {
		r := strings.NewReader(tt)
		p := &Parser{}
		el, err := p.Parse(r)
		if err != nil {
			t.Fatalf("[%d] parsing failed %v", i, err)
		}
		buf := bytes.Buffer{}
		err = el.Xml(&buf, "")
		if err != nil {
			t.Fatalf("[%d] rendering failed %v", i, err)
		}
		if buf.String() != tt {
			t.Fatalf("[%d] failed \ngot %s\nexp %s\n", i, buf.String(), tt)
		}
	}
}
