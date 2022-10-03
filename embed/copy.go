package embed

import (
	"embed"
	"os"
	"path"
)

func Copy(fs embed.FS, dir, targetPath string, cover bool) error {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(targetPath, 0754); err != nil {
		return err
	}
	flag := func() int {
		if cover {
			return os.O_CREATE | os.O_WRONLY
		}
		return os.O_EXCL | os.O_WRONLY
	}()
	for _, file := range files {
		datas, err := fs.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return err
		}

		if err := func() error {
			writer, err := os.OpenFile(path.Join(targetPath, file.Name()), flag, 0666)
			if err == os.ErrExist && !cover {
				return nil
			}
			if err != nil {
				return err
			}
			defer writer.Close()
			if _, err := writer.Write(datas); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return err
		}
	}
	return nil
}
