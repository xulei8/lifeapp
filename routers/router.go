package routers

import (
	"github.com/astaxie/beego"
	"github.com/xulei8/lifeapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}) //test

	beego.Router("/appser/", &controllers.AppSer{})

}
