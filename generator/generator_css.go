package generator

import (
	"app/pkg/html"
	"fmt"
	"strconv"
	"strings"
)

func generateHead(cssClasses UniqueClasses) []html.Tag {
	out := []html.Tag{
		html.Link(html.Attributes{
			html.AttributeHref: "public/base.css",
			html.AttributeType: "text/css",
			html.AttributeRel:  "stylesheet",
		}),
	}

	css := cssClasses.Sorted()
	style := ""

	for _, class := range css {
		switch true {
		case strings.HasPrefix(class, "spacing-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			style += generateSpacing(part[1])
		case strings.HasPrefix(class, "font-size-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			style += generateFontSize(part[2])
		case strings.HasPrefix(class, "p-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				continue
			}
			style += generatePadding(part[1])
		case strings.HasPrefix(class, "b-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				continue
			}
			style += generateBorder(part[1])
		case strings.HasPrefix(class, "br-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				continue
			}
			style += generateBorderRounded(part[1])
		case strings.HasPrefix(class, "width-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				continue
			}
			style += generateWidth(kind, value)
		case strings.HasPrefix(class, "height-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				continue
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				continue
			}
			style += generateHeight(kind, value)
		case strings.HasPrefix(class, "bg-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class)
			if err != nil {
				continue
			}
			style += fmt.Sprintf(`.%s{
  background-color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "fc-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class)
			if err != nil {
				continue
			}
			style += fmt.Sprintf(`.%s{
  color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "bc-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class)
			if err != nil {
				continue
			}
			style += fmt.Sprintf(`.%s{
  border-color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		}
	}
	if style != "" {
		out = append(out, html.Style(style))
	}
	return out
}

func generateWidth(kind string, value int64) string {
	switch kind {
	case "fill":
		return fmt.Sprintf(`.s.r > .width-fill-%d{
  flex-grow: %d;
}`, value, value*100000)
	case "px":
		return fmt.Sprintf(`.width-px-%d{
  width: %dpx;
}`, value, value)
	}
	return ""
}

func generateHeight(kind string, value int64) string {
	switch kind {
	case "fill":
		return fmt.Sprintf(`.s.c > .height-fill-%d{
  flex-grow: %d;
}`, value, value*100000)
	case "px":
		return fmt.Sprintf(`.height-px-%d{
  height: %dpx;
}`, value, value)
	}
	return ""
}

func parseBgClass(attribute string) (uint64, uint64, uint64, error) {
	colorAttr := attribute[3 : len(attribute)-4]
	colorPart := strings.Split(colorAttr, "-")
	if len(colorPart) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid number of arg provider expected 3 got %d in %s", len(colorPart), colorAttr)
	}
	r, err := strconv.ParseUint(colorPart[0], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 1 (%s) : %w", colorPart[0], err)
	}
	g, err := strconv.ParseUint(colorPart[1], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 2 (%s) : %w", colorPart[0], err)
	}
	b, err := strconv.ParseUint(colorPart[2], 10, 8)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid parsing of argument 3 (%s) : %w", colorPart[0], err)
	}

	return r, g, b, nil
}

func generateFontSize(strSize string) string {
	css := `.font-size-%size%{
  font-size: %size%px;
}`
	size, _ := strconv.ParseInt(strSize, 10, 64)
	if size < 33 {
		return ""
	}
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generatePadding(strSize string) string {
	css := `.p-%size%{
  padding: %size%px %size%px %size%px %size%px;
}`
	size, _ := strconv.ParseInt(strSize, 10, 64)
	if size < 25 {
		return ""
	}
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generateBorder(strSize string) string {
	css := `.b-%size%{
  border-width: %size%px %size%px %size%px %size%px;
}`
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generateBorderRounded(strSize string) string {
	css := `.br-%size%{
  border-radius: %size%px;
}`
	return strings.ReplaceAll(css, "%size%", strSize)
}

func generateSpacing(strSpace string) string {
	cssSpacing := `.spacing-%spacing%-%spacing%.r > .s + .s{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.wrp.r > .s{
  margin: %spacing-half%px %spacing-half%px;
}.spacing-%spacing%-%spacing%.c > .s + .s{
  margin-top: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .s + .s{
  margin-top: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .al{
  margin-right: %spacing%px;
}.spacing-%spacing%-%spacing%.pg > .ar{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.p{
  line-height: calc(1em + %spacing%px);
}textarea.s.spacing-%spacing%-%spacing%{
  line-height: calc(1em + %spacing%px);
  height: calc(100% + %spacing%px);
}.spacing-%spacing%-%spacing%.p > .al{
  margin-right: %spacing%px;
}.spacing-%spacing%-%spacing%.p > .ar{
  margin-left: %spacing%px;
}.spacing-%spacing%-%spacing%.p::after{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-top: -%spacing-half%px;
}.spacing-%spacing%-%spacing%.p::before{
  content: '';
  display: block;
  height: 0;
  width: 0;
  margin-bottom: -%spacing-half%px;
}`
	space, _ := strconv.ParseInt(strSpace, 10, 64)
	cssSpacing = strings.ReplaceAll(cssSpacing, "%spacing%", fmt.Sprintf("%d", space))
	return strings.ReplaceAll(cssSpacing, "%spacing-half%", fmt.Sprintf("%d", space/2))
}
