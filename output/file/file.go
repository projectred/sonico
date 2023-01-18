package file

import (
	"context"
	"embed"
	"io/fs"
	"path"

	"github.com/projectred/sonico/walk"
)

type FS struct {
	file              embed.FS
	name              string
	dirDetailFileName string
}

func NewFS(file embed.FS, name string) *FS {
	return &FS{file: file, name: name, dirDetailFileName: "self.json"}
}

func (f FS) Children(ctx context.Context) []walk.Node {
	if !f.isDir() {
		return nil
	}
	files, err := f.file.ReadDir(f.name)
	if err != nil {
		return nil
	}
	nodes := make([]walk.Node, len(files))
	for i := range files {
		nodes[i] = &FSChild{f, nil, files[i], nil}
	}
	return nodes
}

func (f FS) isDir() bool {
	file, err := f.file.Open(f.name)
	if err != nil {
		return true
	}
	stat, err := file.Stat()
	if err != nil {
		return true
	}
	return stat.IsDir()
}

func (f FS) Key() string { return f.name }
func (f FS) Value() interface{} {
	if f.isDir() {
		return nil
	}
	body, err := f.file.ReadFile(f.name)
	if err != nil {
		return nil
	}
	return body
}

type FSChild struct {
	FS
	stack   walk.Nodes
	current fs.DirEntry
	values  []byte
}

func (f *FSChild) Children(ctx context.Context) []walk.Node {
	if !f.current.IsDir() {
		return nil
	}
	files, err := f.file.ReadDir(path.Join(f.FS.name, f.stack.Path(), f.current.Name()))
	if err != nil {
		return nil
	}
	nodes := make([]walk.Node, 0, len(files))
	for i := range files {
		if files[i].Name() == f.FS.dirDetailFileName {
			f.values, err = f.FS.file.ReadFile(path.Join(f.FS.name, f.stack.Path(), files[i].Name()))
			continue
		}
		nodes = append(nodes, &FSChild{f.FS, append(f.stack, f), files[i], nil})
	}
	return nodes
}

func (f FSChild) Key() string { return f.current.Name() }
func (f *FSChild) Value() interface{} {
	if f.values != nil {
		return f.values
	}
	body, err := f.file.ReadFile(path.Join(f.FS.name, f.stack.Path(), f.current.Name()))
	if err != nil {
		return nil
	}
	f.values = body
	return f.values
}
