package initial
import (
	"common"
	"github.com/astaxie/beego"
)

const ImagePaht = "files/image/"
const AttachmentPaht = "files/attchment/"
const ScrawlPath = "files/scrawl/"

var files = []string{ImagePaht, AttachmentPaht, ScrawlPath}

func initFileSys() {
	for i := 0; i<len(files); i++ {
		if common.FileExist(files[i]) {
			beego.Debug("file", files[i], "already exists, skip")
		}else {
			common.CreateFile(files[i])
			beego.Debug("file", files[i], "created")
		}
	}
	beego.SetStaticPath("/files", "files")//静态文件
}