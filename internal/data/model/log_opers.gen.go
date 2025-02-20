// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLogOper = "log_opers"

// LogOper mapped from table <log_opers>
type LogOper struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键id
	Title        string         `gorm:"column:title;not null" json:"title"`                 // 操作的模块
	BusinessType int32          `gorm:"column:business_type;not null" json:"business_type"` // 0其它 1新增 2修改 3删除
	URL          string         `gorm:"column:url;not null" json:"url"`                     // 操作url
	Method       string         `gorm:"column:method;not null" json:"method"`               // 请求方法
	UserName     string         `gorm:"column:user_name;not null" json:"user_name"`         // 操作人员
	UserID       int64          `gorm:"column:user_id;not null" json:"user_id"`             // 用户id
	IP           string         `gorm:"column:ip;not null" json:"ip"`                       // 操作IP
	Agent        string         `gorm:"column:agent;not null" json:"agent"`                 // 代理
	Latency      int32          `gorm:"column:latency;not null" json:"latency"`             // 延迟
	Resp         string         `gorm:"column:resp;not null" json:"resp"`                   // 请求参数
	Status       int32          `gorm:"column:status;not null;default:1" json:"status"`     // 1=正常 2=异常
	ErrorMessage string         `gorm:"column:error_message;not null" json:"error_message"` // 错误信息
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`                // 创建时间
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`                // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                // 删除时间
}

// TableName LogOper's table name
func (*LogOper) TableName() string {
	return TableNameLogOper
}
