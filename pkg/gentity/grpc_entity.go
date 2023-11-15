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
	"google.golang.org/grpc/metadata"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	maxBodyLen = 300
	TraceId    = "trace-id"
)

func getBodyStr(body interface{}) string {
	bodyStr := fmt.Sprintf("%+v", body)
	if len(bodyStr) > maxBodyLen {
		return bodyStr[:maxBodyLen]
	}
	return bodyStr
}

func GetValueFromMD(md metadata.MD, key string) string {
	arr := md.Get(key)
	if nil == arr || len(arr) == 0 {
		return ""
	}
	return arr[0]
}

func GetClientContext(ctx context.Context) context.Context {
	traceId := GetTraceId(ctx)
	md := metadata.Pairs(TraceId, traceId)
	return metadata.NewOutgoingContext(ctx, md)
}

func GetTraceId(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if exists {
		arr := md.Get(TraceId)
		if nil != arr && len(arr) > 0 {
			return arr[0]
		}
	}
	return strconv.FormatInt(time.Now().UnixMicro(), 10)
}

// GrpcServerInterceptor
//
//	@Description: 服务端拦截器
//	@param ctx
//	@param req
//	@param info
//	@param handler
//	@return interface{}
//	@return error
func GrpcServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqStr := getBodyStr(req)
	traceId := GetTraceId(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC SERVER] %s fail - traceId:%s, req:%s, err:%v, stack:%s", info.FullMethod, traceId, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC SERVER] %s param - traceId:%s, req:%s", info.FullMethod, traceId, reqStr)
	bt := time.Now()
	rsp, err := handler(ctx, req)
	if err != nil {
		glog.Errorf("[GRPC SERVER] %s fail - traceId:%s, cost:%dms, req:%s, msg:%s", info.FullMethod, traceId, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(rsp)
		glog.Infof("[GRPC SERVER] %s success - traceId:%s, cost:%dms, req:%s, rsp:%s", info.FullMethod, traceId, time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return rsp, err
}

// GrpcClientInterceptor
//
//	@Description: 客户端拦截器
//	@param ctx
//	@param method
//	@param req
//	@param reply
//	@param cc
//	@param invoker
//	@param opts
//	@return error
func GrpcClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	reqStr := getBodyStr(req)
	traceId := GetTraceId(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC CLIENT] %s fail - traceId:%s, req:%s, err:%v, stack:%s", method, traceId, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC CLIENT] %s param - traceId:%s, req:%s", method, traceId, reqStr)
	bt := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		glog.Errorf("[GRPC CLIENT] %s fail - traceId:%s, cost:%dms, req:%s, msg:%s", method, traceId, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(reply)
		glog.Infof("[GRPC CLIENT] %s success - traceId:%s, cost:%dms, req:%s, rsp:%s", method, traceId, time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return err
}
