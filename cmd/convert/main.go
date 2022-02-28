package main

import (
	"app/generator"
	"app/parser"
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("failed", err)
		os.Exit(1)
	}
}

func run() error {
	var in, out string

	flag.StringVar(&in, "in", "", "input filename")
	flag.StringVar(&out, "out", "out.html", "output filename")

	flag.Parse()

	fIn, err := os.Open(in)
	if err != nil {
		return fmt.Errorf("cannot open input file : %w", err)
	}
	defer fIn.Close()

	fOut, err := os.OpenFile(out, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("cannot open out file : %w", err)
	}
	defer fIn.Close()

	p := parser.Parser{}
	ast, err := p.Parse(fIn)
	if err != nil {
		return err
	}

	return generator.Generate(ast, fOut)
}
