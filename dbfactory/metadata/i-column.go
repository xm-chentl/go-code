package metadata

import "reflect"

type IColumn interface {
	Field() string
	GetValue(reflect.Value) interface{}
}
