package main

import (
	_ "project02/routers"
	beego "github.com/beego/beego/v2/server/web"	
)

func main() {
	beego.Run()
	
}

