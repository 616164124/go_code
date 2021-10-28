package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"project02/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index2", &controllers.IndexController{})
	beego.Router("/index3", &controllers.Index3Controller{})
	//
	//user
	beego.Router("/user", &controllers.UserController{}, "get:Get")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAdd")
	//返回json
	beego.Router("/user/getuser", &controllers.UserController{}, "get:Getuser")
	beego.Router("/user/getEdit", &controllers.UserController{}, "get:GetEdit")
	beego.Router("/user/edit", &controllers.UserController{}, "post:DoEditUser")

	//get put del post
	beego.Router("/goods/get", &controllers.GoodsController{}, "get:GetGoods")
	beego.Router("/goods/put", &controllers.GoodsController{}, "put:PutGoods")
	beego.Router("/goods/post", &controllers.GoodsController{}, "post:PostGoods")
	beego.Router("/goods/del", &controllers.GoodsController{}, "delete:DelGoods")
	//Xml
	beego.Router("/goods/xml", &controllers.GoodsController{}, "post:Xml")
	//crud

}
