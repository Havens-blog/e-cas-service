// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysDept = "sys_depts"

// SysDept mapped from table <sys_depts>
type SysDept struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	ParentID  int64          `gorm:"column:parent_id;not null" json:"parent_id"`        // 上级部门
	DeptPath  string         `gorm:"column:dept_path;not null" json:"dept_path"`        // 部门路径
	DeptName  string         `gorm:"column:dept_name;not null" json:"dept_name"`        // 部门名称
	Sort      int32          `gorm:"column:sort;not null" json:"sort"`                  // 排序
	Leader    string         `gorm:"column:leader;not null" json:"leader"`              // 负责人
	Phone     string         `gorm:"column:phone;not null" json:"phone"`                // 手机
	Email     string         `gorm:"column:email;not null" json:"email"`                // 邮箱
	Status    int32          `gorm:"column:status;not null;default:1" json:"status"`    // 状态 1=正常 2-冻结
	CreateBy  string         `gorm:"column:create_by;not null" json:"create_by"`        // 创建人
	UpdateBy  string         `gorm:"column:update_by;not null" json:"update_by"`        // 修改人
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}
