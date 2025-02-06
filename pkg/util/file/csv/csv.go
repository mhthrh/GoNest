package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/mhthrh/GoNest/pkg/util/directory"
	"github.com/mhthrh/GoNest/pkg/util/file"
	"os"
	"path/filepath"
)

var (
	appPath = ""
)

func init() {
	appPath = directory.GetAppRootDir()
}

type File struct {
	path string
	name string
}

func init() {
	appPath = directory.GetAppRootDir()
}
func New(path, name string) file.IFile {
	return &File{
		path: path,
		name: name,
	}
}

func (f *File) Read() ([]byte, error) {
	var buf bytes.Buffer
	file, e := os.Open(filepath.Join(appPath, f.path, f.name))
	if e != nil {
		return nil, e
	}
	defer func() {
		_ = file.Close()
	}()

	r := csv.NewReader(file)

	c, e := r.ReadAll()
	if e != nil {
		return nil, e
	}

	e = csv.NewWriter(&buf).WriteAll(c)

	if e != nil {
		return nil, errors.New("error writing CSV")

	}

	return buf.Bytes(), nil
}

func (f *File) Write(i []byte) error {
	//TODO implement me
	panic("implement me")
}
