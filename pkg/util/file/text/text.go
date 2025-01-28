package text

import (
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
	data, e := os.ReadFile(filepath.Join(appPath, FilePath, FileName))
	if e != nil {
		return e
	}

	f.Data = data
	return nil
}

func (f *File) Write() error {
	return nil
}
