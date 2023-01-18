package sql

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/projectred/sonico/collection"
	"github.com/projectred/sonico/walk"
)

func SQLInsert(ctx context.Context, source interface{}, table, key, tag string) ([]byte, error) {
	var fields []string
	var values [][]string
	if err := walk.NewWalker([]walk.Visitor{
		func(ctx context.Context, stack walk.Nodes, node walk.Node) error {
			if len(values) == 1 {
				fields = append(fields, node.Key())
			}
			values[len(values)-1] = append(values[len(values)-1], String(reflect.ValueOf(node.Value())))
			return nil
		},
	},
		[]walk.Visitor{
			func(ctx context.Context, stack walk.Nodes, node walk.Node) error {
				if reflect.ValueOf(node.Value()).Kind() != reflect.Slice {
					values = append(values, []string{})
				}
				return nil
			},
		},
	).Walk(ctx, nil, walk.NewReflect(source, key, tag)); err != nil {
		return nil, err
	}
	buffer := bytes.NewBufferString(fmt.Sprintf("INSERT INTO %s (%s) VALUES ", table, (strings.Join(collection.FallIn(fields, "`", "`"), ", "))))

	l := len(values[0])
	for i := range values {
		values[0] = append(values[0], strings.Join(values[i], ", "))
	}
	buffer.WriteString(strings.Join(collection.FallIn(values[0][l:], "(", ")"), ", ") + ";")
	return buffer.Bytes(), nil
}
