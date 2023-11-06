package gentity

const SelectSQLTemplate = "select %s from %s.%s %s %s %s" //  1-字段，2-库名，3-表名，4-where，5-order, 6-limit

const InsertSQLTemplate = "insert into %s.%s(%s) values(%s)" // 1-库名，2-表名，3-字段，4-值

const UpdateSQLTemplate = "update %s.%s set %s where %s" // 1-库名，2-表名，3-字段=值，4-where条件

const DeleteSQLTemplate = "delete from %s.%s where %s" // 1-库名，2-表名，3-where条件

const BatchInsertSQLTemplate = "insert into %s.%s(%s) values %s" // 1-库名，2-表名，3-字段，4-值

const ReplaceSQLTemplate = "replace into %s.%s(%s) values(%s)" // 1-库名，2-表名，3-字段，4-值

const Limit string = "limit"

const Offset string = "offset"

type QueryCondition struct {
	QueryKey   string
	QueryValue interface{}
}
