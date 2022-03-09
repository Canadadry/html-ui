package parser

import (
	"app/ast"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type Parser struct {
	d *xml.Decoder
}

func (p *Parser) Parse(r io.Reader) (ast.El, error) {
	p.d = xml.NewDecoder(r)
	tok, err := p.d.Token()
	if err != nil {
		return ast.El{}, err
	}
	return p.parseItem(tok)
}

func (p *Parser) parseItem(tok xml.Token) (ast.El, error) {
	cd, ok := tok.(xml.CharData)
	if ok {
		return ast.El{
			Type:    ast.TypeElText,
			Content: strings.Trim(string(cd), " \n\r\t"),
		}, nil
	}
	se, ok := tok.(xml.StartElement)
	if !ok {
		return ast.El{}, fmt.Errorf("expected to got a xml.StartElement got a %T", tok)
	}
	_, ok = ast.ValidElType[ast.ElType(se.Name.Local)]
	if !ok {
		return ast.El{}, fmt.Errorf("expected to got a Valid El Type got a '%s', possible values are %v", se.Name.Local, ast.ValidElType)
	}
	el := ast.El{
		Type: ast.ElType(se.Name.Local),
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
			return ast.El{}, err
		}
		ee, ok = tok.(xml.EndElement)
		if ok {
			break
		}
		child, err := p.parseItem(tok)
		if err != nil {
			return el, err
		}
		if child.Type != ast.TypeElText || child.Content != "" {
			el.Children = append(el.Children, child)
		}
	}
	if ee.Name.Local != se.Name.Local {
		return el, fmt.Errorf("expected closing tag %s got %s", se.Name.Local, ee.Name.Local)
	}
	return el, nil
}

func parseAttributes(se xml.StartElement) ([]ast.Attribute, error) {
	attrs := []ast.Attribute{}
	for _, attr := range se.Attr {
		_, ok := ast.ValidAttrType[ast.AttrType(attr.Name.Local)]
		if !ok {
			return nil, fmt.Errorf("while parsing el type '%s' expected to got a valid attr type got a '%s', possible values are %v", se.Name.Local, attr.Name.Local, ast.ValidAttrType)
		}
		att := ast.Attribute{Type: ast.AttrType(attr.Name.Local), Value: attr.Value}
		err := att.Parse()
		if err != nil {
			return nil, err
		}
		attrs = append(attrs, att)
	}
	return attrs, nil
}
