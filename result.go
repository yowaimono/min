package min

const (
    CodeSuccess = 200
    CodeError   = 500
    CodeFail    = 400
)

// Result 是统一的响应体
type Result[T any] struct {
    Msg  string `json:"msg"`
    Code int    `json:"code"`
    Data T      `json:"data"`
}

// NewResult 创建一个新的 Result 实例
func NewResult[T any](msg string, code int, data T) Result[T] {
    return Result[T]{Msg: msg, Code: code, Data: data}
}

// Of 创建一个 Result 实例
func Of[T any](data T) Result[T] {
    return NewResult("成功", CodeSuccess, data)
}

// Error 创建一个 Result 实例表示错误
func Error(msg string) Result[any] {
    return NewResult[any](msg, CodeError, nil)
}

// Success 创建一个 Result 实例表示成功
func Success(data any) Result[any] {
    return Of(data)
}

// Fail 创建一个 Result 实例表示失败
func Fail(msg string) Result[any] {
    return NewResult[any](msg, CodeFail, nil)
}