package functions
import (
	"github.com/astaxie/beego"
	"html/template"
)

func init() {
	beego.AddFuncMap("set", set)
	beego.AddFuncMap("AssetsCss", AssetsCss)
	beego.AddFuncMap("AssetsJs", AssetsJs)
}


//add by ElvizLai
//Usage:当前模板上下文中设置一个变量
//{{set . "var" "Mes"}}
//{{.var}}
func set(renderArgs map[interface{}]interface{}, key string, value interface{}) template.JS {
	renderArgs[key] = value
	return template.JS("")
}

// returns script tag with src string.
func AssetsJs(src string) template.HTML {
	text := string(src)
	text = "<script src=\"" + src + "\"></script>"
	return template.HTML(text)
}

// returns stylesheet link tag with src string.
func AssetsCss(src string) template.HTML {
	text := string(src)
	text = "<link href=\"" + src + "\" rel=\"stylesheet\" />"
	return template.HTML(text)
}
