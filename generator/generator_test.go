package generator

import (
	"app/ast"
	"app/parser"
	"bytes"
	"github.com/sergi/go-diff/diffmatchpatch"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
		"test6",
		"test7",
		"test8",
		"test9",
		"test10",
		"test11",
		"test12",
		"test13",
		"test14",
		"test15",
		"test16",
		"test17",
		"test18_part1",
		"test18_part2",
		"test18_part3",
		"test18_part4",
		"test19",
		"test20",
		"test21",
	}

	for _, tt := range tests {
		fIn, err := os.Open("testcase/" + tt + ".xml")
		if err != nil {
			t.Fatalf("[%s] failed %v", tt, err)
		}
		defer fIn.Close()
		p := parser.Parser{}
		root, err := p.Parse(fIn)
		if err != nil {
			t.Fatalf("[%s] failed parsing %v", tt, err)
		}
		err = ast.Validate(root)
		if err != nil {
			t.Fatalf("[%s] failed validating %v", tt, err)
		}

		result := bytes.Buffer{}
		err = Generate(&result, root)
		if err != nil {
			t.Fatalf("[%s] failed generating %v", tt, err)
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
			t.Errorf("[%s] failed \ngot \n-%s-\nexp \n-%s-\n", tt, result.String(), string(expected))
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(result.String(), string(expected), false)
			if len(diffs) > 1 {
				t.Fatalf("[%s] failed \n%v", tt, dmp.DiffPrettyText(diffs))
			}
		}
	}
}
