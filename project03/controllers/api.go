package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ApiController struct {
	beego.Controller
}
//动态路由
//http://localhost:7073/api/123
func (c *ApiController) Get() {
	//获取值
	id :=c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口-----"+id)

}
