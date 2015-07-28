package article
import (
	"time"
	"models/user"
	"github.com/astaxie/beego/orm"
)


//分类
type Category struct {
	Id       int64
	Name     string `orm:"unique"`                               //分类名
	Created  time.Time `orm:"auto_now_add;type(datetime);index"` //创建时间
	Updated  time.Time `orm:"null"`                              //更新时间
	Count    int `orm:"-"`
	User     *user.User `orm:"rel(fk)"`
	LastUser *user.User `orm:"null;rel(fk)"`                     //记录最后一个操作该分类的用户信息
}

//文章
type Article struct {
	Id       string `orm:"pk"`
	Category *Category `orm:"rel(fk)"`                           //分类
	Title    string                                              //标题
	Tags     string
	Abstract string `orm:"size(6000)"`                           //摘要
	Content  string `orm:"size(10000)"`                          //内容
	IsDraft  bool                                                //是否为草稿
	IsDel    bool                                                //是否已删除
	Created  time.Time `orm:"auto_now_add;type(datetime);index"` //创建时间
	Updated  time.Time `orm:"index;null"`                        //更新时间
	Views    int64 `orm:"index"`                                 //浏览数量
	User     *user.User `orm:"rel(fk)"`                          //作者
	LastUser *user.User `orm:"null;rel(fk)"`                     //记录最后一个操作该分类的用户信息
}

//回复 todo
type Reply struct {
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func init() {
	orm.RegisterModel(new(Category), new(Article))
}