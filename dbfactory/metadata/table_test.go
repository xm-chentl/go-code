package metadata

import (
	"reflect"
	"testing"
)

type testTable struct {
	FieldOne string
	FieldTwo int
}

func Test_GetValueMap(t *testing.T) {
	testTableInst := testTable{
		FieldOne: "field-one",
		FieldTwo: 2,
	}

	columns := make([]IColumn, 0)
	rt := reflect.TypeOf(testTableInst)
	for i := 0; i < rt.NumField(); i++ {
		columns = append(columns, &column{
			structField: rt.Field(i),
		})
	}

	tableInst := table{
		entryType: reflect.TypeOf(testTableInst),
		columns:   columns,
	}
	data := tableInst.GetValueMap(testTableInst)
	if len(data) != 2 {
		t.Error("err")
	}

	v, ok := data["FieldOne"]
	if !ok || v.(string) != "field-one" {
		t.Error("err")
	}

	v, ok = data["FieldTwo"]
	if !ok || v.(int) != 2 {
		t.Error("err")
	}
}
