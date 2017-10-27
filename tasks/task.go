package tasks

import(
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"piTemperature/models"
	"time"
	"os"
	"strconv"
	"bufio"
)

func init() {
	// define tasks
	taskGetTemperature := toolbox.NewTask("taskGetTemperature", "*/5 * * * * *", GetTemperature)

	// add tasks into global task list
	toolbox.AddTask("taskGetTemperature", taskGetTemperature)

	// orm handling of database
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//parameter description for RegisterDataBase
	//para1: aliasname (database name used in this beego project, not the real name in the mysql!!!
	//para2: database type
	//para3: connection by below format:
	//       <user>:<password>@/<databasename in mysql>?<charset>
	//                         ^ be aware of the slash ("/"), this is mandantory
	orm.RegisterDataBase("default", "mysql", "defaultuser:default@/various?charset=utf8")
	// create table
	orm.RunSyncdb("default", true, true)
}


func GetTemperature () error {
	o := orm.NewOrm()
	o.Using("default")

	temp := new(models.Temperatures)
	timeNow:=time.Now()
	temp.Time=timeNow.Format("2006-01-02 15:04:05")
	temp.Temperature, _ =ReadTemperature()

	o.Insert(temp)
	return nil
}

func ReadTemperature() (float64, error) {
	fileString:=`/sys/class/thermal/thermal_zone0/temp`
	file, err := os.Open(fileString)
	defer file.Close()
	// prepare to read file
	var tempStr string
	switch err{
	case nil:
		read :=bufio.NewReader(file)
		line,  _:=read.ReadString('\n')
		//get rid of the '\n' EOF string
		//otherwise Atoi will have error
		line =line[:len(line)-1]
		tempStr=line
	default:
		beego.Critical(err)
		return -1, err
	}
	// convert string to number
	var tempInt int
	var tempFloat float64
	tempInt, err =strconv.Atoi(tempStr)
	switch err{
	case nil:
		tempFloat=float64(tempInt)/1000
	default:
		beego.Critical(err)
		return -1, err
	}
	return tempFloat, nil 
}

