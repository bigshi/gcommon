/**
 * Create Time:2024/1/10
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gutil

import (
	"errors"
	"net"
	"strings"
)

// GetLocalIpx
//
//	@Description: 获取ip地址
//	@param x
//	@return string
//	@return error
func GetLocalIpx(x int8) (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, ifs := range interfaces {
		if !strings.HasPrefix(ifs.Name, "en") && !strings.HasPrefix(ifs.Name, "eth") {
			continue
		}

		addrs, err := ifs.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ipnet, isOk := addr.(*net.IPNet)
			if !isOk || ipnet.IP.IsLoopback() {
				continue
			}
			if x == 4 && ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
			if x == 6 && ipnet.IP.To16() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("net interfaces not match")
}
