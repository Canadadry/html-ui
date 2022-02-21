package html

func basic(name string, attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            name,
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Table(attr Attributes, children ...Node) Tag {
	return basic("table", attr, children...)
}

func Thead(attr Attributes, children ...Node) Tag {
	return basic("thead", attr, children...)
}

func Tbody(attr Attributes, children ...Node) Tag {
	return basic("tbody", attr, children...)
}

func Tr(attr Attributes, children ...Node) Tag {
	return basic("tr", attr, children...)
}

func Td(attr Attributes, children ...Node) Tag {
	return basic("td", attr, children...)
}

func Th(attr Attributes, children ...Node) Tag {
	return basic("th", attr, children...)
}

func P(attr Attributes, children ...Node) Tag {
	return basic("p", attr, children...)
}

func Strong(attr Attributes, children ...Node) Tag {
	return basic("strong", attr, children...)
}

func B(attr Attributes, children ...Node) Tag {
	return basic("b", attr, children...)
}

func Nav(attr Attributes, children ...Node) Tag {
	return basic("nav", attr, children...)
}

func Footer(attr Attributes, children ...Node) Tag {
	return basic("footer", attr, children...)
}

func Div(attr Attributes, children ...Node) Tag {
	return basic("div", attr, children...)
}

func H1(attr Attributes, children ...Node) Tag {
	return basic("h1", attr, children...)
}

func H2(attr Attributes, children ...Node) Tag {
	return basic("h2", attr, children...)
}

func H3(attr Attributes, children ...Node) Tag {
	return basic("h3", attr, children...)
}

func H4(attr Attributes, children ...Node) Tag {
	return basic("h4", attr, children...)
}

func H5(attr Attributes, children ...Node) Tag {
	return basic("h5", attr, children...)
}

func Span(attr Attributes, children ...Node) Tag {
	return basic("span", attr, children...)
}

func I(attr Attributes, children ...Node) Tag {
	return basic("i", attr, children...)
}

func Small(attr Attributes, children ...Node) Tag {
	return basic("small", attr, children...)
}

func Ul(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "ul",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Ol(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "ol",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Li(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "li",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Form(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "form",
		AttributesNames: []attribute{AttributeId, AttributeAction, AttributeMethod, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Input(attr Attributes) Tag {
	return Tag{
		Name:            "input",
		AttributesNames: []attribute{AttributeId, AttributePlaceholder, AttributeType, AttributeName, AttributeValue, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Closed:          true,
	}
}

func Label(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "label",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest, AttributeFor},
		Attributes:      attr,
		Children:        children,
	}
}

func A(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "a",
		AttributesNames: []attribute{AttributeId, AttributeHref, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}
func Button(attr Attributes, children ...Node) Tag {
	return Tag{
		Name:            "button",
		AttributesNames: []attribute{AttributeId, AttributeType, AttributeName, AttributeValue, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Children:        children,
	}
}

func Img(attr Attributes) Tag {
	return Tag{
		Name:            "img",
		AttributesNames: []attribute{AttributeId, AttributeSrc, AttributeClass, AttributeStyle, AttributeDataTest},
		Attributes:      attr,
		Closed:          true,
	}
}

func Br(attr Attributes) Tag {
	return Tag{
		Name:            "br",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Closed:          true,
	}
}

func Hr(attr Attributes) Tag {
	return Tag{
		Name:            "hr",
		AttributesNames: []attribute{AttributeId, AttributeClass, AttributeStyle, AttributeDataTest},
		Closed:          true,
	}
}

func Text(txt string) Tag {
	return Tag{Raw: txt}
}
