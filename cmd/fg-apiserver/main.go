package main

import (
	"os"

	"github.com/TobyIcetea/fastgo/cmd/fg-apiserver/app"
	_ "go.uber.org/automaxprocs"
)

func main() {
	// 创建 Go 极速项目
	command := app.NewFastGoCommand()

	// 执行命令并处理错误
	if err := command.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序（例如 bash 脚本）根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
