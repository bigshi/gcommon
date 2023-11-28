/**
 * Create Time:2023/11/10
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"encoding/json"
	"github.com/qionggemens/gcommon/pkg/glog"
	"io"
	"net/http"
	"runtime/debug"
	"time"
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

// HttpServer http
type HttpServer struct {
	Interceptors []HandleInterceptor
	RequestMap   map[string]func(rw http.ResponseWriter, req *http.Request)
}

// HandleInterceptor 拦截器
type HandleInterceptor interface {
	PreHandle(rw http.ResponseWriter, req *http.Request) bool
	PostHandle(rw http.ResponseWriter, req *http.Request)
}

func (server *HttpServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			glog.Errorf("[HTTP ACCESS] handle fail - path:%s, err:%v, stack:%s", req.URL.Path, e, string(debug.Stack()))
		}
	}()
	urlPath := req.URL.Path
	glog.Infof("[HTTP ACCESS] req begin - path:%s", urlPath)
	bt := time.Now()
	reqMapping, isOk := server.RequestMap[urlPath]
	if !isOk {
		glog.Errorf("[HTTP ACCESS] req end - path:%s, msg:404", req.URL.Path)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	for i := 0; i < len(server.Interceptors); i++ {
		interceptor := server.Interceptors[i]
		isOk = interceptor.PreHandle(rw, req)
		if !isOk {
			glog.Errorf("[HTTP ACCESS] req end - path:%s, interceptor:%v, msg:preHandle fail", req.URL.Path, interceptor)
			return
		}
	}
	reqMapping(rw, req)
	for i := len(server.Interceptors) - 1; i >= 0; i-- {
		interceptor := server.Interceptors[i]
		interceptor.PostHandle(rw, req)
	}
	glog.Infof("[HTTP ACCESS] req end - cost:%dms, path:%s", time.Since(bt).Milliseconds(), req.URL.Path)
}
