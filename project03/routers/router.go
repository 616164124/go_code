package routers
import (
	"project03/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api", &controllers.ApiController{})
	//动态路由
    beego.Router("/api/:id", &controllers.ApiController{})
	//
    beego.Router("/api/cms_:id([0-9]+).html", &controllers.CMSController{})

}
