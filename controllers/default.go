package controllers

import (
	"github.com/astaxie/beego"
    "piTemperature/tasks"
)

type MainController struct {
	beego.Controller
}

type SendJsonController struct{
	beego.Controller
}

func (c *MainController) Get() {
    c.Layout = "base.html"
	c.TplName = "main.html"
    c.Data["temperature"], _=tasks.ReadTemperature()
}

func (c *SendJsonController) Get(){
	hash :=Hash{"1235", 64.113}
	c.Data["json"]=hash
	c.ServeJSON()
}


type Hash struct {
	Time string
	Temperature float64
}
