/**
 * Create Time:2023/12/3
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package ghelper

import (
	"context"
	"github.com/qionggemens/gcommon/pkg/gcomponent"
	"github.com/qionggemens/gcommon/pkg/gentity"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func GetDomainCodeFromCtx(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(gcomponent.MdKeyDomainCode)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func GetAppCodeFromCtx(ctx context.Context) string {
	md, exists := metadata.FromIncomingContext(ctx)
	if !exists {
		return ""
	}
	arr := md.Get(gcomponent.MdKeyAppCode)
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
	arr := md.Get(gcomponent.MdKeyUserId)
	if arr != nil && len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func GetDomainCodeFromHeader(req *http.Request) string {
	return req.Header.Get(gentity.HeaderKeyDomainId)
}

func GetAppCodeFromHeader(req *http.Request) string {
	return req.Header.Get(gentity.HeaderKeyAppId)
}

func GetUserIdFromHeader(req *http.Request) string {
	return req.Header.Get(gentity.HeaderKeyUserId)
}
