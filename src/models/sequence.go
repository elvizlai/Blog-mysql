package models
import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
)

//获取序列号
func CurrVal(name string) int64 {
	o := orm.NewOrm()
	sql := fmt.Sprintf("SELECT currval('%v') uid", name)
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	uidStr := maps[0]["uid"].(string)
	uid64, _ := strconv.ParseInt(uidStr, 10, 64)
	return uid64
}

//序列号+1
func NextVal(name string) {
	o := orm.NewOrm()
	sql := fmt.Sprintf("SELECT nextval('%v') uid", name)
	o.Raw(sql).Exec()
}