package utils

import (
	"kcers-survey/biz/dal/config"
	"kcers-survey/biz/pkg/consts"
	"os"
	"strconv"
	"time"
)

func ExportFilePath(v string) (string, string) {
	timePath := time.Now().Format(time.DateOnly) + "/"
	exportFilePath := consts.ExportFilePath + timePath
	if err := os.MkdirAll(exportFilePath, 0o777); err != nil {
		panic(err)
	}
	ing := strconv.FormatInt(time.Now().Unix(), 10)
	name := v + ing + ".xlsx"
	files := exportFilePath + name
	domain := config.GlobalServerConfig.Domain + "export/" + timePath + name
	return files, domain
}
