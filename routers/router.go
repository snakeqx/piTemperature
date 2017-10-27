package routers

import (
	"piTemperature/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/getdata", &controllers.SendJsonController{})
}
