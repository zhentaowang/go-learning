package main

import (
	_ "data-process/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"code.aliyun.com/wyunshare/thrift-server"
	"data-process/routers"
)

func init() {
	//database connection
	mysql_user := beego.AppConfig.String("mysql_user")
	mysql_pass := beego.AppConfig.String("mysql_pass")
	mysql_urls := beego.AppConfig.String("mysql_urls")
	mysql_port := beego.AppConfig.String("mysql_port")
	mysql_name := beego.AppConfig.String("mysql_name")
	orm.RegisterDataBase("default", "mysql", mysql_user + ":" + mysql_pass + "@tcp("+ mysql_urls + ":" + mysql_port + ")/" + mysql_name)
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//thrift server address
	thrift_server_address := beego.AppConfig.String("thrift_server_address")
	thrift_server_port := beego.AppConfig.String("thrift_server_port")

	thriftserver.StartSingleServer(thrift_server_address, thrift_server_port, "businessService", routers.GetHandler()) // thrift server
	//beego.Run() // swagger api
}

