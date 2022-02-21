package parser

import (
	"bytes"
	"testing"
)

func TestElXml(t *testing.T) {
	tests := []struct {
		in  El
		out string
	}{
		{
			in: El{
				Type: "layout",
				Children: []El{
					{
						Type: "row",
						Children: []El{
							{
								Type: "el",
								Children: []El{
									{
										Type:    "text",
										Content: "stylish!",
									},
								},
							},
							{
								Type: "el",
								Children: []El{
									{
										Type:    "text",
										Content: "stylish!",
									},
								},
							},
							{
								Type: "el",
								Children: []El{
									{
										Type: "el",
										Children: []El{
											{
												Type:    "text",
												Content: "stylish!",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			out: `<layout>
	<row>
		<el>
			stylish!
		</el>
		<el>
			stylish!
		</el>
		<el>
			<el>
				stylish!
			</el>
		</el>
	</row>
</layout>`,
		},
	}

	for i, tt := range tests {
		buf := bytes.Buffer{}
		err := tt.in.Xml(&buf, "")
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}
		if buf.String() != tt.out {
			t.Fatalf("[%d] failed \ngot %s\nexp %s\n", i, buf.String(), tt.out)
		}
	}
}
