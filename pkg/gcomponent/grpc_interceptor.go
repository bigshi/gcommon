/**
 * Create Time:2023/11/13
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gcomponent

import (
	"context"
	"fmt"
	"github.com/qionggemens/gcommon/pkg/gentity"
	"github.com/qionggemens/gcommon/pkg/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"runtime/debug"
	"strconv"
	"time"
)

const (
	maxBodyLen = 512
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
		return metadata.AppendToOutgoingContext(ctx, gentity.MdKeyTraceId, traceId), metadata.Pairs(gentity.MdKeyTraceId, traceId)
	}
	out := metadata.NewOutgoingContext(ctx, md.Copy())
	arr := md.Get(gentity.MdKeyTraceId)
	if arr == nil || len(arr) == 0 {
		traceId := strconv.FormatInt(time.Now().UnixMicro(), 10)[4:]
		md.Append(gentity.MdKeyTraceId, traceId)
		return metadata.AppendToOutgoingContext(out, gentity.MdKeyTraceId, traceId), md
	}
	traceId := arr[0]
	md.Append(gentity.MdKeyTraceId, traceId)
	return out, md
}

func getMdOfServer(ctx context.Context) metadata.MD {
	// 拿上游的ctx
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return metadata.MD{}
	}
	arr := md.Get(gentity.MdKeyTraceId)
	if arr == nil || len(arr) == 0 {
		md.Append(gentity.MdKeyTraceId, "")
	}
	return md
}

func getAddr(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ""
	}
	if p.Addr == net.Addr(nil) {
		return ""
	}
	return p.Addr.String()
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
	clientAddr := getAddr(ctx)
	md := getMdOfServer(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC SERVER] %s fail [From:%s] - md:%+v, req:%s, err:%v, stack:%s", info.FullMethod, clientAddr, md, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC SERVER] %s begin [From:%s] - md:%+v, req:%s", info.FullMethod, clientAddr, md, reqStr)
	bt := time.Now()
	rsp, err := handler(ctx, req)
	if err != nil {
		glog.Errorf("[GRPC SERVER] %s fail [From:%s] - cost:%dms, md:%+v, req:%s, msg:%s", info.FullMethod, clientAddr, md, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(rsp)
		glog.Infof("[GRPC SERVER] %s success [From:%s] - cost:%dms, md:%+v, req:%s, rsp:%s", info.FullMethod, clientAddr, md, time.Since(bt).Milliseconds(), reqStr, rspStr)
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
	serverAddr := getAddr(ctx)
	outCtx, md := getMdOfClient(ctx)
	defer func() {
		if p := recover(); p != nil {
			glog.Errorf("[GRPC CLIENT] %s fail [To:%s] - md:%+v, req:%s, err:%v, stack:%s", method, serverAddr, md, reqStr, p, string(debug.Stack()))
		}
	}()
	glog.Infof("[GRPC CLIENT] %s begin [To:%s] - md:%+v, req:%s", method, serverAddr, md, reqStr)
	bt := time.Now()
	err := invoker(outCtx, method, req, reply, cc, opts...)
	if err != nil {
		glog.Errorf("[GRPC CLIENT] %s fail [To:%s] - cost:%dms, md:%+v, req:%s, msg:%s", method, serverAddr, md, time.Since(bt).Milliseconds(), reqStr, err.Error())
	} else {
		rspStr := getBodyStr(reply)
		glog.Infof("[GRPC CLIENT] %s success [To:%s] - cost:%dms, md:%+v, req:%s, rsp:%s", method, serverAddr, md, time.Since(bt).Milliseconds(), reqStr, rspStr)
	}
	return err
}
