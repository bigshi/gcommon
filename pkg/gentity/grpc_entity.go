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
)
const (
	TraceId  = "trace-id"
	DomainId = "domain-id"
	AppId    = "app-id"
	UserId   = "user-id"
)

func getBodyStr(body interface{}) string {
	bodyStr := fmt.Sprintf("%+v", body)
	if len(bodyStr) > maxBodyLen {
		return bodyStr[:maxBodyLen]
	}
	return bodyStr
}

func getTraceIdOfClient(ctx context.Context) (context.Context, string) {
	// 拿上游的ctx
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		traceId := strconv.FormatInt(time.Now().UnixMicro(), 10)[4:]
		return metadata.AppendToOutgoingContext(ctx, TraceId, traceId), traceId
	}
	out := metadata.NewOutgoingContext(ctx, md.Copy())
	arr := md.Get(TraceId)
	if arr == nil || len(arr) == 0 {
		traceId := strconv.FormatInt(time.Now().UnixMicro(), 10)[4:]
		return metadata.AppendToOutgoingContext(out, TraceId, traceId), traceId
	}
	traceId := arr[0]
	return out, traceId
}

func getTraceIdOfServer(ctx context.Context) string {
	// 拿上游的ctx
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(TraceId)
	if arr == nil || len(arr) == 0 {
		return ""
	}
	return arr[0]
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
	traceId := getTraceIdOfServer(ctx)
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
	outCtx, traceId := getTraceIdOfClient(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC CLIENT] %s fail - traceId:%s, req:%s, err:%v, stack:%s", method, traceId, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC CLIENT] %s param - traceId:%s, req:%s", method, traceId, reqStr)
	bt := time.Now()
	err := invoker(outCtx, method, req, reply, cc, opts...)
	if err != nil {
		glog.Errorf("[GRPC CLIENT] %s fail - traceId:%s, cost:%dms, req:%s, msg:%s", method, traceId, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(reply)
		glog.Infof("[GRPC CLIENT] %s success - traceId:%s, cost:%dms, req:%s, rsp:%s", method, traceId, time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return err
}

type GrpcHeader struct {
	AppId    string
	DomainId string
	UserId   string
}

func ToGrpcHeader(ctx context.Context) *GrpcHeader {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return &GrpcHeader{}
	}
	var domainId, appId, userId string

	arr := md.Get(DomainId)
	if arr != nil && len(arr) > 0 {
		domainId = arr[0]
	}

	arr = md.Get(AppId)
	if arr != nil && len(arr) > 0 {
		appId = arr[0]
	}

	arr = md.Get(UserId)
	if arr != nil && len(arr) > 0 {
		userId = arr[0]
	}
	return &GrpcHeader{AppId: appId, DomainId: domainId, UserId: userId}
}
