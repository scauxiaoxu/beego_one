package main

import (
	_ "apiproject/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库配置
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	runmode := beego.AppConfig.String("runmode")
	if runmode == "dev" {
		orm.Debug = true
	}

	beego.Run()
}
