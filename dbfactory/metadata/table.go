package metadata

import "reflect"

type table struct {
	entryType reflect.Type
	columns   []IColumn
}

func (t table) GetValueMap(value interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	for _, column := range t.columns {
		res[column.Field()] = column.GetValue(rv)
	}

	return
}
