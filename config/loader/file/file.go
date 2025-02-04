package file

import (
	"encoding/json"
	"github.com/mhthrh/common-lib/config/loader"
	customModelError "github.com/mhthrh/common-lib/errors"
	customError "github.com/mhthrh/common-lib/errors/config"
	textFile "github.com/mhthrh/common-lib/pkg/util/file/text"
	"strings"
)

type FileConfig struct {
	Path string
	name string
}

func New(val ...interface{}) (loader.IConfig, *customModelError.XError) {

	if len(val) < 2 {
		return nil, customError.FileParameter(nil)
	}
	for _, v := range val {
		if strings.Trim(v.(string), " ") == "" {
			return nil, customError.FileParameter(nil)
		}
	}
	if v1, ok := val[0].(string); ok {
		if v2, ok := val[1].(string); ok {
			return &FileConfig{
				Path: v1,
				name: v2,
			}, nil
		}
	}
	return nil, customError.FileParameter(nil)
}

func (f FileConfig) Initialize() (*loader.Config, *customModelError.XError) {
	var config loader.Config
	textFile.FileName = f.name
	textFile.FilePath = f.Path

	b := textFile.New(nil)
	if err := b.Read(); err != nil {
		return nil, customError.FileInitializerError(customModelError.RunTimeError(err))
	}
	err := json.Unmarshal(b.Data, &config)
	if err != nil {
		return nil, customError.FileInitializerError(customModelError.RunTimeError(err))
	}
	return &config, nil
}

func (f FileConfig) PrintConfig() *customModelError.XError {
	//TODO implement me
	panic("implement me")
}
