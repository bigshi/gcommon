/**
 * Create Time:2023/11/10
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"encoding/json"
)

const ErrorCode = 500
const SuccessCode = 200

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
