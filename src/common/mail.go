package common
import (
	"net/smtp"
	"github.com/astaxie/beego"
)

type msgType string

const (
	Html msgType = "html"
	Text msgType = "text"
)

func SendMail(msgtype msgType, email, title, msg string) {
	host := "smtp.qiye.163.com"
	port := "25"
	emailUsr := "shlyz@huagai.com"
	emailPwd := "8215579lyz"

	auth := smtp.PlainAuth(
	"",
	emailUsr,
	emailPwd,
	host,
	)

	var content_type string

	if (msgtype=="html") {
		content_type = "Content-Type: text/html" + "; charset=UTF-8"
	}else {
		content_type = "Content-Type: text/text" + "; charset=UTF-8"
	}

	mailmsg := []byte("To: " + email + "\r\nFrom: " + emailUsr + "<" + emailUsr +
	">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" +  msg)

	err := smtp.SendMail(
	host+":"+port,
	auth,
	emailUsr,
	[]string{email},
	mailmsg,
	)

	if err != nil {
		beego.Debug(err)
	}
}



