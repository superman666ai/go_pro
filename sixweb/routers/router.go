package routers

import (
	"sixweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //beego.Router("/user", &controllers.UserController{})
}
