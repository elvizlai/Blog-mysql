package routers
import (
	"github.com/astaxie/beego"
	"controllers"
	"controllers/admin"
)

func init() {
	beego.Debug("init routers")

	//首页文章列表
	beego.Router("/", &controllers.Home{},"get:Index")

	//首页分类
	beego.Router("/category", &controllers.Home{},"get:Category")

	//文章详情
	beego.Router("/article", &controllers.Home{},"get:Article")

	//归档 todo
	beego.Router("/archives", &controllers.Archive{})

	//关于 todo
	beego.Router("/me", &controllers.About{})

	//邮箱验证
//	beego.Router("/verifyEmail",&controllers.Home{},"*:VerifyEmail")



	//登陆
	beego.Router("/login", &controllers.Home{}, "get:Login")
	//注册
	beego.Router("/register",&controllers.Home{},"get:Register")

	//找回密码 todo

	//退出
	beego.Router("/logout", &controllers.Home{}, "*:Logout")


	//-----------------管理员相关------------------------
	//获取文章列表
	beego.Router("/articleList", &admin.Article{}, "get:ArticleList")
	//获取草稿列表
	beego.Router("/draftList", &admin.Article{}, "get:DraftList")
	//获取回收站列表
	beego.Router("/trashList", &admin.Article{}, "get:TrashList")
	//增加文章
	beego.Router("/newArticle", &admin.Article{}, "get:NewArticle")
	//增加分类
	beego.Router("/addCategory", &admin.Article{}, "post:AddCategory")
	//修改分类
	beego.Router("/modifyCategory", &admin.Article{}, "post:ModifyCategory")
	//发布文章
	beego.Router("/postArticle", &admin.Article{}, "*:PostArticle")
	//修改文章
	beego.Router("/modifyArticle", &admin.Article{}, "*:ModifyArticle")

	beego.Debug("init routers done")
}