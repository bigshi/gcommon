package gutil

import (
	"database/sql"
	"fmt"
	"github.com/frameundulate/gcommon/pkg/gmodel"
	"strings"
)

const Limit string = "limit"
const Offset string = "offset"

//
// GetWhereSql
//  @Description: 获取where
//  @param queryList
//  @return string
//  @return []interface{}
//
func GetWhereSql(queryList []gmodel.QueryCondition) (string, []interface{}) {
	var ands = make([]string, 0)
	var values = make([]interface{}, 0)

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

//
// GetOrderSql
//  @Description: 获取order
//  @param orderMap
//  @return string
//
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

//
// GetLimitSql
//  @Description: 获取limit
//  @param limitMap
//  @return string
//
func GetLimitSql(limitMap map[string]int32) string {
	if limitMap == nil || len(limitMap) == 0 {
		return ""
	}
	limit, limitOk := limitMap[Limit]
	offset, offsetOk := limitMap[Offset]
	if limitOk && offsetOk {
		return fmt.Sprintf("limit %d offset %d", limit, offset)
	}
	if limitOk && !offsetOk {
		return fmt.Sprintf("limit %d", limit)
	}
	return ""
}

//
// Insert
//  @Description: 新增
//  @param db
//  @param sql
//  @param args
//  @return err
//  @return id
//
func Insert(db *sql.DB, sql string, args ...interface{}) (err error, id int64) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err, 0
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, insertId
}

//
// InsertByTx
//  @Description: 新增-事务
//  @param tx
//  @param sql
//  @param args
//  @return err
//  @return id
//
func InsertByTx(tx *sql.Tx, sql string, args ...interface{}) (err error, id int64) {
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return err, 0
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, insertId
}

//
// Modify
//  @Description: 修改（编辑、删除）
//  @param db
//  @param sql
//  @param args
//  @return err
//  @return id
//
func Modify(db *sql.DB, sql string, args ...interface{}) (err error, id int64) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err, 0
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err, 0
	}
	return nil, affectedRows
}

//
// ModifyTx
//  @Description: 修改（编辑、删除）- 事务
//  @param tx
//  @param sql
//  @param args
//  @return err
//  @return id
//
func ModifyTx(tx *sql.Tx, sql string, args ...interface{}) (err error, id int64) {
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return err, 0
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err, 0
	}
	return nil, affectedRows
}
