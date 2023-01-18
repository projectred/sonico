package embed

import (
	"context"
	"embed"
	"os"
	"path"

	"github.com/projectred/sonico/output/file"
	"github.com/projectred/sonico/walk"
)

func CoverCopy(fs embed.FS, dir, targetPath string, rename bool) error {
	return Copy(fs, dir, targetPath, rename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
}

func AppendCopy(fs embed.FS, dir, targetPath string, rename bool) error {
	return Copy(fs, dir, targetPath, rename, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
}

func Copy(fs embed.FS, dir, targetPath string, rename bool, flag int) error {
	return walk.NewWalker([]walk.Visitor{
		func(ctx context.Context, stack walk.Nodes, node walk.Node) error {
			if rename {
				stack = stack[1:]
			}
			if err := os.MkdirAll(path.Join(targetPath, stack.Path()), 0754); err != nil {
				return err
			}
			datas, err := fs.ReadFile(path.Join(dir, stack.Path(), node.Key()))
			if err != nil {
				return err
			}

			writer, err := os.OpenFile(path.Join(targetPath, stack.Path(), node.Key()), flag, 0644)
			defer writer.Close()
			if err != nil {
				return err
			}
			if _, err := writer.Write(datas); err != nil {
				return err
			}
			return nil
		}}, nil).Walk(context.Background(), nil, file.NewFS(fs, dir))
}
