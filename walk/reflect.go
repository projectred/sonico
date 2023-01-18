package walk

import (
	"context"
	"reflect"
	"strconv"
)

func elemType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

func elemValue(t reflect.Value) reflect.Value {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

type Reflect struct {
	self interface{}
	key  string
	tag  string
}

func NewReflect(self interface{}, key string, tag string) *Reflect {
	return &Reflect{self: self, key: key, tag: tag}
}

func (r Reflect) Children(ctx context.Context) []Node {
	t := elemType(reflect.TypeOf(r.self))
	v := elemValue(reflect.ValueOf(r.self))
	if t.Kind() == reflect.Struct {
		result := make([]Node, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			result[i] = Reflect{v.Field(i).Interface(), t.Field(i).Tag.Get(r.tag), r.tag}
		}
		return result
	}
	if t.Kind() == reflect.Slice {
		result := make([]Node, v.Len())
		for i := 0; i < v.Len(); i++ {
			result[i] = Reflect{v.Index(i).Interface(), strconv.FormatInt(int64(i), 10), r.tag}
		}
		return result
	}
	return nil
}
func (r Reflect) Key() string        { return r.key }
func (r Reflect) Value() interface{} { return r.self }
