package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller

}
type Index3Controller struct {
	beego.Controller
	
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
	
}


func (c *Index3Controller) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Ctx.WriteString("你好！！！！")

	c.TplName = "index3.html"
	
}


func (c *Index3Controller) Get2() {
	c.TplName = "index3.html"
	
}


