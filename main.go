package main

import (
	"fmt"

	_ "github.com/PBP-API-Framework-1120043/models"
	_ "github.com/PBP-API-Framework-1120043/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/joho/godotenv"
)

func main() {
	// if in develop mode
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// get database configuration from environment variables
	// dbUser := os.Getenv("DB_USER")
	// dbPwd := os.Getenv("DB_PWD")
	// dbName := os.Getenv("DB_NAME")
	dbUser := "root"
	dbPwd := ""
	dbName := "explorasi_framework_beego"
	fmt.Println(dbUser, ", ", dbPwd, ", ", dbName)
	dbString := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=utf8"

	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)

	// autosync
	// db alias
	name := "default"

	// drop table and re-create
	force := false

	// print log
	verbose := true

	// error
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
