package retcode

const (
	SUCCESS = 0
	Fail    = -1

	//-10~-99表示系统级错误
	ErrParam = -10 //参数错误

	//-100~-999表示业务逻辑
	ErrValidateFailUrl = -101 //不是合法的URL
	ErrHashIdNotFound  = -102 //不存在该HashId
	ErrUrlNotFound     = -103 //URI不存在
	TokenNotFound      = -104 //Token不存在
)
