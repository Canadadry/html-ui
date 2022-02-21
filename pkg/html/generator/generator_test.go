package generator

import (
	"bytes"
	"strings"
	"testing"
)

func TestTrimWhiteLine(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "test",
			out: "test",
		},
		{
			in:  "",
			out: "",
		},
		{
			in:  "\ntest\n",
			out: "test",
		},
		{
			in:  " \n \ttest\n",
			out: "test",
		},
		{
			in:  "\nte\nst\n",
			out: "te\\nst",
		},
		{
			in:  "\n\nte\nst\n\n",
			out: "te\\nst",
		},
	}

	for i, tt := range tests {
		result := trimWhiteLine(tt.in)
		if result != tt.out {
			t.Fatalf("[%d] got \n%s\nexp\n%s\n", i, result, tt.out)
		}
	}
}

type TestCase struct {
	in  string
	out string
}

func TestGenerate(t *testing.T) {
	tests := testCase()

	for i, tt := range tests {
		buf := bytes.Buffer{}
		err := Generate(strings.NewReader(tt.in), &buf)
		if err != nil {
			t.Fatalf("%d failed %v", i, err)
		}
		outFormated := bytes.Buffer{}
		err = gofmt(tt.out, &outFormated)
		if err != nil {
			t.Fatalf("%d cannot format output %v", i, err)
		}
		diff(t, i, buf.String(), outFormated.String())
	}
}

func diff(t *testing.T, i int, left string, right string) {
	t.Helper()
	line := 0
	if left == right {
		return
	}
	min := len(left) - 1
	if min > len(right)-1 {
		min = len(right) - 1
	}
	for j := 0; j < min; j++ {
		if left[j] == '\n' {
			line++
		}
		if left[j] != right[j] {
			t.Fatalf("[%d] diff at char %d:%d got \n%s\nexp\n%s\n", i, line, j, left, right)
		}
	}
	t.Fatalf("[%d] size not equal got %d want %d \ngot \n-%s-\nexp\n-%s-\n", i, len(left), len(right), left, right)
}

func testCase() []TestCase {
	return []TestCase{
		{
			in: `<a href="/public/content.html" target="_blank">
	TestChild
</a>`,
			out: `package generated

import (
	"app/pkg/html"
)

func Html() html.Node {
	return html.A(Attributes{
		html.AttributeHref:   "/public/content.html",
		html.AttributeTarget: "_blank",
	},
		html.Text("TestChild"),
	)
}
`,
		},
		{
			in: `<a href="/public/content.html" target="_blank">
<span class="fa fa-users">
    Test
    Child
</span>
</a>
`,
			out: `package generated

import (
	"app/pkg/html"
)

func Html() html.Node {
	return html.A(Attributes{
		html.AttributeHref:   "/public/content.html",
		html.AttributeTarget: "_blank",
	},
		html.Span(Attributes{
			html.AttributeClass:   "fa fa-users",
		},
			html.Text("Test\nChild"),
		),
	)
}`,
		},
		{
			in: `<a href="mailto:test@gmail.com?subject=test
newline" target="_blank">
<span class="fa fa-users">
    Test
    Child
</span>
</a>
`,
			out: `package generated

import (
	"app/pkg/html"
)

func Html() html.Node {
	return html.A(Attributes{
		html.AttributeHref:   "mailto:test@gmail.com?subject=test\nnewline",
		html.AttributeTarget: "_blank",
	},
		html.Span(Attributes{
			html.AttributeClass:   "fa fa-users",
		},
			html.Text("Test\nChild"),
		),
	)
}`,
		},
	}
}
