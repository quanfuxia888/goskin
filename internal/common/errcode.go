package common

type ErrCode struct {
	Code   int
	MsgKey string
}

func (e ErrCode) Error() string {
	return e.MsgKey
}

func (e ErrCode) WithMsg(msg string) ErrCode {
	return ErrCode{Code: e.Code, MsgKey: msg}
}

// 注册示例
var (
	CodeSuccess          = newErr(0, "ok")
	ErrInternal          = newErr(1001, "internal_error")
	ErrInvalidParam      = newErr(1002, "invalid_param")
	ErrUnauthorized      = newErr(1003, "unauthorized")
	ErrUserAlreadyExists = newErr(2001, "user_exists")
	ErrUserNotFound      = newErr(2002, "user_not_found")
	ErrTokenGenerateFail = newErr(2003, "token_generate_fail")
	ErrTokenInvalid      = newErr(2004, "token_invalid")
	ErrTokenMissing      = newErr(2005, "token_missing")
)

func newErr(code int, msgKey string) ErrCode {
	RegisterCode(code, msgKey)
	return ErrCode{Code: code, MsgKey: msgKey}
}
