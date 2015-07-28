package user
import (
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"common"
	"errors"
	"time"
	"strings"
	"enum"
	"models"
	"github.com/astaxie/beego"
	"fmt"
)

//增加用户
func AddUser(email, nickname, password string) enum.Result {
	o := orm.NewOrm()

	user := &User{Id:models.CurrVal("user"), Email:email, Nickname:nickname}
	user.Salt = common.RandString(10)
	user.Password = common.Md5(password + user.Salt)
	user.EmailToken = uuid.NewV4().String()
	_, err := o.Insert(user)

	fmt.Println(err)

	switch{
	case err==nil:
	case strings.Contains(err.Error(), "nick"):
		return enum.NickNameAlreadyExist
	case strings.Contains(err.Error(), "email"):
		return enum.EmailAlreadyExist
	case strings.Contains(err.Error(), "token")://uuid重复--基本不会发生
		AddUser(email, nickname, password)
	default:
	}

	models.NextVal("user")

	//创建用户成功后，联动生成token表
	for {
		userToken := &UserToken{Token:common.CreateGUID(), Updated:time.Now(), User:user}
		_, err=o.Insert(userToken)
		if err==nil {
			break
		}else {
			beego.Error(err)
		}
	}

	return enum.OK
}

//通过邮箱查找用户
func FindUser(email string) *User {
	o := orm.NewOrm()
	user := &User{Email: email}
	if o.Read(user, "email")!=nil {
		return nil
	}
	return user
}

//cookietoken更新
func UpdateCookieToken(user User) (string) {
	o := orm.NewOrm()

	var token string
	for {
		token = common.CreateGUID()
		_, err := o.QueryTable("user_token").Filter("user_id", user.Id).Update(orm.Params{"token": token, "updated":time.Now()})
		if err==nil {
			break
		}else {
			beego.Error(err)
		}
	}

	return token
}

//token验证
func VerifyToken(token string) (int64, bool) {
	o := orm.NewOrm()

	userToken := &UserToken{Token:token}
	err := o.Read(userToken, "token")
	if err!=nil {
		return 0, false
	}
	return userToken.User.Id, true
}

//todo
//获取邮箱认证
func ReqVerifyEmail(email string) (string, error) {
	user := &User{}
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Limit(-1)

	err := qs.Filter("email_verified", false).Filter("email", email).One(user)
	if err!=nil {
		return "", err
	}else {
		return user.EmailToken, nil
	}
}

//验证邮箱
//先列出所有未认证的邮箱，然后匹配emailToken
func VerifyEmail(emailToken string) error {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	qs.Limit(-1)
	num, err := qs.Filter("email_verified", false).Filter("email_token", emailToken).Update(orm.Params{"email_verified": true})
	if err!=nil {
		return err
	}else if num==0 {
		if qs.Filter("email_token", emailToken).Exist() {
			return errors.New(enum.EmailAlreadyVerified.String())
		}else {
			return errors.New(enum.NotVerifiedToken.String())
		}
	}else {
		return nil
	}
}

//请求重置密码
//func ReqResetPassword(email string) (string, error) {
//	o := orm.NewOrm()
//	user := &User{Email:email}
//
//	if o.Read(user, "email") == nil {
//		if !user.EmailVerified {
//			return "", errors.New(enum.EmailNeedVerify.String())
//		}
//		token := common.CreateGUID()
//		verify := &Verify{User:user, InsertTime:time.Now(), ResetToken:token}
//		if created, _, err := o.ReadOrCreate(verify, "user_id"); err == nil&&!created {
//			verify.ResetToken = token
//			verify.InsertTime = time.Now()
//			_, err := o.Update(verify, "token", "insert_time")
//			return token, err
//		}
//		return token, nil
//	}else {
//		return "", errors.New(enum.UserNotExist.String())
//	}
//}

//修改密码
//func ChangePassword(username, oldPassword, newPassword, repeatNewPassword string) error {
//	if newPassword!=repeatNewPassword {
//		return errors.New(enum.PasswordConflict.String())
//	}
//	user := &User{Username:username}
//	o := orm.NewOrm()
//
//	if o.Read(user, "username") != nil {
//		return errors.New(enum.UserNotExist.String())
//	}else if common.Md5(oldPassword + user.Salt)==user.Password {
//		user.Password = common.Md5(newPassword + user.Salt)
//		_, err := o.Update(user)
//		return err
//	}else {
//		return errors.New(enum.PasswordError.String())
//	}
//}

//判断重置密码的链接是否有效
func IsVerifyValid(token string) bool {
	o := orm.NewOrm()
	validTime := time.Now().Add(-time.Hour*8)
	return o.QueryTable("verify").Filter("insert_time__lt", validTime).Filter("token", token).Exist()
}

//重置密码
//func ResetPassword(token, password, confirmPassword string) error {
//	o := orm.NewOrm()
//	verify := &Verify{}
//	qs := o.QueryTable(verify)
//
//	if qs=qs.Filter("token", token); qs.One(verify)==nil {
//		validTime := time.Now().Add(-time.Hour*8)
//		if validTime.Sub(verify.InsertTime)<=0 {
//			salt := common.RandString(10)
//			password = common.Md5(password+salt)
//			o.QueryTable("user").Filter("Id", verify.User).Update(orm.Params{"password": password, "salt":salt})
//			return nil
//		}else {
//			return errors.New(enum.VerificationTimeOut.String())
//		}
//	}else {
//		return errors.New(enum.NotVerifiedToken.String())
//	}
//}

// 修改邮箱
//func ChangeEmail(username string, email string) error {
//	o := orm.NewOrm()
//
//	reg := regexp.MustCompile(`^(\w)+(\.\w+)*@(\w)+((\.\w+)+)$`)
//	result := reg.MatchString(email)
//	if !result {
//		return errors.New(enum.NotAnEmail.String())
//	}
//
//	num, err := o.QueryTable("user").Filter("username", username).Update(orm.Params{"email": email})
//
//	if nil != err {
//		return err
//	} else if 0 == num {
//		return errors.New(enum.UserNotExist.String())
//	}
//
//	return nil
//}

// 修改昵称
//func ChangeUsername(oldUsername, newUsername string) error {
//	o := orm.NewOrm()
//
//	num, err := o.QueryTable("user").Filter("username", oldUsername).Update(orm.Params{"username": newUsername})
//	if err!=nil {
//		return err
//	}else if num==0 {
//		return errors.New(enum.UserNotExist.String())
//	}
//	return nil
//}

//删除用户
//func DelUser(username, password string) error {
//	user := &User{Username:username}
//	o := orm.NewOrm()
//	if o.Read(user, "username") != nil {
//		return errors.New(enum.UserNotExist.String())
//	}else if common.Md5(password + user.Salt)==user.Password {
//		_, err := o.Delete(user)
//		return err
//	}else {
//		return errors.New(enum.PasswordError.String())
//	}
//}