package article
import (
	"github.com/astaxie/beego/orm"
	"models/user"
	"common"
	"time"
	"strconv"
)

//添加分类
func AddCategory(name string, user user.User) (id int64, err error) {
	o := orm.NewOrm()
	category := &Category{Name:name, User:&user}
	_, id, err = o.ReadOrCreate(category, "Name")
	return
}

//修改分类
func ModifyCategory(categoryId int64, categoryName string, modifyUserId int64) (id int64, err error) {
	o := orm.NewOrm()

	id, err=o.QueryTable("category").Filter("Id", categoryId).Update(orm.Params{"Name":categoryName, "Updated":time.Now(), "last_user_id":modifyUserId})
	return
}

//获取所有分类,isAll==true数量最大的在上面
func GetCategories(isAll bool) []*Category {
	var categories []*Category
	o := orm.NewOrm()
	var maps []orm.Params
	if isAll {
		o.Raw("SELECT category.id,category.name,count(*) as count FROM category LEFT JOIN article ON category.id=article.category_id GROUP BY id ORDER BY count DESC;").Values(&maps)
	}else {
		o.Raw("SELECT category.id,category.name,count(*) as count FROM category INNER JOIN article ON category.id=article.category_id WHERE article.is_draft=FALSE AND article.is_del=FALSE GROUP BY category_id ORDER BY count DESC;").Values(&maps)
	}

	for _, v := range maps {
		var category Category
		category.Id, _=strconv.ParseInt(v["id"].(string), 10, 64)
		category.Name=v["name"].(string)
		category.Count, _=strconv.Atoi(v["count"].(string))
		categories=append(categories, &category)
	}

	return categories
}

// 添加文章
func AddArticle(title string, content string, categoryId int64, tags string, userId int64, isDraft bool) {
	o := orm.NewOrm()

	abstract := common.SubHtml(content, 100)

	topic := &Article{Title:title, Abstract:abstract, Content:content, Category:&Category{Id:categoryId}, Tags:tags, IsDraft:isDraft, User:&user.User{Id:userId}}
	topic.Id = common.CreateGUID()

	o.Insert(topic)

	return
}

//获取所有文章
func GetAllArticles(limit, page int) ([]*Article, int64) {
	var articles []*Article
	o := orm.NewOrm()
	setter := o.QueryTable("article").Filter("isDel", false).Filter("is_draft", false).OrderBy("-created").RelatedSel()
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&articles)

	return articles, count
}

//获取所有草稿
func GetAllDrafts(limit, page int) ([]*Article, int64) {
	var articles []*Article
	o := orm.NewOrm()
	setter := o.QueryTable("article").Filter("isDel", false).Filter("is_draft", true).OrderBy("-created").RelatedSel()
	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&articles)
	return articles, count
}

//根据id获取文章 todo 点击量模型未构建 返回改文章对应的标签
func GetArticleById(id interface{}) (Article, error) {
	article := Article{}
	o := orm.NewOrm()
	err := o.QueryTable("article").Filter("Id", id).RelatedSel().One(&article)

	if err==nil {
		o.QueryTable("article").Filter("Id", id).Update(orm.Params{
			"Views": orm.ColValue(orm.Col_Add, 1),
		})
		//查询该文章对应的标签
		//		var tags []*Tag
		//		_, err = o.QueryTable("tag").Filter("Articles__Article__Id", id).All(&tags)
		//		article.Tags = tags
	}

	return article, err
}

//修改文章
func ModifyArticleById(articleId, title, content string, categoryId int64, tags string, modifyUserId int64, isDraft bool) (id int64, err error) {
	o := orm.NewOrm()
	abstract := common.SubHtml(content, 100)

	//是否是草稿,如果为草稿，则需要更新创建时间
	var maps []orm.Params
	o.Raw("SELECT is_draft draft FROM article WHERE id=?;", articleId).Values(&maps)
	now := time.Now()
	params := orm.Params{}
	params["title"]=title
	params["abstract"]=abstract
	params["content"]=content
	params["category_id"]=categoryId
	params["is_draft"]=isDraft
	params["tags"]=tags
	params["updated"]=now
	params["last_user_id"]=modifyUserId


	if maps[0]["draft"].(string)=="1" {
		params["created"]=now
	}

	id, err=o.QueryTable("article").Filter("Id", articleId).Update(params)

	return
}

//删除文章
func DelArticleById(articleId string) {

}

//readorcreat
//func AddTags(tags []string) (tagList []*Tag) {
//	o := orm.NewOrm()
//
//	for i := 0; i<len(tags); i++ {
//		tagList = append(tagList, &Tag{Name:tags[i]})
//		o.ReadOrCreate(tagList[i], "Name")
//	}
//
//	return tagList
//}

//根据categoryId获取文章
func GetArticlesByCategoryId(categoryId interface{}, limit, page int64) ([]*Article, int64) {
	o := orm.NewOrm()

	var articles []*Article

	setter := o.QueryTable("article").Filter("isDel", false).Filter("is_draft", false).Filter("category_id", categoryId).OrderBy("-created").RelatedSel()

	count, _ := setter.Count()
	setter.Limit(limit, (page-1)*limit).All(&articles)
	return articles, count
}