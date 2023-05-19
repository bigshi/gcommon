/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package configcenter

import (
	"fmt"
	"strconv"
)

//
// GetBool
//  @Description: 获取bool值
//  @param key
//  @return bool
//
func GetBool(key string) bool {
	value, ok := configMap[key]
	if !ok {
		return false
	}
	val, isOk := value.(bool)
	if isOk {
		return val
	}
	return false
}

//
// GetString
//  @Description: 获取字符串值
//  @param key
//  @return string
//
func GetString(key string) string {
	value, isOk := configMap[key]
	if !isOk {
		return ""
	}
	val, isOk := value.(string)
	if isOk {
		return val
	}
	return ""
}

//
// GetStrList
//  @Description: 获取字符串数组
//  @param key
//  @return []string
//
func GetStrList(key string) []string {
	result := make([]string, 0)
	i := int64(0)
	for true {
		k := fmt.Sprintf("%s[%s]", key, strconv.FormatInt(i, 10))
		value, isOk := configMap[k]
		if !isOk {
			break
		}
		val, isOk := value.(string)
		if isOk {
			result = append(result, val)
		} else {
			break
		}
		i++
	}
	return result
}
