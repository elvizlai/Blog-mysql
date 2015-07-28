package controllers
import (
	"github.com/astaxie/beego"
)

type base struct {
	beego.Controller
}

func (this *base) Prepare() {
	this.Layout = "layout.html"
}