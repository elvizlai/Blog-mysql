package initial
import "github.com/astaxie/beego"

var mysql, pq, sqlite string

func initConfig() {
	mysql = beego.AppConfig.String("mysql")
	pq = beego.AppConfig.String("pq")
	sqlite = beego.AppConfig.String("sqlite")
}