package controllers
import (
	"models/user"
	"models/article"
	"github.com/astaxie/beego/utils/pagination"
	"strconv"
)

const pageLimit = 10

type Home struct {
	base
}

func (this *Home) Index() {
	this.TplNames="home.html"
	this.Data["Title"]="首页"

	currentPage := 1
	//获取上一页的页码
	if page, err := strconv.Atoi(this.Input().Get("p")); err==nil {
		currentPage = page
	}

	articleList, count := article.GetAllArticles(pageLimit, currentPage)

	pagination.SetPaginator(this.Ctx, pageLimit, count)

	this.Data["topicList"]=articleList

	this.Data["categoryList"]=article.GetCategories(false)
}

//
//根据分类列出文章 todo 错误处理
func (this *Home)Category() {
	this.TplNames="home.html"

	var id int
	var err error
	if id, err = strconv.Atoi(this.Input().Get("id")); err==nil {

	}

	currentPage := 1
	//获取上一页的页码
	if page, err := strconv.Atoi(this.Input().Get("p")); err==nil {
		currentPage = page
	}

	var count int64
	this.Data["topicList"], count = article.GetArticlesByCategoryId(int64(id), pageLimit, int64(currentPage))//todo

	this.Data["categoryList"]=article.GetCategories(false)

	pagination.SetPaginator(this.Ctx, pageLimit, count)
}

//todo 上一篇，下一篇
func (this *Home) Article() {
	this.TplNames = "article.html"
	topicId := this.Input().Get("id")
	if topicId=="" {
		this.Redirect("/", 302)
		return
	}

	article, err := article.GetArticleById(topicId)
	if err!=nil {
		//找不到该文章
		this.Abort("404")
	}

	this.Data["article"] = article
}

//登录
func (this *Home) Login() {
	this.TplNames = "login.html"
}

//注册
func (this *Home) Register() {
	this.TplNames = "register.html"
}

func (this *Home) VerifyEmail() {
	code := this.Input().Get("code")
	if code=="" {
		this.Ctx.WriteString("failed")
	}else {
		err := user.VerifyEmail(code)
		if err!=nil {
			this.Ctx.WriteString(err.Error())
		}else {
			this.Ctx.WriteString("chenggong")
		}
	}
}

//登出
func (this *Home)Logout() {
	this.DelSession(this.Ctx.GetCookie("token"))
	this.Redirect("/", 302)
}