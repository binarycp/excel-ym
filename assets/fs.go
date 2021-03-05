//Package 使用embed编译静态资源
package assets

import (
	"embed"
	"excel-ym/app"
	"github.com/binarycp/gutils/files"
	"os"
)

var (
	//用于读
	//go:embed template
	FS embed.FS

	//用于写入文件
	//go:embed resources/imagines
	wFS embed.FS
)

func init() {
	generateResource()
}

// 创建资源文件
func generateResource() {
	resources := [...]string{app.InfoImg, app.WarningImg}
	for _, v := range resources {
		bytes, _ := wFS.ReadFile(v)
		err := files.CreateDir(v)
		if err != nil {
			return
		}
		err = os.WriteFile(v, bytes, 0644)
		if err != nil {
			return
		}
	}
}
