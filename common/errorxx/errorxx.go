package errorxx

func NewCustomError(code int, msg string) error {
	return ApiCustomError{
		Msg:  msg,
		Code: code,
	}
}

type ApiCustomError struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"err_msg"`
	Code int         `json:"err_code"`
}

func (api ApiCustomError) Error() string {
	return api.Msg
}

func (api ApiCustomError) ErrCode() int {
	return api.Code
}
