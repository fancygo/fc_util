package fc_util

import (
	"os"
	"path"
)

var (
	workDir string
	logDir  string
	resDir  string
)

func init() {
	workDir, _ = os.Getwd()
	logDir = path.Join(workDir, "log")
	resDir = path.Join(workDir, "res")
}

func GetWorkDir() string {
	return workDir
}

func GetLogDir() string {
	return logDir
}

func GetResDir() string {
	return resDir
}
