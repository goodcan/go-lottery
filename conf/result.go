package conf

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this *Result) SetError(code int, msg string, data ...interface{}) {
	this.Code = code
	this.Msg = msg
	if data != nil {
		this.Data = data
	}
}
