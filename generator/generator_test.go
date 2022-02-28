package generator

import (
	"app/parser"
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []string{
		"test1",
		// "test2",
		// "test3",
		// "test4",
	}

	for _, tt := range tests {
		fIn, err := os.Open("testcase/" + tt + ".xml")
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}
		defer fIn.Close()
		p := parser.Parser{}
		ast, err := p.Parse(fIn)
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}
		result := bytes.Buffer{}
		err = Generate(ast, &result)
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}

		fOut, err := os.Open("testcase/" + tt + ".html")
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}
		defer fOut.Close()

		expected, err := ioutil.ReadAll(fOut)
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}

		if result.String() != string(expected) {
			t.Fatalf("[%s] failed \ngot \n-%s-\nexp \n-%s-\n", tt, result.String(), string(expected))
		}
	}
}
