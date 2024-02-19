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

// GetNacosUsername
//
//	@Description: 获取nacos客户端账号
//	@return string
func GetNacosUsername() string {
	tmp := os.Getenv("NACOS_USERNAME")
	if tmp == "" {
		return "nacos"
	}
	return strings.Trim(tmp, " ")
}

// GetNacosPassword
//
//	@Description:获取nacos客户端密码
//	@return string
func GetNacosPassword() string {
	tmp := os.Getenv("NACOS_PASSWORD")
	if tmp == "" {
		return "nacos"
	}
	return strings.Trim(tmp, " ")
}
