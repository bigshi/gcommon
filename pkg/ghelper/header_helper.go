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
	}
}

func ToReqHeaderFromReq(req *http.Request) *gentity.ReqHeader {
	return &gentity.ReqHeader{
		DomainCode: req.Header.Get(gentity.HeaderKeyDomainCode),
		AppCode:    req.Header.Get(gentity.HeaderKeyAppCode),
		UserId:     req.Header.Get(gentity.HeaderKeyUserId),
	}
}

func WriteOkResponse(rw http.ResponseWriter, data interface{}, msg string) {
	io.WriteString(rw, string(gentity.NewOkResultBytes(data, msg)))
}

func WriteFailResponse(rw http.ResponseWriter, msg string) {
	io.WriteString(rw, string(gentity.NewFailResultBytes(msg)))
}
