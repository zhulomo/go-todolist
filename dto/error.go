package dto

type AppError struct {
	Code int
	Msg  string
	Err  error // 原始错误 日志用
}

func (e AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}
