package loader

import (
	"encoding/json"
	textFile "github.com/mhthrh/common-lib/pkg/util/file/text"
)

type FileConfig struct {
	path string
	name string
}

func (f FileConfig) Initialize() (*Config, error) {
	var config Config
	textFile.FileName = f.name
	textFile.FilePath = f.path

	b := textFile.New(nil)
	if err := b.Read(); err != nil {
		return nil, err
	}
	err := json.Unmarshal(b.Data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (f FileConfig) PrintConfig() error {
	//TODO implement me
	panic("implement me")
}
