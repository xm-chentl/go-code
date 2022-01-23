package metadata

import (
	"reflect"
	"testing"
)

func Test_Column_Field(t *testing.T) {
	testTableInst := testTable{}
	rt := reflect.TypeOf(testTableInst)
	field, _ := rt.FieldByName("FieldOne")
	columnInst := column{
		structField: field,
	}
	if columnInst.Field() != field.Name {
		t.Error("err")
	}
}
