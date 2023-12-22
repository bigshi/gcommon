package gutil

import (
	"github.com/qionggemens/gcommon/pkg/gentity"
	"strings"
)

// GetWhereSql
//
//	@Description: 获取where
//	@param queryList
//	@return string
//	@return []interface{}
func GetWhereSql(queryList []gentity.QueryCondition) (string, []interface{}) {
	var ands = make([]string, 0)
	var values = make([]interface{}, 0)
	if queryList == nil || len(queryList) == 0 {
		return "", values
	}

	for _, query := range queryList {
		k := query.QueryKey
		v := query.QueryValue
		ands = append(ands, k)
		if strings.Contains(k, "like") {
			v = "%" + v.(string) + "%"
		}
		values = append(values, v)
	}
	return strings.Join(ands, " and "), values
}
