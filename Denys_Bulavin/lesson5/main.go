package main

import (
	_ "lesson5/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type GetForm struct {
	Title   string `form:"art-title"`
	Article string `form:"art-article"`
	Tags    string `form:"art-tags"`
}

func main() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123@/lesson5")

	beego.Run()
}
