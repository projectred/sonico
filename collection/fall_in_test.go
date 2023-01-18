package collection

import (
	"reflect"
	"testing"
)

func TestFallIn(t *testing.T) {
	if !reflect.DeepEqual(
		FallIn([]string{"id", "name"}, "`", "`"),
		[]string{"`id`", "`name`"},
	) {
		t.FailNow()
	}
}
