/**
 * Create Time:2023/12/3
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package ghelper

import (
	"context"
	"github.com/qionggemens/gcommon/pkg/gentity"
	"google.golang.org/grpc/metadata"
	"io"
	"net/http"
	"strings"
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

func ToReqHeaderFromCtx(ctx context.Context) *gentity.ReqHeader {
	return &gentity.ReqHeader{
		DomainCode: GetMdValueFromCtx(ctx, gentity.MdKeyDomainCode),
		AppCode:    GetMdValueFromCtx(ctx, gentity.MdKeyAppCode),
		UserId:     GetMdValueFromCtx(ctx, gentity.MdKeyUserId),
		SourceIp:   GetMdValueFromCtx(ctx, gentity.MdKeySourceIp),
	}
}

func ToReqHeaderFromReq(req *http.Request) *gentity.ReqHeader {
	// 来源IP
	sourceIps := req.Header.Get("x-forwarded-for")
	if sourceIps == "" {
		sourceIps = req.Header.Get("Proxy-Client-IP")
	}
	if sourceIps == "" {
		sourceIps = req.Header.Get("WL-Proxy-Client-IP")
	}
	if sourceIps == "" {
		sourceIps = req.Header.Get("HTTP_CLIENT_IP")
	}
	if sourceIps == "" {
		sourceIps = req.Header.Get("HTTP_X_FORWARDED_FOR")
	}
	if sourceIps == "" {
		sourceIps = req.RemoteAddr
	}
	return &gentity.ReqHeader{
		DomainCode: req.Header.Get(gentity.HeaderKeyDomainCode),
		AppCode:    req.Header.Get(gentity.HeaderKeyAppCode),
		UserId:     req.Header.Get(gentity.HeaderKeyUserId),
		SourceIp:   strings.Split(sourceIps, ",")[0],
	}
}

func WriteOkResponse(rw http.ResponseWriter, data interface{}, msg string) {
	io.WriteString(rw, string(gentity.NewOkResultBytes(data, msg)))
}

func WriteFailResponse(rw http.ResponseWriter, msg string) {
	io.WriteString(rw, string(gentity.NewFailResultBytes(msg)))
}
