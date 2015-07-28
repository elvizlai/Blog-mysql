package controllers
import (
	"github.com/astaxie/beego"
	"common"
	"fmt"
	"time"
//    "os"
	"encoding/base64"
	"io/ioutil"
	"initial"
	"io"
	"mime/multipart"
)


//todo 文件大小限制
type Upload struct {
	beego.Controller
}

func (this *Upload) Get() {
	conf := common.ReadFile("conf/ueditor.json")
	this.Ctx.WriteString(conf)
}

func (this *Upload) Post() {
	action := this.Input().Get("action")

	f, h, err := this.GetFile("upfile")
	if err==nil {
		defer f.Close()
	}

	switch action{
		//图片上传
		case "uploadimage":
		if fileExceed(f, 2048000) {
			this.uploadFailed("file size exceed")
		}else {
			this.SaveToFile("upfile", initial.ImagePaht + h.Filename)
			url, _ := this.saveToOss(initial.ImagePaht, h.Filename)
			this.uploadSucceed(url, h.Filename)//this.uploadSucceed(initial.ImagePaht+h.Filename, "")
		}

		//附件上传
		case "uploadfile":
		if fileExceed(f, 10240000) {
			this.uploadFailed("file size exceed")
		}else {
			this.SaveToFile("upfile", initial.AttachmentPaht+h.Filename)
			url, _ := this.saveToOss(initial.AttachmentPaht, h.Filename)
			this.uploadSucceed(url, h.Filename)
		}

		//涂鸦上传
		case "uploadscrawl":
		upfile := this.Ctx.Request.Form.Get("upfile")
		data, _ := base64.StdEncoding.DecodeString(upfile)
		scrawlName := fmt.Sprint(time.Now().UnixNano())+".png"
		ioutil.WriteFile(initial.ScrawlPath +scrawlName, data, 0666)//buffer输出到png文件中(不做处理，直接写到文件)
		url, _ := this.saveToOss(initial.ScrawlPath, scrawlName)
		this.uploadSucceed(url, scrawlName)

		default:
		this.uploadFailed("unknown")
	}
	this.ServeJson()
}


func fileExceed(file multipart.File, limitSize int64) bool {
	var isExceed bool
	read_buf := make([]byte, 1024)
	var pos int64 = 0
	for {
		n, err := file.ReadAt(read_buf, pos)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		pos = pos +(int64)(n)
		if pos > limitSize {
			isExceed = true
			break
		}
	}
	return isExceed
}

func (this *Upload) uploadFailed(errorMsg string) {
	this.Data["json"] = map[string]interface{}{
		"state": errorMsg,
	}
}

func (this *Upload) uploadSucceed(url, filename string) {
	this.Data["json"] = map[string]interface{}{
		"state":    "SUCCESS",
		"url":      url, //保存后的文件路径
		"title":    filename, //文件描述，对图片来说在前端会添加到title属性上
		"original": filename, //原始文件名
	}
}

func (this *Upload) saveToOss(filePath, fileName string) (url string, err error) {
	t := time.Now()
	ossFilename := fmt.Sprintf("%s%d/%d/%d/%s", filePath, t.Year(), t.Month(), t.Day(), fileName)
	err = common.OssStore(ossFilename, filePath+fileName)
	url = common.OssGetURL(ossFilename)
	return
}

