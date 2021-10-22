package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type CMSController struct {
	beego.Controller
}
//路由伪静态
//http://localhost:7073/api/cms_12321314.html
func (c *CMSController) Get() {
	//获取值(获取:id后面所有的值包括.html)
	id :=c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("cms接口-----"+id)

}
