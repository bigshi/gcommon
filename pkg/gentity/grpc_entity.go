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
	maxBodyLen = 512
)
const (
	MdKeyTraceId  = "trace-id"
	MdKeyDomainId = "domain-id"
	MdKeyAppId    = "app-id"
	MdKeyUserId   = "user-id"
)

func getBodyStr(body interface{}) string {
	bodyStr := fmt.Sprintf("%+v", body)
	if len(bodyStr) > maxBodyLen {
		return bodyStr[:maxBodyLen]
	}
	return bodyStr
}

func getMdOfClient(ctx context.Context) (context.Context, metadata.MD) {
	// 拿上游的ctx
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		traceId := strconv.FormatInt(time.Now().UnixMicro(), 10)[4:]
		return metadata.AppendToOutgoingContext(ctx, MdKeyTraceId, traceId), metadata.Pairs(MdKeyTraceId, traceId)
	}
	out := metadata.NewOutgoingContext(ctx, md.Copy())
	arr := md.Get(MdKeyTraceId)
	if arr == nil || len(arr) == 0 {
		traceId := strconv.FormatInt(time.Now().UnixMicro(), 10)[4:]
		md.Append(MdKeyTraceId, traceId)
		return metadata.AppendToOutgoingContext(out, MdKeyTraceId, traceId), md
	}
	traceId := arr[0]
	md.Append(MdKeyTraceId, traceId)
	return out, md
}

func getMdOfServer(ctx context.Context) metadata.MD {
	// 拿上游的ctx
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return metadata.MD{}
	}
	arr := md.Get(MdKeyTraceId)
	if arr == nil || len(arr) == 0 {
		md.Append(MdKeyTraceId, "")
	}
	return md
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
	md := getMdOfServer(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC SERVER] %s fail - md:%+v, req:%s, err:%v, stack:%s", info.FullMethod, md, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC SERVER] %s param - md:%+v, req:%s", info.FullMethod, md, reqStr)
	bt := time.Now()
	rsp, err := handler(ctx, req)
	if err != nil {
		glog.Errorf("[GRPC SERVER] %s fail - cost:%dms, md:%+v, req:%s, msg:%s", info.FullMethod, md, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(rsp)
		glog.Infof("[GRPC SERVER] %s success - cost:%dms, md:%+v, req:%s, rsp:%s", info.FullMethod, md, time.Since(bt).Milliseconds(), reqStr, rspStr)
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
	outCtx, md := getMdOfClient(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC CLIENT] %s fail - md:%+v, req:%s, err:%v, stack:%s", method, md, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC CLIENT] %s param - md:%+v, req:%s", method, md, reqStr)
	bt := time.Now()
	err := invoker(outCtx, method, req, reply, cc, opts...)
	if err != nil {
		glog.Errorf("[GRPC CLIENT] %s fail - cost:%dms, md:%+v, req:%s, msg:%s", method, md, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(reply)
		glog.Infof("[GRPC CLIENT] %s success - cost:%dms, md:%+v, req:%s, rsp:%s", method, md, time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return err
}
