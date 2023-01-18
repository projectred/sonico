package sql

import (
	"fmt"
	"reflect"
)

func String(v reflect.Value) string {
	if v.Kind() == reflect.Pointer && v.IsNil() {
		return "null"
	}
	if v.Kind() == reflect.String {
		return fmt.Sprintf("'%v'", v.Interface())
	}
	return fmt.Sprintf("%v", v.Interface())
}
