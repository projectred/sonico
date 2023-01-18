package file

import (
	"context"
	"embed"
	"fmt"
	"testing"

	"github.com/projectred/sonico/walk"
)

//go:embed test
var dir embed.FS

//go:embed test/a.json
var file embed.FS

func TestFS(t *testing.T) {
	walk.NewWalker([]walk.Visitor{
		func(ctx context.Context, stack walk.Nodes, node walk.Node) error {
			fmt.Println(stack.Path(), node.Key(), string(node.Value().([]byte)), "///")
			return nil
		},
	}, nil).Walk(context.Background(), nil, NewFS(dir, "test"))
	walk.NewWalker([]walk.Visitor{
		func(ctx context.Context, stack walk.Nodes, node walk.Node) error {
			fmt.Println(stack.Path(), node.Key(), string(node.Value().([]byte)), "///")
			return nil
		},
	}, nil).Walk(context.Background(), nil, NewFS(file, "test/a.json"))
}
