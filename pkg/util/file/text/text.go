package text

import (
	"github.com/mhthrh/common-lib/pkg/util/directory"
	"github.com/mhthrh/common-lib/pkg/util/file"
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

func New(path, name string) file.IFile {
	return &File{
		path: path,
		name: name,
	}
}
func (f *File) Read() ([]byte, error) {
	data, e := os.ReadFile(filepath.Join(appPath, f.path, f.name))
	if e != nil {
		return nil, e
	}
	return data, nil
}

func (f *File) Write(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}
