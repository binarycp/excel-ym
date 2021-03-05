package utils

import (
	"excel-ym/app"
	"github.com/gen2brain/beeep"
	"os"
)

func Err(vars ...interface{}) {
	err := TypeIsError(vars)
	if err != nil {
		_ = beeep.Alert("错误信息", err.Error(), app.WarningImg)
		os.Exit(1)
	}
}
