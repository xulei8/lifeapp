package main

import (
	"github.com/astaxie/beego"
	m "github.com/xulei8/lifeapp/models"
	_ "github.com/xulei8/lifeapp/routers"
	"os"
)

func main() {
	initArgs()
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "--newdb" || v == "new" || v == "newdb" {
			m.Syncdb()
			os.Exit(0)

		}
	}
}
