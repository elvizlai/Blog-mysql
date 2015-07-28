package enum

type Result int

const (
	OK Result = iota
	UserNotExist	//用户名不存在
	NickNameAlreadyExist	//昵称已存在
	PasswordError	//密码错误
	PasswordConflict
	EmailAlreadyExist	//邮箱已存在
	EmailNeedVerify		//邮箱需要验证
	EmailAlreadyVerified	//邮箱已验证
	NotAnEmail
	NotVerifiedToken	//token非法
	VerificationTimeOut		//已经超时
	CategoryAlreadyExist	//分类已经存在
)

func (r Result) String() string {
	switch r{
		case OK:
		return "请求成功"

		case UserNotExist:
		return "不存在此用户"

		case NickNameAlreadyExist:
		return "昵称已存在"

		case PasswordError:
		return "密码错误"
		case PasswordConflict:
		return "两次密码输入不一致"

		case EmailAlreadyExist:
		return "邮箱已存在"
		case EmailNeedVerify:
		return "邮箱未激活"
		case EmailAlreadyVerified:
		return "邮箱已激活"

		case NotAnEmail:
		return "邮箱格式不正确"
		case NotVerifiedToken:
		return "验证失败，链接无效"
		case VerificationTimeOut:
		return "验证码已失效"

		case CategoryAlreadyExist:
		return "该分类已存在"

		default:
		return "未知错误"
	}
}