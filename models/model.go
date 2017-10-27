package models

import (
	"github.com/astaxie/beego/orm"
)

type Temperatures struct {
	Id int
	Time string
	Temperature float64
}

func init() {
	orm.RegisterModel(new(Temperatures))
}
