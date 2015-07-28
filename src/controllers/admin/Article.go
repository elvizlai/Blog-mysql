package admin
import (
	"models/article"
	"models/user"
	"strconv"
	"strings"
	"enum"
	"github.com/astaxie/beego/utils/pagination"
)

type Article struct {
	baseController
}

const limit = 20

//创建新文章
func (this *Article) NewArticle() {
	this.TplNames = "new_article.html"
	this.Data["categories"]=article.GetCategories(true)
}

//新建分类
func (this *Article) AddCategory() {
	categoryName := this.GetString("categoryName")
	id, err := article.AddCategory(categoryName, user.User{Id:this.userId})
	if err!=nil {
		this.Failed(enum.CategoryAlreadyExist)
	}else {
		result := map[string]interface{}{
			"Id":id,
			"Name":categoryName,
		}
		this.Succeed(enum.OK, result)
	}
	this.ServeJson()
}

//修改分类
func (this *Article) ModifyCategory() {
	categoryIdStr := this.GetString("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)
	categoryName := this.GetString("categoryName")

	id, err := article.ModifyCategory(int64(categoryId), categoryName, this.userId)
	if err!=nil {
		this.Failed(enum.CategoryAlreadyExist)
	}else {
		result := map[string]interface{}{
			"Id":id,
			"Name":categoryName,
		}
		this.Succeed(enum.OK, result)
	}
	this.ServeJson()
}

//新文章提交 ?isDraft=true 则表示当前提交为草稿
func (this *Article) PostArticle() {
	isDraft, _ := strconv.ParseBool(this.Input().Get("isDraft"))
	categoryId, _ := strconv.Atoi(this.GetString("categoryId"))
	topicTitle := this.GetString("articleTitle")
	content := this.GetString("content")
	articleTags := this.GetString("articleTags")

	//所有中文分号替换为英文分号
	articleTags = strings.Replace(articleTags, "；", ";", -1)

	article.AddArticle(topicTitle, content, int64(categoryId), articleTags, this.userId, isDraft)

//	for i:=0;i<1000;i++{
//		article.AddArticle(common.RandString(10), common.RandString(500), int64(categoryId), articleTags, this.userId, isDraft)
//	}

	if isDraft {
		this.Redirect("/draftList", 302)
	}else {
		this.Redirect("/articleList", 302)
	}
}

//修改文章 文章id
func (this *Article) ModifyArticle() {

	if this.Ctx.Request.Method=="GET" {
		this.TplNames = "modify_article.html"

		articleIdStr := this.Input().Get("id")

		if articleIdStr=="" {
			this.Redirect("articleList", 302)
			return
		}

		articleRes, err := article.GetArticleById(articleIdStr)
		if err!=nil {
			this.Redirect("articleList", 302)
			return
		}

		this.Data["articleId"]=articleIdStr
		this.Data["categories"]=article.GetCategories(true)
		this.Data["Title"]=articleRes.Title
		this.Data["Tags"] = articleRes.Tags
		this.Data["Content"]=articleRes.Content
		this.Data["Id"] = articleRes.Category.Id
	}else {
		articleId := this.GetString("articleId")
		isDraft, _ := strconv.ParseBool(this.Input().Get("isDraft"))
		topicTitle := this.GetString("articleTitle")
		content := this.GetString("content")
		categoryId, _ := strconv.Atoi(this.GetString("categoryId"))
		articleTags := this.GetString("articleTags")
		articleTags=strings.Replace(articleTags, "；", ";", -1)

		article.ModifyArticleById(articleId, topicTitle, content, int64(categoryId), articleTags, int64(this.userId), isDraft)
		if isDraft {
			this.Redirect("/draftList", 302)
		}else {
			this.Redirect("/articleList", 302)
		}
	}
}

//文章列表
func (this *Article) ArticleList() {
	this.TplNames = "article_list.html"

	currentPage := 1
	//获取上一页的页码
	if page, err := strconv.Atoi(this.Input().Get("p")); err==nil {
		currentPage = page
	}

	this.Data["isArticle"] = true
	var count int64
	this.Data["topics"], count = article.GetAllArticles(limit, currentPage)
	pagination.SetPaginator(this.Ctx, limit, count)
}

//草稿列表
func (this *Article) DraftList() {
	this.TplNames = "draft_list.html"
	this.Data["isDraft"] = true

	currentPage := 1
	//获取上一页的页码
	if page, err := strconv.Atoi(this.Input().Get("p")); err==nil {
		currentPage = page
	}

	var count int64
	this.Data["topics"], count = article.GetAllDrafts(limit, currentPage)
	pagination.SetPaginator(this.Ctx, limit, count)
}

//回收站列表
func (this *Article) TrashList() {
	this.TplNames = "trash_list.html"
	this.Data["isTrash"] = true
}




//删除文章 文章id
func (this *Article) DelArticle() {

}

func (this *Article) Succeed(e enum.Result, result interface{}) {
	this.Data["json"]=map[string]interface{}{
		"errcode":e,
		"msg":e.String(),
		"result":result,
	}
}

func (this *Article) Failed(e enum.Result) {
	this.Data["json"]=map[string]interface{}{
		"errcode":e,
		"msg":e.String(),
	}
}