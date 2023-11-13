/**
 * Create Time:2023/11/13
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"context"
	"fmt"
	"github.com/qionggemens/gcommon/pkg/glog"
	"google.golang.org/grpc"
	"runtime/debug"
	"time"
)

const maxBodyLen = 300

func getBodyStr(body interface{}) string {
	bodyStr := fmt.Sprintf("%+v", body)
	if len(bodyStr) > maxBodyLen {
		return bodyStr[:maxBodyLen]
	}
	return bodyStr
}

func GrpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqStr := getBodyStr(req)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("%s fail - req:%s, err:%v, stack:%s", info.FullMethod, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("%s param - req:%s", reqStr)
	bt := time.Now()
	rsp, err := handler(ctx, req)
	if err != nil {
		glog.Errorf("%s fail - cost:%dms, req:%s, msg:%s", time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(req)
		glog.Infof("%s success - cost:%dms, req:%s, rsp:%s", time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return rsp, err
}