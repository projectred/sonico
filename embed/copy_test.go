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
	if err := Copy(fs, "test", targetDir, os.O_CREATE|os.O_WRONLY); err != nil {
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

func TestConveyCopy(t *testing.T) {
	targetDir := path.Join(t.TempDir(), "config")
	if err := Copy(fs, "test", targetDir, os.O_CREATE|os.O_WRONLY); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path.Join(targetDir, "test_data.json"), []byte("just a test"), 0666); err != nil {
		t.Fatal(err)
	}
	if err := CoverCopy(fs, "test", targetDir); err != nil {
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

func TestAppendCopy(t *testing.T) {
	targetDir := path.Join(t.TempDir(), "config")
	if err := Copy(fs, "test", targetDir, os.O_CREATE|os.O_WRONLY); err != nil {
		t.Fatal(err)
	}
	if err := AppendCopy(fs, "test", targetDir); err != nil {
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
	for start, end := 0, len(datas)/2; end == len(datas); start, end = end, len(datas) {
		if err := json.Unmarshal(datas[start:end], &outPut); err != nil {
			t.Fatal(err)
		}
		if outPut.Work != "ok" {
			t.Fatalf("it should be 'ok', but it is %s", outPut.Work)
		}
		outPut.Work = ""
	}
}
