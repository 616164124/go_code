package controllers

import (
	"encoding/xml"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type GoodsController struct {
	beego.Controller
}

type Goods struct {
	Titil   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

//func 方法名必须大写开头
func (c *GoodsController) GetGoods() { //get
	c.Ctx.WriteString("12313")
	p := Goods{}
	err := c.ParseForm(&p)
	if err!=nil{
		return 
	}
	fmt.Printf("%#v", p)
	c.Ctx.WriteString("get操作")

}

func (c *GoodsController) PostGoods() { //post
	p :=Goods{}
   err :=c.ParseForm(&p)
	if err!=nil {
		c.Ctx.WriteString("postgoods 失败")
		return 
	}
	if p.Titil!=""{
		c.Ctx.WriteString("postgoods失败 Titil 不能为空")
		return
	}
	fmt.Printf("%#v",p)
	c.Ctx.WriteString("post执行成功")


}

func (c *GoodsController) PutGoods() { //put

}

func (c *GoodsController) DelGoods() { //delete

}
//获取xml数据 xml中的名字 配置文件添加 copyrequestbody=true
/*
标签的名字一一对应（大小写一致）
xml的结构体
type Goods struct {
	Titil   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}
xml的提交的数据
<?xml version="1.0" encoding="UTF-8" ?>
<goods>
	<title type="string">语文</title>
    <content type="string">鲁迅文章</content>
</goods>
*/
func (c *GoodsController) Xml(){
	p:=Goods{}
	var err error
	if e:=xml.Unmarshal(c.Ctx.Input.RequestBody,&p);e!=nil {
		c.Data["json"]=err.Error()
		c.ServeJSON()
	}else {
		fmt.Printf("%#v",p)
		c.Data["json"]=p
		c.ServeJSON()
	}
}
