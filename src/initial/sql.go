package initial

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
//    _ "github.com/lib/pq"
//    _ "github.com/mattn/go-sqlite3"
	_ "models/user"
	_ "models/article"
	_ "models/attachment"
	"common"
)

func initSql() {
	beego.Debug("init sql")

	Using := beego.AppConfig.String("USINGDB")

	switch Using{
	case "mysql":
		orm.RegisterDataBase("default", "mysql", mysql)
	case "pq":
		orm.RegisterDataBase("default", "postgres", pq)
	case "sqlite":
		if !common.FileExist(sqlite) {
			common.CreateFile(sqlite)
		}
		orm.RegisterDataBase("default", "sqlite3", sqlite)
	}
	orm.RunSyncdb("default", false, true)
	beego.Debug("init sql done")
}