package html

func MetaCharset(charset string) Tag {
	return Tag{
		Name:            "meta",
		AttributesNames: []attribute{AttributeCharset},
		Attributes:      Attributes{AttributeCharset: charset},
		Closed:          true,
	}
}

func Meta(name, content string) Tag {
	return Tag{
		Name:            "meta",
		AttributesNames: []attribute{AttributeName, AttributeContent},
		Attributes: Attributes{
			AttributeName:    name,
			AttributeContent: content,
		},
		Closed: true,
	}
}

func Link(attr Attributes) Tag {
	return Tag{
		Name:            "link",
		AttributesNames: []attribute{AttributeHref, AttributeType, AttributeRel, AttributeIntegrity, AttributeCrossorigin},
		Attributes:      attr,
		Closed:          true,
	}
}

func Title(title string) Tag {
	return Tag{
		Name:     "title",
		Children: []Tag{Text(title)},
	}
}

func Script(attr Attributes, script string) Tag {
	return Tag{
		Name:            "script",
		AttributesNames: []attribute{AttributeType, AttributeSrc},
		Attributes:      attr,
		Children:        []Tag{Text(script)},
	}
}

func Style(style string) Tag {
	return Tag{
		Name:            "style",
		AttributesNames: []attribute{},
		Attributes:      nil,
		Children:        []Tag{Text(style)},
	}
}
