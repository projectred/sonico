package filter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	if !reflect.DeepEqual(
		Filter([]string{"a", "u", "c"}, func(s string) bool { return s < "u" }),
		[]string{"a", "c"},
	) {
		t.FailNow()
	}
}
