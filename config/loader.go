package config

import (
	"encoding/json"
	"github.com/mhthrh/common-lib/config/db"
	textFile "github.com/mhthrh/common-lib/pkg/util/file/text"
)

const (
	path = ""
)

type Config interface {
	Initialize() error
	PrintConfig() error
}
type File struct {
	DataBase db.DB
}

type Vault struct {
	Database db.DB
}

func (f *File) Initialize() error {
	textFile.FileName = ""
	textFile.FilePath = ""
	b := textFile.New(nil)
	if err := b.Read(); err != nil {
		return err
	}
	err := json.Unmarshal(b.Data, &f.DataBase)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) PrintConfig() error {
	return nil
}

func (v *Vault) Initialize() error {
	return nil
}

func (v *Vault) PrintConfig() error {
	return nil
}
