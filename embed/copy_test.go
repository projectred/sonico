package embed

import (
	"embed"
	"encoding/json"
	"os"
	"path"
	"testing"
)

//go:embed test/*.json
var fs embed.FS

func TestCopy(t *testing.T) {
	targetDir := path.Join(t.TempDir(), "config")
	if err := Copy(fs, "test", targetDir, true); err != nil {
		t.Fatal(err)
	}
	files, err := os.ReadDir(targetDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("is should has len 1, but it is %d", len(files))
	}
	if files[0].Name() != "test_data.json" {
		t.Fatalf("it's should be 'test_data.json', but it is %s", files[0].Name())
	}
	datas, err := os.ReadFile(path.Join(targetDir, files[0].Name()))
	if err != nil {
		t.Fatal(err)
	}

	outPut := struct {
		Work string `json:"work"`
	}{}
	if err := json.Unmarshal(datas, &outPut); err != nil {
		t.Fatal(err)
	}
	if outPut.Work != "ok" {
		t.Fatalf("it should be 'ok', but it is %s", outPut.Work)
	}
}
