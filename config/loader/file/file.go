package file

import (
	"encoding/json"
	"github.com/mhthrh/common-lib/config/loader"
	customModelError "github.com/mhthrh/common-lib/model/error"
	customError "github.com/mhthrh/common-lib/pkg/util/file"
	textFile "github.com/mhthrh/common-lib/pkg/util/file/text"
)

type FileConfig struct {
	Path string
	name string
}

func New(path, name string) loader.IConfig {
	return &FileConfig{
		Path: path,
		name: name,
	}
}

func (f FileConfig) Initialize() (*loader.Config, *customModelError.XError) {
	var config loader.Config
	text := textFile.New(f.Path, f.name)

	b, err := text.Read()
	if err != nil {
		return nil, customError.FileInitializerError(customModelError.RunTimeError(err))
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, customError.FileInitializerError(customModelError.RunTimeError(err))
	}
	return &config, nil
}

func (f FileConfig) PrintConfig() *customModelError.XError {
	//TODO implement me
	panic("implement me")
}
