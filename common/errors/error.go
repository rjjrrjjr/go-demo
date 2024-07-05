package errs

import (
	"fmt"
	"io"
	"net/http"
	"runtime"

	"git.code.oa.com/trpc-go/trpc-go/errs"
	"git.code.oa.com/trpc-go/trpc-go/log"
	"github.com/pkg/errors"
	"go.uber.org/multierr"
)

const (
	CodeInternalError = 10000
)

type ErrCode = int

var statusCodeMap = make(map[ErrCode]int)
var MsgMap = make(map[ErrCode]string)

// BizErr 同时支持 trpc 和 restful 的错误结构
// 为了解决 error 空值判断问题，这里的类型不实现 Error 接口，要转成 error 手动调用 AsError 方法
type BizErr struct {
	code     ErrCode
	msg      string
	cause    error
	httpCode int
}

type MsgContent struct {
	Msg  string `json:"msg"`
	Desc string `json:"desc"`
}

type MsgStatusOk struct {
	ErrorCode int32       `json:"errorCode"`
	Body      interface{} `json:"body"`
}

func GetMsgStatusOk(msg interface{}) *MsgStatusOk {
	return &MsgStatusOk{
		ErrorCode: 0,
		Body:      msg,
	}
}

// AsError 转成 error 接口类型，直接返回 trpc.Error
func (e *BizErr) AsError() error {
	if e.cause != nil {
		return errs.Wrap(e.cause, e.code, e.msg)
	}
	return errs.New(e.code, e.msg)
}

func (e *BizErr) Error() string {
	return e.msg
}

func (e *BizErr) AsCauseError() error {
	if e.cause != nil {
		return errs.Wrap(e.cause, e.code, e.msg)
	}
	return errs.New(e.code, e.msg)
}

// StatusCode 状态码
func (e *BizErr) StatusCode() int {
	return e.httpCode
}

// Cause 返回实际的 error 错误
func (e *BizErr) Cause() error {
	return e.cause
}

// Unwrap 用于 errors.Unwrap 检测封装类型的接口
func (e *BizErr) Unwrap() error {
	return e.cause
}

// ErrCode 错误码
func (e *BizErr) ErrCode() ErrCode {
	return e.code
}

// ErrMsg 错误信息
func (e *BizErr) ErrMsg() string {
	return e.msg
}

// String 格式化输出
func (e *BizErr) String() string {
	return fmt.Sprintf("BizErr{code:%d, msg:%s, httpcode:%d, err:%v}",
		e.code, e.msg, e.httpCode, e.Cause())
}

func GetStatusCodeByErrorCode(eCode int) int {
	statusCode, ok := statusCodeMap[eCode]
	if !ok {
		statusCode = 500
	}
	return statusCode
}

// Format 格式化输出时用
func (e *BizErr) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		// +v% 或 5xx 错误打印堆栈
		if s.Flag('+') || e.httpCode >= http.StatusInternalServerError {
			_, _ = fmt.Fprintf(s, "BizErr{code:%d, msg:%s, httpcode:%d, err:%+v}",
				e.code, e.msg, e.httpCode, e.Cause())
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.String())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.String())
	default:
		// unknown format
		_, _ = fmt.Fprintf(s, "%%!%c(BizErr=%s)", verb, e.String())
	}
}

// trpcInternalHttpStatusMap 默认的 tRPC 错误码到 http 状态码的映射
// Copy from git.code.oa.com/trpc-go/trpc-go/restful/errors.go
var trpcInternalHttpStatusMap = map[int32]int{
	errs.RetServerDecodeFail:   http.StatusBadRequest,
	errs.RetServerEncodeFail:   http.StatusInternalServerError,
	errs.RetServerNoService:    http.StatusNotFound,
	errs.RetServerNoFunc:       http.StatusNotFound,
	errs.RetServerTimeout:      http.StatusGatewayTimeout,
	errs.RetServerOverload:     http.StatusTooManyRequests,
	errs.RetServerSystemErr:    http.StatusInternalServerError,
	errs.RetServerAuthFail:     http.StatusUnauthorized,
	errs.RetServerValidateFail: http.StatusBadRequest,
	errs.RetUnknown:            http.StatusInternalServerError,
	errs.RetClientThrottled:    http.StatusTooManyRequests,
}

