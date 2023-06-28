package chef

import (
	"fmt"
	"reflect"
)

type Name struct {
	value string
}

func NewName(n string) (*Name, error) {
	if n == "" {
		return nil, fmt.Errorf("name is required")
	}
	name := new(Name)
	name.value = n
	return name, nil
}

// Equals 等価性の比較、今回は値が一つだが複数の場合はすべての値の等価性を比較する
func (name *Name) Equals(otherName Name) bool {
	return reflect.DeepEqual(name.value, otherName.value)
}
