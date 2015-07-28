package controllers
import "github.com/astaxie/beego"

type Test struct {
    beego.Controller
}

func (this *Test) Get(){
    this.TplNames = "test.html"
}