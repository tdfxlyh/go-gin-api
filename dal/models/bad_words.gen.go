// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameBadWord = "bad_words"

// BadWord mapped from table <bad_words>
type BadWord struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                           // 主键id
	Content      string    `gorm:"column:content" json:"content"`                                               // 敏感词内容
	CreateUserID int64     `gorm:"column:create_user_id" json:"create_user_id"`                                 // 创建用户id
	Extra        string    `gorm:"column:extra" json:"extra"`                                                   // 扩展信息
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3)" json:"create_time"` // 创建时间
	ModifyTime   time.Time `gorm:"column:modify_time;not null;default:CURRENT_TIMESTAMP(3)" json:"modify_time"` // 修改时间
	Status       int32     `gorm:"column:status" json:"status"`                                                 // 0存在，1删除
}

// TableName BadWord's table name
func (*BadWord) TableName() string {
	return TableNameBadWord
}
