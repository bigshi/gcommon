/**
 * Create Time:2023/10/13
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gutil

import (
	"os"
	"strings"
)

// GetNacosAddr
//
//	@Description: 获取 nacos 配置地址
//	@return string
func GetNacosAddr() string {
	tmp := os.Getenv("NACOS_ADDR")
	if tmp == "" {
		return "nacos.qionggemen.com:8848"
	}
	return strings.Trim(tmp, " ")
}
