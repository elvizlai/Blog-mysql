package attachment
import (
	"github.com/astaxie/beego/orm"
	"time"
)

func AddAttachment(path, url string) {
	o := orm.NewOrm()
	attachment := Attachment{Path:path, Url:url, Created:time.Now()}
	o.Insert(&attachment)
}