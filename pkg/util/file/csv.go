package csvFile

import (
	"bytes"
	"common-lib/pkg/util/directory"
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Path string
	Data []byte
}

func NewFile(path, name string, data []byte) *File {
	return &File{
		Name: name,
		Path: path,
		Data: data,
	}
}

func (f *File) Read() error {
	appPath, e := directory.GetWorkingDir()
	if e != nil {
		return e
	}
	var buf bytes.Buffer
	file, e := os.Open(filepath.Join(appPath, f.Path, f.Name))
	if e != nil {
		return e
	}
	defer func() {
		_ = file.Close()
	}()

	r := csv.NewReader(file)

	c, e := r.ReadAll()
	if e != nil {
		return e
	}

	e = csv.NewWriter(&buf).WriteAll(c)

	if e != nil {
		return errors.New("error writing CSV")

	}

	f.Data = buf.Bytes()
	return nil
}

func (f *File) Write() error {
	return nil
}
