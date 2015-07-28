package attachment
import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Attachment struct {
	Id      int64
	Path    string
	Url     string
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Attachment))
}