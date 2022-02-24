package html

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Node interface {
	Render(w io.Writer, indent string) error
}

type attribute string
type Attributes map[attribute]string
type Tag struct {
	Raw             string
	Name            string
	Closed          bool
	AttributesNames []attribute
	Attributes      map[attribute]string
	Children        []Tag
	InlineChildren  bool
}

func Render(w io.Writer, t Tag, indent, indentIncr, linebreak string) error {
	if t.Raw != "" {
		rawPart := strings.Split(t.Raw, linebreak)
		for i, part := range rawPart {
			_, err := fmt.Fprintf(w, "%s%s%s", indent, part, linebreak)
			if err != nil {
				return fmt.Errorf("on tag raw part %d:'%s' : %w", i, part, err)
			}
		}
		return nil
	}
	content := t.Name
	sort.Slice(t.AttributesNames, func(i, j int) bool { return t.AttributesNames[i] < t.AttributesNames[j] })
	for _, attrName := range t.AttributesNames {
		value, ok := t.Attributes[attrName]
		if !ok {
			continue
		}
		if value == "" {
			content += " " + string(attrName)
			continue
		}
		content += " " + string(attrName) + `="` + value + `"`
	}
	l := linebreak
	if t.InlineChildren {
		l = ""
	}
	if t.Closed {
		_, err := fmt.Fprintf(w, "%s<%s/>%s", indent, content, l)
		return err
	}
	_, err := fmt.Fprintf(w, "%s<%s>%s", indent, content, l)
	if err != nil {
		return fmt.Errorf("on tag %s before : %w", t.Name, err)
	}
	for i, r := range t.Children {
		if t.InlineChildren {
			err = Render(w, r, "", "", "")
		} else {
			err = Render(w, r, indent+indentIncr, indentIncr, linebreak)
		}
		if err != nil {
			return fmt.Errorf("on tag %s child %d : %w", t.Name, i, err)
		}
	}
	i := indent
	if t.InlineChildren {
		i = ""
	}
	_, err = fmt.Fprintf(w, "%s</%s>%s", i, t.Name, linebreak)
	if err != nil {
		return fmt.Errorf("on tag %s after : %w", t.Name, err)
	}
	return err
}
