package errcode

var (
	ErrorService                  = NewError(20000, "服务异常，请稍后重试")
	ErrorUserRecordNotFound       = NewError(20001, "查询用户不存在")
	ErrorPlatformRecordNotFound   = NewError(20002, "查询平台不存在")
	ErrorUserPhoneExists          = NewError(20003, "当前手机号已绑定账户")
	ErrorUserEmailExists          = NewError(20004, "当前邮箱已绑定账户")
	ErrorUserNameExists           = NewError(20005, "当前用户名已绑定账户")
	ErrorUserAuthFailed           = NewError(20006, "用户验证失败, 请输入正确的账号和密码")
	ErrorVerifyCodeNoPhoneNumbers = NewError(20007, "输入正确的手机号")
	ErrorGenerateVerifyCode       = NewError(20008, "生成验证码失败")
	ErrorSendVerifyCode           = NewError(20009, "发送验证码失败")
	ErrorCheckCode                = NewError(20010, "验证异常")
	ErrorVerifyCodeFailed         = NewError(20011, "验证失败")
	ErrorUnknownService           = NewError(29999, "未知的服务异常，请联系负责人排查异常")
)
