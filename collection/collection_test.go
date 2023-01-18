package collection

import (
	"reflect"
	"testing"
)

type item struct {
	ID   int64
	Name string
}

func TestCollection(t *testing.T) {
	if !reflect.DeepEqual(
		Collection([]*item{{1, "apple"}, {2, "banana"}, {3, "orange"}}, func(i *item) string { return i.Name }),
		[]string{"apple", "banana", "orange"},
	) {
		t.FailNow()
	}
}
