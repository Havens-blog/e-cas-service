// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysAPI = "sys_apis"

// SysAPI mapped from table <sys_apis>
type SysAPI struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	Path        string         `gorm:"column:path;not null" json:"path"`                  // api路径
	Description string         `gorm:"column:description;not null" json:"description"`    // api中文描述
	APIGroup    string         `gorm:"column:api_group;not null" json:"api_group"`        // api组
	Method      string         `gorm:"column:method;not null" json:"method"`              // 方法
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
}

// TableName SysAPI's table name
func (*SysAPI) TableName() string {
	return TableNameSysAPI
}
