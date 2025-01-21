package directory

import (
	"os"
	"strings"
)

const (
	appMainDir = "common-lib"
)

var (
	appPath = ""
)

func GetWorkingDir() (path string, e error) {

	defer func() {
		if err := recover(); err != nil {
			path = ""
			e = err.(error)
		}
	}()
	appPath, _ = os.Getwd()
	if l := strings.Index(appPath, appMainDir); l > 0 {
		appPath = appPath[:l+len(appMainDir)+1]
	}
	return appPath, nil

}
