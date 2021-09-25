// Package schema 参数封装（数据表映射）
package schema

// Topic 题目结构体
type Topic struct {
	ID    int
	Title string
}

// Paper 试卷结构体
type Paper struct {
	Name  string
	Topic []Topic
}
