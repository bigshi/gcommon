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

func grpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqStr := fmt.Sprintf("%+v", req)[:500]
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
		rspStr := fmt.Sprintf("%+v", rsp)[:500]
		glog.Infof("%s success - cost:%dms, req:%s, rsp:%s", time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return rsp, err
}
