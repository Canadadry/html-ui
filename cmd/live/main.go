package main

import (
	"app/ast"
	"app/generator"
	"app/parser"
	"app/pkg/router"
	"embed"
	"flag"
	"fmt"
	"net/http"
	"os"
)

//go:embed public
var publicFiles embed.FS

func main() {
	if err := run(); err != nil {
		fmt.Println("failed", err)
		os.Exit(1)
	}
}

func run() error {
	var in string

	flag.StringVar(&in, "in", "", "input filename")
	flag.Parse()

	r := &router.Group{}
	r.Mount("/public/", http.FileServer(http.FS(publicFiles)))
	r.Get("/", handler(in))

	server := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	return server.ListenAndServe()
}

func handler(in string) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		fIn, err := os.Open(in)
		if err != nil {
			return fmt.Errorf("cannot open input file : %w", err)
		}
		defer fIn.Close()

		p := parser.Parser{}
		root, err := p.Parse(fIn)
		if err != nil {
			return err
		}
		err = ast.Validate(root)
		if err != nil {
			return err
		}
		return generator.Generate(w, root)
	}
}