// GetStatusCode 提取错误中的 Http 状态码
func GetStatusCode(err error) int {
	err = unwrapMultiErr(err)
	statusCode := http.StatusInternalServerError
	code := errs.Code(err)
	if statusFromMap, ok := statusCodeMap[ErrCode(code)]; ok {
		statusCode = statusFromMap
	}
	log.Debugf("GetStatusCode- code: %v, statusCodeMap got statusCode:%v ", code, statusCode)
	if statusFromMap, ok := trpcInternalHttpStatusMap[int32(code)]; ok {
		statusCode = statusFromMap
	}
	log.Debugf("GetStatusCode- code: %v, trpcInternalHttpStatusMap got statusCode:%v ", code, statusCode)
	return statusCode
}

func GetErrCode(err error) int {
	if err == nil {
		return 0
	}
	err = unwrapMultiErr(err)
	for {
		if e, ok := err.(*errs.Error); ok {
			return int(e.Code)
		}
		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}
	return CodeInternalError
}

// GetErrorCodeMessage 获取错误中的 code/message 结构
func GetErrorCodeMessage(err error) (code int, msg string) {
	err = unwrapMultiErr(err)
	return errs.Code(err), errs.Msg(err)
}

// 针对 multierr，且只有一个 err，则展开并返回第一个
// 如果有多个错误，则不做处理，code 和 msg 都将当做内部错误来处理
func unwrapMultiErr(err error) error {
	merr := multierr.Errors(err)
	if len(merr) == 1 {
		err = merr[0]
	}
	return err
}

func IsAny(err error, codes ...ErrCode) bool {
	err = unwrapMultiErr(err)
	code := errs.Code(err)
	for _, c := range codes {
		if c == code {
			return true
		}
	}
	return false
}

func Msg(err error) string {
	return errs.Msg(err)
}

func Code(err error) int {
	return errs.Code(err)
}

// Not thread-safe
func register(code ErrCode, statusCode int, msg string) {
	statusCodeMap[code] = statusCode
	MsgMap[code] = msg
}

func NewError(code ErrCode, cause error, args ...interface{}) *BizErr {
	return newError(code, cause, args...)
}

func newError(code ErrCode, cause error, args ...interface{}) *BizErr {
	statusCode, ok := statusCodeMap[code]
	var msg string
	if !ok {
		statusCode = statusCodeMap[CodeInternalError]
		msg = fmt.Sprintf(MsgMap[CodeInternalError], args...)
	} else {
		msg = fmt.Sprintf(MsgMap[code], args...)
	}
	return &BizErr{
		code:     code,
		msg:      msg,
		cause:    newWithStack(cause),
		httpCode: statusCode,
	}
}

func newWithHttpCode(httpCode int, code ErrCode, cause error, args ...interface{}) *BizErr {
	var msg string
	if _, ok := MsgMap[code]; !ok {
		msg = fmt.Sprintf(MsgMap[CodeInternalError], args...)
	} else {
		msg = fmt.Sprintf(MsgMap[code], args...)
	}
	return &BizErr{
		code:     code,
		msg:      msg,
		cause:    newWithStack(cause),
		httpCode: httpCode,
	}
}

func newWithStack(err error) error {
	if err == nil {
		return nil
	}
	return &withStack{
		err,
		callers(2),
	}
}

type stack []uintptr

type withStack struct {
	error
	stack
}

// Cause 实际错误
func (w *withStack) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withStack) Unwrap() error { return w.error }

// Format 格式化
func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v", w.Cause())
			for _, pc := range w.stack {
				_, _ = fmt.Fprintf(s, "\n%+v", errors.Frame(pc))
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}

}

func callers(skip int) stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3+skip, pcs[:])
	return pcs[0:n]
}
