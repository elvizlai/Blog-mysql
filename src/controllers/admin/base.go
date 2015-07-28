package admin

import (
	"github.com/astaxie/beego"
	"models/user"
)

type baseController struct {
	beego.Controller
	token  string
	userId int64
}

func (this *baseController) Prepare() {
	this.Layout = "layout.html"

	//todo 精简
	if token := this.GetSession(this.Ctx.GetCookie("token")); token!=nil {
		if uid, ok := user.VerifyToken(token.(string)); ok {
			this.token = token.(string)
			this.userId = uid
			this.Data["isLogin"] = true
		}else {
			this.Redirect("/login", 302)
		}
	}else {
		this.Redirect("/login", 302)
	}
}