package html

const (
	AttributeXmlns   attribute = "xmlns"
	AttributeHeight  attribute = "height"
	AttributeVersion attribute = "version"
	AttributeViewBox attribute = "viewBox"
	AttributeD       attribute = "d"
)

func Svg(attr Attributes, children ...Tag) Tag {
	return Tag{
		Name:            "svg",
		AttributesNames: []attribute{AttributeXmlns},
		Children:        children,
	}
}
func G(attr Attributes, children ...Tag) Tag {
	return Tag{
		Name:            "g",
		AttributesNames: []attribute{AttributeId},
		Children:        children,
	}
}

func Path(attr Attributes, children ...Tag) Tag {
	return Tag{
		Name:            "path",
		AttributesNames: []attribute{AttributeD, AttributeStyle},
		Children:        children,
	}
}
