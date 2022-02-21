package html

import (
	"io"
)

type If struct {
	Cond bool
	Node Node
}

func (i If) Render(w io.Writer, indent string) error {
	if i.Cond {
		return i.Node.Render(w, indent)
	}
	return nil
}
