//关于
package controllers

type About struct {
    base
}

func (this *About) Get(){
    this.TplNames="me.html"
}