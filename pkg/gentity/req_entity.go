/**
 * Create Time:2024/2/20
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	HeaderKeyDomainCode = "Domain-Code"
	HeaderKeyAppCode    = "App-Code"
	HeaderKeyUserId     = "User-Id"
)

const (
	MdKeyTraceId    = "trace-id"
	MdKeyDomainCode = "domain-code"
	MdKeyAppCode    = "app-code"
	MdKeyUserId     = "user-id"
)

func GetMdValueFromCtx(ctx context.Context, mdKey string) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(mdKey)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

type ReqHeader struct {
	DomainCode string
	AppCode    string
	UserId     string
}

func ToReqHeaderFromCtx(ctx context.Context) *ReqHeader {
	return &ReqHeader{
		DomainCode: GetMdValueFromCtx(ctx, MdKeyDomainCode),
		AppCode:    GetMdValueFromCtx(ctx, MdKeyAppCode),
		UserId:     GetMdValueFromCtx(ctx, MdKeyUserId),
	}
}

func ToReqHeaderFromReq(req *http.Request) *ReqHeader {
	return &ReqHeader{
		DomainCode: req.Header.Get(HeaderKeyDomainCode),
		AppCode:    req.Header.Get(HeaderKeyAppCode),
		UserId:     req.Header.Get(HeaderKeyUserId),
	}
}
