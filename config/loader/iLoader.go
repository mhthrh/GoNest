package loader

import (
	"github.com/mhthrh/common-lib/config/model"
	customModelError "github.com/mhthrh/common-lib/model/error"
)

type Config struct {
	MetaData model.MetaData `json:"metaData"`
	Secret   model.Secret   `json:"secret"`
	DataBase model.DB       `json:"db"`
}

type IConfig interface {
	Initialize() (*Config, *customModelError.XError)
	PrintConfig() *customModelError.XError
}
