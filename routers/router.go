package routers

import (
	"jms/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
//	beego.Router("/main",&controllers.MainController{})
}
