package config

import (
	"encoding/json"
	"github.com/mhthrh/common-lib/config/model"
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
	MetaData model.MetaData `json:"metaData"`
	Secret   model.Secret   `json:"secret"`
	DataBase model.DB       `json:"db"`
}

type Vault struct {
	MetaData model.MetaData `json:"metaData"`
	Secret   model.Secret   `json:"secret"`
	DataBase model.DB       `json:"db"`
}

func (f *File) Initialize() error {
	var config File
	textFile.FileName = ""
	textFile.FilePath = ""
	b := textFile.New(nil)
	if err := b.Read(); err != nil {
		return err
	}
	err := json.Unmarshal(b.Data, &config)
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
