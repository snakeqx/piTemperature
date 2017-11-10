package tasks

import(
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"piTemperature/models"
	"time"
	"os"
    "os/exec"
	"strconv"
	"bufio"
)

func init() {
	// define tasks
	taskStoreTemperature := toolbox.NewTask("taskStoreTemperature", beego.AppConfig.String("taskfrequency"), StoreTemperature)

	// add tasks into global task list
	toolbox.AddTask("taskStoreTemperature", taskStoreTemperature)

	// orm handling of database
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//parameter description for RegisterDataBase
	//para1: aliasname (database name used in this beego project, not the real name in the mysql!!!
	//para2: database type
	//para3: connection by below format:
	//       <user>:<password>@/<databasename in mysql>?<charset>
	//                         ^ be aware of the slash ("/"), this is mandantory
	orm.RegisterDataBase("default", "mysql", "defaultuser:default@/various?charset=utf8&loc=Asia%2FShanghai")
	// create table
	orm.RunSyncdb("default", true, true)
}


func StoreTemperature () error {
	o := orm.NewOrm()
	o.Using("default")

	temp := new(models.Temperatures)
	timeNow:=time.Now()
	temp.Time=timeNow//.Format("2006-01-02 15:04:05")
    var err error
	temp.CpuTemperature, err= ReadCpuTemperature()
    if err !=nil{
        return err
    }
    temp.GpuTemperature, err = ReadGpuTemperature()
    if err != nil{
        return err
    }

	o.Insert(temp)

	return nil
}

func ReadCpuTemperature() (float64, error) {
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

func ReadGpuTemperature() (float64, error) {
    cmd := exec.Command("vcgencmd","measure_temp")
    output, err := cmd.Output()
    if err != nil{
        return -1, err
    }
    str := string(output)
    str = str[5:len(str)-3]
    f,err :=strconv.ParseFloat(str,64)
    if err!=nil{
        return -1, err
    }
    return f, err

}
