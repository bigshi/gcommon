/**
 * Create Time:2023/11/10
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"github.com/qionggemens/gcommon/pkg/glog"
	"net/http"
	"runtime/debug"
	"time"
)

// HttpProxy http代理
type HttpProxy struct {
	interceptors []HandleInterceptor
	requestMap   map[string]func(rw http.ResponseWriter, req *http.Request)
}

// HandleInterceptor 拦截器
type HandleInterceptor interface {
	PreHandle(rw http.ResponseWriter, req *http.Request) bool
	PostHandle(rw http.ResponseWriter, req *http.Request)
}

func (pxy *HttpProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			glog.Errorf("process http request fail - path:%s, err:%v, stack:%s", req.URL.Path, e, string(debug.Stack()))
		}
	}()
	urlPath := req.URL.Path
	glog.Infof("ACCESS req begin - path:%s", urlPath)
	bt := time.Now()
	reqMapping, isOk := pxy.requestMap[urlPath]
	if !isOk {
		glog.Errorf("ACCESS req end - msg:404, path:%s", req.URL.Path)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	for i := 0; i < len(pxy.interceptors); i++ {
		interceptor := pxy.interceptors[i]
		isOk = interceptor.PreHandle(rw, req)
		if !isOk {
			glog.Errorf("ACCESS req end - msg:preHandle fail, path:%s, interceptor:%v", req.URL.Path, interceptor)
			return
		}
	}
	reqMapping(rw, req)
	for i := len(pxy.interceptors) - 1; i >= 0; i-- {
		interceptor := pxy.interceptors[i]
		interceptor.PostHandle(rw, req)
	}
	glog.Infof("ACCESS req end - cost:%dms, path:%s", time.Since(bt).Milliseconds(), req.URL.Path)
}
