package routers
import (
    "github.com/astaxie/beego"
    "controllers/user"
)

func init(){
    //注册
    beego.Router("/register",&user.User{},"post:Register")
    //登陆
    beego.Router("/login",&user.User{},"post:Login")
}