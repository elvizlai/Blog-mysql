package user
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/bitly/go-simplejson"
	"fmt"
	"models/user"
	"enum"
	"common"
)

type User struct {
	beego.Controller
}

//注册 todo 错误处理的优雅封装
func (this *User) Register() {
	reqBody := this.Ctx.Input.RequestBody
	reqJson, err := simplejson.NewJson(reqBody)

	//非法请求
	if err!=nil {
		this.Abort("400")
	}

	email := reqJson.Get("email").MustString()
	nickname := reqJson.Get("nickname").MustString()
	password := reqJson.Get("password").MustString()


	valid := validation.Validation{}
	valid.Email(email, "email")//邮箱
	valid.MinSize(nickname, 5, "nickname")//昵称至少5位
	valid.MinSize(password, 6, "password")//密码至少6位

	//非法请求
	if valid.HasErrors() {
		this.Abort("400")
	}

	enumResult := user.AddUser(email,nickname,password)
	fmt.Println(enumResult)

	this.Data["json"]=map[string]interface{}{"code":enumResult, "msg":enumResult.String()}
	this.ServeJson()
}

//登录
func (this *User) Login() {
	fmt.Println(this)
	reqBody := this.Ctx.Input.RequestBody
	fmt.Println(string(reqBody))
	reqJson, err := simplejson.NewJson(reqBody)


	//非法请求
	if err!=nil {
		this.Abort("400")
	}
	fmt.Println(reqJson)

	email := reqJson.Get("email").MustString()
	password := reqJson.Get("password").MustString()

	valid := validation.Validation{}
	valid.Email(email,"email")
	valid.MinSize(password, 6, "password")//密码至少6位

	//非法请求
	if valid.HasErrors() {
		this.Abort("400")
	}
	fmt.Println(reqJson)

	if currentUser := user.FindUser(email); currentUser==nil {
		this.Data["json"]=map[string]interface{}{"code":enum.UserNotExist, "msg":enum.UserNotExist.String()}
	}else {
		if currentUser.Password!=common.Md5(password+currentUser.Salt){
			this.Data["json"]=map[string]interface{}{"code":enum.PasswordError, "msg":enum.PasswordError.String()}
		}else{
			this.Data["json"]=map[string]interface{}{"code":enum.OK, "msg":enum.OK.String()}
			//讲token写入cookie
			token:=user.UpdateCookieToken(*currentUser)
			this.SetSession(this.Ctx.GetCookie("token"),token)
		}
	}
	this.ServeJson()
}