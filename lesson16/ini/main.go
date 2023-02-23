package main

import (
	"fmt"

	ini "gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(), // MustInt("初期値")

		DbName: cfg.Section("db").Key("name").String(),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)           // int 8080
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)       // string webapp.sql
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver) // string sqlite3
}
