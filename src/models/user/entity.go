package user
import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id            int64  `orm:"pk"`           //使用sequence自增
	Nickname      string `orm:"unique"`       //昵称
	Email         string `orm:"unique;index"` //用户名
	Password      string
	Salt          string


	EmailVerified bool
	EmailToken    string
}

//登录成功后token存储
type UserToken struct {
	Id      int64
	Token   string `orm:"unique"`
	Updated time.Time
	User    *User  `orm:"rel(fk)"`
}

//密码重置
//type Verify struct {
//	Id         int64
//	ResetToken string
//	InsertTime time.Time
//	User       *User `orm:"rel(fk)"`
//}

func init() {
	orm.RegisterModel(new(User), new(UserToken))
}