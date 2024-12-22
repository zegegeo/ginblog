package biz

type Error struct{
	Code int    `json:"code"`
	Msg string  `json:"message"`
}

func NewError(code int, msg string) *Error{
	return &Error{
		Code: code,
		Msg: msg,
	}
}

func (e *Error) Error() string{
	return e.Msg
}