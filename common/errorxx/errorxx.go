package errorxx

const defaultCode = 1001

type CodeError struct {
	Code int    `json:"err_code"`
	Msg  string `json:"err_msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"err_code"`
	Msg  string `json:"err_msg"`
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
