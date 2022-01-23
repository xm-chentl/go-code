package metadata

import (
	"reflect"
	"sync"
)

var pool sync.Map

func Get(entry interface{}) ITable {
	entryType := reflect.TypeOf(entry)
	if entryType.Kind() == reflect.Ptr {
		entryType = entryType.Elem()
	}

	value, ok := pool.Load(entryType.Name())
	if ok {
		return value.(ITable)
	}

	columns := make([]IColumn, 0)
	for i := 0; i < entryType.NumField(); i++ {
		columns = append(columns, &column{
			structField: entryType.Field(i),
		})
	}
	tableInst := &table{
		entryType: entryType,
		columns:   columns,
	}
	pool.Store(entryType.Name(), tableInst)

	return tableInst
}
