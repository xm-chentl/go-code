package metadata

import "reflect"

type column struct {
	structField reflect.StructField
}

func (c column) Field() string {
	return c.structField.Name
}

func (c column) GetValue(entryValue reflect.Value) interface{} {
	return entryValue.FieldByName(c.Field()).Interface()
}
