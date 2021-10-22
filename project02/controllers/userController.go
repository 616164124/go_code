package controllers

import (
	"fmt"
	"strconv"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}
//User结构体  
// form 中的username 跟页面中的username要一致
type User struct{
	Username string `form:"username"`
	Password string `form:"Password"`
	Hobby []string `form:"hobby"`
	
}

func (c *UserController) Get() {
	c.TplName = "user.html"
}

//登录模式练习
func (c *UserController) DoAdd() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id 必须为数字类型")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")
	c.Ctx.WriteString("id " + strconv.Itoa(id) + "username" + username + "\t" + password)

}
//GetEdit
func (c *UserController) GetEdit(){

	c.TplName="edit.html"
}

//返回json数据
func (c *UserController) Getuser(){
	u :=User{
		Username :"张三",
		Password : "12334",
		Hobby:[]string{"1","123"},
	}
	//这两句可以返回json数据
	c.Data["json"]=u
	c.ServeJSON()
}
//将页面的数据（json）放在结构体里
func (c *UserController) DoEditUser(){
	u :=User{}
	err :=c.ParseForm(&u);
	if err!=nil{
		c.Ctx.WriteString("post提交失败")
	}
	//打印结构体内的数据
	fmt.Printf("%#v",u)
	c.Ctx.WriteString("post提交成功")
}