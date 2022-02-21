package html

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	tests := []struct {
		t Tag
		v string
	}{
		{
			t: Tag{
				Name:            "a",
				AttributesNames: []attribute{attribute("href"), attribute("target")},
				Attributes: map[attribute]string{
					attribute("href"):   "/public/content.html",
					attribute("target"): "_blank",
					attribute("fake"):   "attribute",
				},
				Children: []Node{Tag{Raw: "TestChild"}},
			},
			v: `<a href="/public/content.html" target="_blank">
	TestChild
</a>
`,
		},
		{
			t: Tag{
				Name:            "a",
				AttributesNames: []attribute{attribute("href"), attribute("target")},
				Attributes: map[attribute]string{
					attribute("href"):   "/public/content.html",
					attribute("target"): "_blank",
					attribute("fake"):   "attribute",
				},
				Children: []Node{Tag{
					Name:            "a",
					AttributesNames: []attribute{attribute("href"), attribute("target")},
					Attributes: map[attribute]string{
						attribute("href"):   "/public/content.html",
						attribute("target"): "_blank",
						attribute("fake"):   "attribute",
					},
					Children: []Node{Tag{Raw: "Test\nChild"}},
				}},
			},
			v: `<a href="/public/content.html" target="_blank">
	<a href="/public/content.html" target="_blank">
		Test
		Child
	</a>
</a>
`,
		},
		{
			t: Tag{
				Name: "input",
				AttributesNames: []attribute{
					attribute("type"),
					attribute("placeholder"),
					attribute("value"),
				},
				Attributes: map[attribute]string{
					attribute("type"):        "text",
					attribute("placeholder"): "your name",
				},
				Closed:   true,
				Children: []Node{Tag{Raw: "TestChild"}},
			},
			v: `<input placeholder="your name" type="text"/>
`,
		},
		{
			t: Tag{
				Name: "script",
				AttributesNames: []attribute{
					attribute("type"),
					attribute("async"),
				},
				Attributes: map[attribute]string{
					attribute("type"):  "text",
					attribute("async"): "",
				},
				Closed: true,
			},
			v: `<script async type="text"/>
`,
		},
	}

	for i, tt := range tests {
		result := bytes.Buffer{}
		err := tt.t.Render(&result, "")
		if err != nil {
			t.Fatalf("[%d] Failed : %v", i, err)
		}
		if result.String() != tt.v {
			t.Fatalf("[%d] got \n%s\nexp\n%s\n", i, result.String(), tt.v)
		}
	}
}
