/**
 * Create Time:2024/1/12
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package nacos

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/qionggemens/gcommon/pkg/glog"
	"gopkg.in/yaml.v2"
	"reflect"
	"strconv"
	"strings"
)

var configMap = make(map[string]interface{}, 0)

func LoadYamlConfig(namespaceId string, dataId string) error {
	cc := getClientConfig(namespaceId, dataId)
	sc := getServerConfig()
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		glog.Errorf("LoadYamlConfig fail - namespaceId:%s, dataId:%s, msg:%s", namespaceId, dataId, err.Error())
		return errors.New("LoadYamlConfig fail")
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		glog.Errorf("LoadYamlConfig fail - namespaceId:%s, dataId:%s, msg:%s", namespaceId, dataId, err.Error())
		return errors.New("LoadYamlConfig fail")
	}
	if content == "" {
		glog.Warningf("LoadYamlConfig finish - namespaceId:%s, dataId:%s, msg:config is empty", namespaceId, dataId)
		return nil
	}

	confMap := make(map[string]interface{}, 0)
	err = yaml.Unmarshal([]byte(content), confMap)
	if err != nil {
		return err
	}
	buildFlattenedMap(configMap, confMap, "")
	glog.Infof("-------------------------- nacos config --------------------------")
	for k, v := range configMap {
		glog.Infof("---  %s: %v", k, v)
	}
	glog.Infof("LoadYamlConfig finish - namespaceId:%s, dataId:%s", namespaceId, dataId)
	return nil
}

// buildFlattenedMap
//
//	@Description: Spring 原生转换
//	@param result
//	@param source
//	@param path
func buildFlattenedMap(result map[string]interface{}, source map[string]interface{}, path string) {
	for k, v := range source {
		if len(path) != 0 && path != "" {
			if strings.HasPrefix(k, "[") {
				k = path + k
			} else {
				k = path + "." + k
			}
		}
		vn := reflect.TypeOf(v).Kind()
		if vn == reflect.String {
			result[k] = v
		} else if vn == reflect.Map {
			value := v.(map[interface{}]interface{})
			son := make(map[string]interface{}, 0)
			for mk, mv := range value {
				son[mk.(string)] = mv
			}
			buildFlattenedMap(result, son, k)
		} else if vn == reflect.Array || vn == reflect.Slice {
			value := v.([]interface{})
			if len(value) == 0 {
				result[k] = ""
			} else {
				for i, j := range value {
					m := make(map[string]interface{}, 0)
					m["["+strconv.FormatInt(int64(i), 10)+"]"] = j
					buildFlattenedMap(result, m, k)
				}
			}
		} else {
			if v != nil {
				result[k] = v
			} else {
				result[k] = ""
			}
		}
	}
}

// GetBool
//
//	@Description: 获取bool值
//	@param key
//	@return bool
func GetBool(key string, defValue bool) bool {
	value, ok := configMap[key]
	if !ok {
		return defValue
	}
	val, isOk := value.(bool)
	if isOk {
		return val
	}
	return defValue
}

// GetString
//
//	@Description: 获取字符串值
//	@param key
//	@return string
func GetString(key string, defValue string) string {
	value, isOk := configMap[key]
	if !isOk {
		return defValue
	}
	val, isOk := value.(string)
	if isOk {
		return val
	}
	return defValue
}

// GetStrList
//
//	@Description: 获取字符串数组
//	@param key
//	@return []string
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

// GetInt
//
//	@Description:
//	@param key
//	@return int
func GetInt(key string, defValue int) int {
	value, isOk := configMap[key]
	if !isOk {
		return defValue
	}
	val, isOk := value.(int)
	if isOk {
		return val
	}
	return defValue
}

// GetInt32
//
//	@Description:
//	@param key
//	@return int32
func GetInt32(key string, defValue int32) int32 {
	value, isOk := configMap[key]
	if !isOk {
		return defValue
	}
	val, isOk := value.(int32)
	if isOk {
		return val
	}
	return defValue
}

// GetInt64
//
//	@Description:
//	@param key
//	@return int64
func GetInt64(key string, defValue int64) int64 {
	value, isOk := configMap[key]
	if !isOk {
		return defValue
	}
	val, isOk := value.(int64)
	if isOk {
		return val
	}
	return defValue
}
