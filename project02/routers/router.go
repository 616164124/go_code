package routers

import (
	"project02/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index2",&controllers.IndexController{})
	beego.Router("/index3",&controllers.Index3Controller{})
}
