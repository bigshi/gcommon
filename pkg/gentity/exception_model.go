/**
 * Create Time:2023/4/14
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

type ApiException struct {
	Code int32
	Msg  string
}

func (error *ApiException) Error() string {
	return error.Msg
}

func NewApiException(code int32, msg string) *ApiException {
	return &ApiException{
		Code: code,
		Msg:  msg,
	}
}

func NewException(msg string) *ApiException {
	return &ApiException{
		Code: ErrorCode,
		Msg:  msg,
	}
}
