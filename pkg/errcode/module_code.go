package errcode

var (
	UserAlreadyExist       = NewError(20001, "用户已存在")
	UsernameLengthLimit    = NewError(20002, "用户名长度3-21")
	UsernameCharacterLimit = NewError(20003, "用户名只能包含数字、字母")
	PasswordLengthLimit    = NewError(20004, "密码长度6-32")
	UserRegisterFailed     = NewError(20005, "注册失败")
	UserHasBeenBanned      = NewError(20006, "该账号已被封停")
)
