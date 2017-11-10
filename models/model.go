package models

import (
	"github.com/astaxie/beego/orm"
    "time"
)

type Temperatures struct {
	Id int
	Time time.Time
	CpuTemperature float64
    GpuTemperature float64
}

func init() {
	orm.RegisterModel(new(Temperatures))
}
