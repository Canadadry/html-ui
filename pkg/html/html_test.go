package html

import (
	"bytes"
	"testing"
)

func TestDocumentWith(t *testing.T) {
	tests := []struct {
		r   Document
		exp string
	}{
		{
			r: Document{
				Lang: "en",
				Head: []Node{
					MetaCharset("utf-8"),
					Meta("viewport", "width=device-width, initial-scale=1"),
					Link(Attributes{AttributeHref: "/public/css/style.css", AttributeType: "text/css", AttributeRel: "stylesheet"}),
					Link(Attributes{AttributeHref: "/public/img/favicon.ico", AttributeRel: "icon"}),
					Title("example"),
				},
				Body: []Node{
					Div(Attributes{AttributeClass: "container-fluid"},
						Div(Attributes{AttributeClass: "row"},
							Div(Attributes{AttributeClass: "col pt-2"},
								A(Attributes{AttributeHref: "/"},
									Img(Attributes{AttributeSrc: "/public/img/logo.png"}),
									Text("Hello world!"),
								),
							),
						),
					),
				},
			},
			exp: `<!doctype html>
<html lang="en">
	<head>
		<meta charset="utf-8"/>
		<meta content="width=device-width, initial-scale=1" name="viewport"/>
		<link href="/public/css/style.css" rel="stylesheet" type="text/css"/>
		<link href="/public/img/favicon.ico" rel="icon"/>
		<title>
			example
		</title>
	</head>
	<body>
		<div class="container-fluid">
			<div class="row">
				<div class="col pt-2">
					<a href="/">
						<img src="/public/img/logo.png"/>
						Hello world!
					</a>
				</div>
			</div>
		</div>
	</body>
</html>
`,
		},
	}

	for i, tt := range tests {
		result := bytes.Buffer{}
		err := tt.r.Render(&result)
		if err != nil {
			t.Fatalf("[%d] Failed : %v", i, err)
		}
		if result.String() != tt.exp {
			t.Fatalf("[%d] got \n%s\nexp\n%s\n", i, result.String(), tt.exp)
		}
	}
}
