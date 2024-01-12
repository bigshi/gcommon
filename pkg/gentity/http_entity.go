/**
 * Create Time:2023/11/10
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"encoding/json"
	"io"
	"net/http"
)

const ErrorCode = 500
const SuccessCode = 200

const (
	HeaderKeyDomainId = "Domain-Id"
	HeaderKeyAppId    = "App-Id"
	HeaderKeyUserId   = "User-Id"
)

type ApiResult struct {
	Code    int32       `json:"code"`
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func NewOkResult(data interface{}, msg string) ApiResult {
	result := ApiResult{
		Msg:     msg,
		Code:    SuccessCode,
		Success: true,
		Data:    data,
	}
	return result
}

func NewOkResultBytes(data interface{}, msg string) []byte {
	bytes, _ := json.Marshal(NewOkResult(data, msg))
	return bytes
}

func NewFailResult(msg string) ApiResult {
	result := ApiResult{
		Msg:     msg,
		Code:    ErrorCode,
		Success: false,
	}
	return result
}

func NewFailResultBytes(msg string) []byte {
	bytes, _ := json.Marshal(NewFailResult(msg))
	return bytes
}

func WriteOkResponse(rw http.ResponseWriter, data interface{}, msg string) {
	io.WriteString(rw, string(NewOkResultBytes(data, msg)))
}

func WriteFailResponse(rw http.ResponseWriter, msg string) {
	io.WriteString(rw, string(NewFailResultBytes(msg)))
}
