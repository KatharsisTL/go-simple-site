package routers

import (
	"go-simple-site/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/index", &controllers.MainController{})
}
