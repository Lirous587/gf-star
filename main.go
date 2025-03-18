package main

import (
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"star/internal/cmd"
	_ "star/internal/packed"
)

func main() {
	var err error

	//全局设置i18n
	g.I18n().SetLanguage("zh-CN")

	err = connDb()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}

func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接数据库失败")
	}
	return nil
}
