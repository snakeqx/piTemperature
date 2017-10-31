package controllers

import (
	"github.com/astaxie/beego"
    "piTemperature/tasks"
    "piTemperature/models"
    "github.com/astaxie/beego/orm"
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
    // define the return value is a list of struct temperatures
    var temperatures []*models.Temperatures
    // define queity
    o := orm.NewOrm()
    qs :=o.QueryTable("temperatures")
    qs.OrderBy("-id").Limit(20).All(&temperatures)
	c.Data["json"]=temperatures
	c.ServeJSON()
}


