package models

// 数据库查询的结果集类型定义
type Result struct {
	TotalCount int
	TotalPages int
	List       []string
}
