// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCasbinRule = "casbin_rule"

// CasbinRule mapped from table <casbin_rule>
type CasbinRule struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Ptype      string         `gorm:"column:ptype" json:"ptype"`
	V0         string         `gorm:"column:v0" json:"v0"`
	V1         string         `gorm:"column:v1" json:"v1"`
	V2         string         `gorm:"column:v2" json:"v2"`
	V3         string         `gorm:"column:v3" json:"v3"`
	V4         string         `gorm:"column:v4" json:"v4"`
	V5         string         `gorm:"column:v5" json:"v5"`
	Desc       string         `gorm:"column:desc;not null;comment:描述" json:"desc"`
	CreateAt   time.Time      `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP;comment:记录创建时间" json:"create_at"`
	UpdateAt   time.Time      `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;comment:记录修改时间" json:"update_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreateUser string         `gorm:"column:create_user;not null;comment:创建人" json:"create_user"`
	UpdateUser string         `gorm:"column:update_user;not null;comment:修改人" json:"update_user"`
}

// TableName CasbinRule's table name
func (*CasbinRule) TableName() string {
	return TableNameCasbinRule
}
