package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me" //绑定html中的数据
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index2.html"
	
}

