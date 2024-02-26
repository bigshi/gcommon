/**
 * Create Time:2024/2/20
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gentity

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
	MdKeySourceIp   = "source-ip"
)

type ReqHeader struct {
	DomainCode string
	AppCode    string
	UserId     string
	SourceIp   string
}
