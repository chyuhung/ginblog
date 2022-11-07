package errmsg

const (
	SUCCSE = 200
	ERROR  = 500
	// code=1000... model user error
	ERROR_USERNAME_UESD    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code=2000... model article error
	ERROR_ARTICLE_NOT_EXIST = 2001

	// code=3000... model category error
	ERROR_CATEGORY_USED      = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCSE: "OK",
	ERROR:  "FAIL",
	// user
	ERROR_USERNAME_UESD:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN错误",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	// category
	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
	// article
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
