package embed

import (
	"embed"
	"os"
	"path"
)

func CoverCopy(fs embed.FS, dir, targetPath string) error {
	return Copy(fs, dir, targetPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC)
}

func AppendCopy(fs embed.FS, dir, targetPath string) error {
	return Copy(fs, dir, targetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND)
}

func Copy(fs embed.FS, dir, targetPath string, flag int) error {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(targetPath, 0754); err != nil {
		return err
	}
	for _, file := range files {
		datas, err := fs.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return err
		}

		if err := func() error {
			writer, err := os.OpenFile(path.Join(targetPath, file.Name()), flag, 0644)
			defer writer.Close()
			if err != nil {
				return err
			}
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
