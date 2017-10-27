package main

import (
	_ "piTemperature/routers"
	_ "piTemperature/tasks"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)


func main() {
	// start task regitstered in /tasks/task.go
	toolbox.StartTask()
	defer toolbox.StopTask()
	beego.Run()
}

