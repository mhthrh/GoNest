package csv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/mhthrh/common-lib/pkg/util/directory"
	"os"
	"path/filepath"
)

var (
	FilePath = ""
	FileName = ""
	appPath  = ""
)

type File struct {
	Data []byte
}

func init() {
	appPath = directory.GetAppRootDir()
}
func New(data []byte) *File {
	return &File{
		Data: data,
	}
}

func (f *File) Read() error {

	var buf bytes.Buffer
	file, e := os.Open(filepath.Join(appPath, FilePath, FileName))
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
