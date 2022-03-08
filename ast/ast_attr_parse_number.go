package ast

import (
	"strconv"
)

func ParseNumberAttr(v string) (int, error) {
	num, err := strconv.ParseInt(v, 10, 64)
	return int(num), err
}
