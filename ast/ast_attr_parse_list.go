package ast

import (
	"fmt"
)

func ParseListAttr(v string, list []string) error {
	for _, item := range list {
		if item == v {
			return nil
		}
	}
	return fmt.Errorf("%s is not a valid value : can have %v", v, list)
}
