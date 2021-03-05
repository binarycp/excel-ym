package excel

import (
	"excel-ym/app"
	"excel-ym/assets"
	"excel-ym/utils"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gen2brain/beeep"
	"strings"
)

type Info struct {
	Path string
	F    *excelize.File
}

const (
	SHEET = "下载excel服务"
)

var (
	message = `统计报表生产成功
{%order}
(%device)
(%org)`
)

func init() {
	c := make(chan struct{})
	go func() {
		replacer := strings.NewReplacer("{%order}", Order.Path, "(%device)", device.Path, "(%org)", org.Path)
		utils.Err(beeep.Notify("执行成功", replacer.Replace(message), app.InfoImg))
		c <- struct{}{}
	}()
	<-c
}

func New(path string, template string) *Info {
	return &Info{
		Path: setPath(path),
		F:    getFileFS(template),
	}
}

func setPath(path string) string {
	return strings.ReplaceAll(utils.GetRoot(path), "{$date}", utils.Day())
}

func getFile(template string) *excelize.File {
	file, err := excelize.OpenFile(template)
	utils.Err(err)
	return file
}

func getFileFS(template string) *excelize.File {
	open, err := assets.FS.Open(template)
	if open != nil {
		defer open.Close()
	}
	utils.Err(err)
	openReader, err := excelize.OpenReader(open)
	utils.Err(err)
	return openReader
}
