package main

import (
	_ "hello_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hello_api/kafka"
)

func init(){
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	logLevel, err:= beego.AppConfig.Int("logLevel")
	if nil != err{
		logLevel = beego.LevelDebug
	}
	beego.SetLevel(logLevel)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		beego.AppConfig.String("DBUserName"), beego.AppConfig.String("DBPassword"),
		beego.AppConfig.String("DBAddr"), beego.AppConfig.String("DBName"))
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default","mysql",dataSource,maxIdle, maxConn)
	orm.RunSyncdb("default",false,true)
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//test()
	//models.AddApp("kevin","bbbbbbbbbb",1)
	//models.DeleteApp("kevin")
	//models.UpdateApp("kevin","2222222222")
	beego.Run()
	kafka.Start()
}
