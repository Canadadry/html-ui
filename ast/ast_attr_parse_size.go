package ast

import (
	"strconv"
	"strings"
)

type AttrSize interface {
	Get() int
	Type() SizeType
}

type SizeType string

const (
	SizePxType      SizeType = "SizePxType"
	SizePortionType          = "SizePortionType"
	SizeFillType             = "SizeFillType"
)

type SizePx int

func (s SizePx) Get() int       { return int(s) }
func (s SizePx) Type() SizeType { return SizePxType }

type SizePortion int

func (s SizePortion) Get() int       { return int(s) }
func (s SizePortion) Type() SizeType { return SizePortionType }

type SizeFill struct{}

func (SizeFill) Get() int       { return 0 }
func (SizeFill) Type() SizeType { return SizeFillType }

func ParseSizeAttr(width string) (AttrSize, error) {
	if strings.HasPrefix(width, "px:") {
		num, err := strconv.ParseInt(width[3:], 10, 64)
		return SizePx(num), err
	}
	if strings.HasPrefix(width, "portion:") {
		num, err := strconv.ParseInt(width[8:], 10, 64)
		return SizePortion(num), err
	}
	return SizeFill{}, nil
}
