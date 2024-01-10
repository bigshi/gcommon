/**
 * Create Time:2023/12/3
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

import (
	"context"
	"github.com/qionggemens/gcommon/pkg/component"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func GetDomainIdFromCtx(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(component.MdKeyDomainId)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func GetAppIdFromCtx(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(component.MdKeyAppId)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func GetUserIdFromCtx(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(component.MdKeyUserId)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func GetDomainIdFromHeader(req *http.Request) string {
	return req.Header.Get(HeaderKeyDomainId)
}

func GetAppIdFromHeader(req *http.Request) string {
	return req.Header.Get(HeaderKeyAppId)
}

func GetUserIdFromHeader(req *http.Request) string {
	return req.Header.Get(HeaderKeyUserId)
}
