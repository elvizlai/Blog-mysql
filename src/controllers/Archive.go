//归档
package controllers

type Archive struct {
	base
}

func (this *Archive) Get() {
	this.TplNames="archives.html"
}