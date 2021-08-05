package gormtogo

import (
	"fmt"
	"strings"

	"github.com/yishuihanj/db2go/findSql"
)

var Gorm bool

//向字段中添加 Gorm tag
func AddGormTag(column *findSql.Column) string {
	//如果没有开启gorm 则 不需要转换
	if !Gorm {
		return ""
	}
	tag := fmt.Sprintf("\t`gorm:\"column:%s", column.ColumnName)
	if column.NotNull == "true" {
		tag += fmt.Sprintf(";not null")
	}
	if column.DefaultValue != "" && column.IsPrimaryKey != "true" {
		if strings.Contains(column.DefaultValue, "''") { //如果有 ''则置为 空字符串
			column.DefaultValue = fmt.Sprintf("''")
		}
		tag += fmt.Sprintf(";default:%s", column.DefaultValue)
	}
	if column.IsPrimaryKey == "true" {
		tag += fmt.Sprintf(";primaryKey")
	}
	if column.ColumnType != "" {
		tag += fmt.Sprintf(";type:%s", column.ColumnType)
	}
	if column.Comment != "" {
		tag += fmt.Sprintf(";comment:'%s'", column.Comment)
	}
	end := fmt.Sprintf("\"`")
	return tag + end
}
