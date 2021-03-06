package generator

import (
	"app/pkg/html"
	"fmt"
	"strconv"
	"strings"
)

func generateHead(cssClasses UniqueClasses, fonts fontDefiniton) ([]html.Tag, error) {
	out := []html.Tag{
		html.Link(html.Attributes{
			html.AttributeHref: "public/base.css",
			html.AttributeType: "text/css",
			html.AttributeRel:  "stylesheet",
		}),
	}

	css := cssClasses.Sorted()
	style := ""

	for name, url := range fonts {
		style += generateFontDefinition(string(name), string(url))
	}

	predefinedCssClasses := UniqueClassesFrom("wf hf")
	for _, class := range css {
		switch true {
		case predefinedCssClasses.Has(class):
			continue
		case strings.HasPrefix(class, "ff-"):
			continue
		case strings.HasPrefix(class, "box-"):
			part := strings.Split(class, "-")
			if len(part) != 8 && len(part) != 9 {
				return nil, fmt.Errorf("invalid shadow css class %s with %d part", class, len(part))
			}
			style += generateShadow(class, part[1], part[2], part[3], part[4], part[5], part[6], part[7], len(part) == 9)
			continue
		case strings.HasPrefix(class, "spacing-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				return nil, fmt.Errorf("invalid spacing css class %s", class)
			}
			style += generateSpacing(part[1])
		case strings.HasPrefix(class, "font-size-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				return nil, fmt.Errorf("invalid font-size css class %s", class)
			}
			style += generateFontSize(part[2])
		case strings.HasPrefix(class, "p-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				return nil, fmt.Errorf("invalid p css class %s", class)
			}
			style += generatePadding(part[1])
		case strings.HasPrefix(class, "b-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				return nil, fmt.Errorf("invalid b css class %s", class)
			}
			style += generateBorder(part[1])
		case strings.HasPrefix(class, "br-"):
			part := strings.Split(class, "-")
			if len(part) != 2 {
				return nil, fmt.Errorf("invalid br css class %s", class)
			}
			style += generateBorderRounded(part[1])
		case strings.HasPrefix(class, "width-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				return nil, fmt.Errorf("invalid width css class %s", class)
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid width css class %s: %w", class, err)
			}
			style += generateWidth(kind, value)
		case strings.HasPrefix(class, "height-"):
			part := strings.Split(class, "-")
			if len(part) != 3 {
				return nil, fmt.Errorf("invalid height css class %s", class)
			}
			kind := part[1]
			value, err := strconv.ParseInt(part[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid height css class %s: %w", class, err)
			}
			style += generateHeight(kind, value)
		case strings.HasPrefix(class, "bg-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class, 3, 4)
			if err != nil {
				return nil, fmt.Errorf("invalid bg css class %s : %w", class, err)
			}
			style += fmt.Sprintf(`.%s{
  background-color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "fc-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class, 3, 4)
			if err != nil {
				return nil, fmt.Errorf("invalid cf css class %s: %w", class, err)
			}
			style += fmt.Sprintf(`.%s{
  color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "bc-") && strings.HasSuffix(class, "-255"):
			r, g, b, err := parseBgClass(class, 3, 4)
			if err != nil {
				return nil, fmt.Errorf("invalid bg css class %s: %w", class, err)
			}
			style += fmt.Sprintf(`.%s{
  border-color: rgba(%d,%d,%d,1);
}`, class, r, g, b)
		case strings.HasPrefix(class, "bg-") && strings.HasSuffix(class, "-255-fs"):
			r, g, b, err := parseBgClass(class, 3, 7)
			if err != nil {
				return nil, fmt.Errorf("invalid bg css class %s: %w", class, err)
			}
			style += fmt.Sprintf(`.%s:focus {
  background-color: rgba(%d,%d,%d,1);
}.s:focus .%s  {
  background-color: rgba(%d,%d,%d,1);
}.%s:focus-within {
  background-color: rgba(%d,%d,%d,1);
}.ui-slide-bar:focus + .s .focusable-thumb.%s {
  background-color: rgba(%d,%d,%d,1);
}`, class, r, g, b, class, r, g, b, class, r, g, b, class, r, g, b)
		case strings.HasPrefix(class, "w"):
			continue
		default:
			return nil, fmt.Errorf("cannot generate css class for  %s", class)
		}
	}
	if style != "" {
		out = append(out, html.Style(style))
	}
	return out, nil
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

func parseBgClass(attribute string, prefixSize, suffixSize int) (uint64, uint64, uint64, error) {
	colorAttr := attribute[prefixSize : len(attribute)-suffixSize]
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

func generateFontDefinition(name, url string) string {
	style := `@import url('%url%');
.ff-%lowerName%sans-serif .font-open-sanshelveticaverdanasans-serif.cap, .ff-%lowerName%sans-serif .font-open-sanshelveticaverdanasans-serif .cap {line-height: 1;}
.ff-%lowerName%sans-serif .font-open-sanshelveticaverdanasans-serif.cap> .t, .ff-%lowerName%sans-serif .font-open-sanshelveticaverdanasans-serif .cap > .t {vertical-align: 0;line-height: 1;}
.font-open-sanshelveticaverdanasans-serif .ff-%lowerName%sans-serif.cap, .font-open-sanshelveticaverdanasans-serif .ff-%lowerName%sans-serif .cap {line-height: 1;}
.font-open-sanshelveticaverdanasans-serif .ff-%lowerName%sans-serif.cap> .t, .font-open-sanshelveticaverdanasans-serif .ff-%lowerName%sans-serif .cap > .t {vertical-align: 0;line-height: 1;}
.ff-%lowerName%sans-serif.cap, .ff-%lowerName%sans-serif .cap {line-height: 1;}
.ff-%lowerName%sans-serif.cap> .t, .ff-%lowerName%sans-serif .cap > .t {vertical-align: 0;line-height: 1;}
.ff-%lowerName%sans-serif{
  font-family: "%name%", sans-serif;
  font-feature-settings: ;
  font-variant: normal;
}`
	style = strings.ReplaceAll(style, "%name%", name)
	style = strings.ReplaceAll(style, "%lowerName%", strings.ToLower(name))
	style = strings.ReplaceAll(style, "%url%", url)
	return style
}

func generateShadow(class, ox, oy, blur, size, r, g, b string, inner bool) string {
	styleLeft := `.%s{
  box-shadow: `
	styleRight := `%spx %spx %spx %spx rgba(%s,%s,%s,1);
}`
	style := styleLeft + styleRight
	if inner {
		style = styleLeft + "inset " + styleRight
	}
	return fmt.Sprintf(style, class, ox, oy, blur, size, r, g, b)
}
