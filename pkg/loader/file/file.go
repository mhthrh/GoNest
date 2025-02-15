package file

import (
	"encoding/json"
	customModelError "github.com/mhthrh/GoNest/model/error"
	"github.com/mhthrh/GoNest/model/loader"
	textFile "github.com/mhthrh/GoNest/pkg/util/file/text"
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
		return nil, FileInitializerError(customModelError.RunTimeError(err))
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, FileInitializerError(customModelError.RunTimeError(err))
	}
	return &config, nil
}

func (f FileConfig) PrintConfig() *customModelError.XError {
	//TODO implement me
	panic("implement me")
}
