package utils

// 响应结构体
type BusinessCode int

type Result struct {
	Code    BusinessCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data,omitempty"` // 使用omitempty，在data为nil时不生成该字段
}

func (r *Result) Success(data interface{}) *Result {
	r.Code = 200
	r.Message = "Success ok 🚀"
	r.Data = nil // 清空原先的数据
	if data != nil {
			r.Data = data
	}
	return r
}

func (r *Result) Fail(code BusinessCode, message string) *Result {
	r.Code = code
	r.Message = message
	r.Data = nil // 清空原先的数据
	return r
}
