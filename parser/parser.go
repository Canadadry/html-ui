package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	d *xml.Decoder
}

func (p *Parser) Parse(r io.Reader) (El, error) {
	p.d = xml.NewDecoder(r)
	tok, err := p.d.Token()
	if err != nil {
		return El{}, err
	}
	return p.parseItem(tok)
}

func (p *Parser) parseItem(tok xml.Token) (El, error) {
	cd, ok := tok.(xml.CharData)
	if ok {
		return El{
			Type:    TypeElText,
			Content: strings.Trim(string(cd), " \n\r\t"),
		}, nil
	}
	se, ok := tok.(xml.StartElement)
	if !ok {
		return El{}, fmt.Errorf("expected to got a xml.StartElement got a %T", tok)
	}
	_, ok = ValidElType[ElType(se.Name.Local)]
	if !ok {
		return El{}, fmt.Errorf("expected to got a Valid El Type got a '%s', possible values are %v", se.Name.Local, ValidElType)
	}
	el := El{
		Type: ElType(se.Name.Local),
	}
	attrs, err := parseAttributes(se)
	if err != nil {
		return el, err
	}
	el.Attr = attrs
	var ee xml.EndElement
	for {
		tok, err := p.d.Token()
		if err != nil {
			return El{}, err
		}
		ee, ok = tok.(xml.EndElement)
		if ok {
			break
		}
		child, err := p.parseItem(tok)
		if err != nil {
			return el, err
		}
		if child.Type != TypeElText || child.Content != "" {
			el.Children = append(el.Children, child)
		}
	}
	if ee.Name.Local != se.Name.Local {
		return el, fmt.Errorf("expected closing tag %s got %s", se.Name.Local, ee.Name.Local)
	}
	return el, nil
}

func parseAttributes(se xml.StartElement) ([]Attribute, error) {
	attrs := []Attribute{}
	for _, attr := range se.Attr {
		_, ok := ValidAttrType[AttrType(attr.Name.Local)]
		if !ok {
			return nil, fmt.Errorf("expected to got a Valid Attr Type got a '%s', possible values are %v", attr.Name.Local, ValidAttrType)
		}
		attrs = append(attrs, Attribute{Type: AttrType(attr.Name.Local), Value: attr.Value})
	}
	return attrs, nil
}
