package gutil

import (
	"fmt"
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
	return fmt.Sprintf("where %s", strings.Join(ands, " and ")), values
}

// GetOrderSql
//
//	@Description: 获取order
//	@param orderMap
//	@return string
func GetOrderSql(orderMap map[string]bool) string {
	if orderMap == nil || len(orderMap) == 0 {
		return ""
	}
	orderBys := make([]string, 0)
	for k, v := range orderMap {
		if v {
			orderBys = append(orderBys, k+" asc")
		} else {
			orderBys = append(orderBys, k+" desc")
		}
	}
	return fmt.Sprintf("order by %s", strings.Join(orderBys, ","))
}

// GetLimitSql
//
//	@Description: 获取limit
//	@param limitMap
//	@return string
func GetLimitSql(limitMap map[string]int32) string {
	if limitMap == nil || len(limitMap) == 0 {
		return ""
	}
	limit, limitOk := limitMap[gentity.Limit]
	offset, offsetOk := limitMap[gentity.Offset]
	if limitOk && offsetOk {
		return fmt.Sprintf("limit %d offset %d", limit, offset)
	}
	if limitOk && !offsetOk {
		return fmt.Sprintf("limit %d", limit)
	}
	return ""
}

// GetSelectSql
//
//	@Description: 获取查询sql
//	@param dbName
//	@param tbName
//	@param selectFieldStr
//	@param queryList
//	@param orderMap
//	@param limitMap
//	@return string
//	@return []interface{}
func GetSelectSql(dbName string, tbName string, selectFieldStr string, queryList []gentity.QueryCondition, orderMap map[string]bool, limitMap map[string]int32) (string, []interface{}) {
	whereSql, whereValues := GetWhereSql(queryList)
	selectSql := fmt.Sprintf(gentity.SelectSQLTemplate, selectFieldStr, dbName, tbName, whereSql, GetOrderSql(orderMap), GetLimitSql(limitMap))
	return strings.Trim(selectSql, " "), whereValues
}
