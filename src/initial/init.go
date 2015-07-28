package initial

import "github.com/astaxie/beego"



//todo 文件初始化
func init() {
	beego.Debug("init...")
	initConfig()//读取配置文件
	initFileSys()//初始化文件系统
	initSql()//连接数据库
	beego.Debug("init done...")
}