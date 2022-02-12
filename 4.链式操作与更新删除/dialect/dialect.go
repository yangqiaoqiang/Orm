package dialect

import (
	"reflect"
)

// Dialect 抽象出各个数据库差异的部分
var dialectsMap = map[string]Dialect{}

type Dialect interface {
	//将GO语言数据类型转换为该数据库数据类型
	DataTypeOf(typ reflect.Value) string
	//返回某个表是否存在的SQL语句，参数是表名
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect

}
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
