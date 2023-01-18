package walk

import (
	"context"
	"path"

	"github.com/projectred/sonico/collection"
)

type Node interface {
	Children(ctx context.Context) []Node
	Key() string
	Value() interface{}
}

type Nodes []Node

func (ns Nodes) Path() string {
	return path.Join(collection.Collection(ns, func(n Node) string { return n.Key() })...)
}

type WriteAbleNode interface {
	Node
	AppendNode() WriteAbleNode
	WriteAbleChildren() []WriteAbleNode
}

type Visitor func(ctx context.Context, stack Nodes, node Node) error

type Walker struct {
	leafVisitors, nonLeafVisitors []Visitor
}

func NewWalker(leafVisitors, nonLeafVisitors []Visitor) *Walker {
	return &Walker{leafVisitors: leafVisitors, nonLeafVisitors: nonLeafVisitors}
}

func (w *Walker) Walk(ctx context.Context, stack Nodes, node Node) error {
	if len(node.Children(ctx)) == 0 {
		for i := range w.leafVisitors {
			if err := w.leafVisitors[i](ctx, stack, node); err != nil {
				return err
			}
		}
		return nil
	}
	for i := range w.nonLeafVisitors {
		if err := w.nonLeafVisitors[i](ctx, stack, node); err != nil {
			return err
		}
	}
	for _, child := range node.Children(ctx) {
		if err := w.Walk(ctx, append(stack, node), child); err != nil {
			return err
		}
	}
	return nil
}
