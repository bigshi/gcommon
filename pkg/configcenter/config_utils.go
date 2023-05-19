/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package configcenter

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
