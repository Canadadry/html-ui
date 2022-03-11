package html

import (
	"fmt"
	"io"
)

const (
	AttributeLang            attribute = "lang"
	AttributeId              attribute = "id"
	AttributeCharset         attribute = "charset"
	AttributeName            attribute = "name"
	AttributeContent         attribute = "content"
	AttributeHref            attribute = "href"
	AttributeType            attribute = "type"
	AttributeRel             attribute = "rel"
	AttributeIntegrity       attribute = "integrity"
	AttributeCrossorigin     attribute = "crossorigin"
	AttributeClass           attribute = "class"
	AttributeStyle           attribute = "style"
	AttributeAction          attribute = "action"
	AttributeMethod          attribute = "method"
	AttributePlaceholder     attribute = "placeholder"
	AttributeValue           attribute = "value"
	AttributeSrc             attribute = "src"
	AttributeDataTest        attribute = "data-test"
	AttributeFor             attribute = "for"
	AttributeAsync           attribute = "async"
	AttributeRole            attribute = "role"
	AttributeAlt             attribute = "alt"
	AttributeAria            attribute = "aria"
	AttributeAriaDescribedby attribute = "aria-describedby"
	AttributeTitle           attribute = "title"
	AttributeColspan         attribute = "colspan"
	AttributeTarget          attribute = "target"
	AttributeSpellCheck      attribute = "spellcheck"
)

type Document struct {
	Lang string
	Head []Tag
	Body []Tag
}

func (d Document) Render(w io.Writer) error {
	_, err := fmt.Fprintln(w, "<!doctype html>")
	if err != nil {
		return fmt.Errorf("on doctype : %w", err)
	}
	htmlAttr := map[attribute]string{}
	if d.Lang != "" {
		htmlAttr[AttributeLang] = d.Lang
	}
	t := Tag{
		Name:            "html",
		AttributesNames: []attribute{AttributeLang},
		Attributes:      htmlAttr,
		Children: []Tag{
			Tag{
				Name:     "head",
				Children: d.Head,
			},
			Tag{
				Name:     "body",
				Children: d.Body,
			},
		},
	}
	return Render(w, t, "", "\t", "\n")
}
