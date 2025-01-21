// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCloudregionsTbl = "cloudregions_tbl"

// CloudregionsTbl mapped from table <cloudregions_tbl>
type CloudregionsTbl struct {
	CreatedAt     time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	UpdateVersion int32          `gorm:"column:update_version;not null" json:"update_version"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Deleted       bool           `gorm:"column:deleted;not null" json:"deleted"`
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	Description   string         `gorm:"column:description" json:"description"`
	IsEmulated    bool           `gorm:"column:is_emulated;not null" json:"is_emulated"`
	Name          string         `gorm:"column:name;not null" json:"name"`
	Status        string         `gorm:"column:status;not null;default:init" json:"status"`
	Progress      float32        `gorm:"column:progress;default:100" json:"progress"`
	Enabled       bool           `gorm:"column:enabled" json:"enabled"`
	ExternalID    string         `gorm:"column:external_id" json:"external_id"`
	ImportedAt    time.Time      `gorm:"column:imported_at" json:"imported_at"`
	Source        string         `gorm:"column:source" json:"source"`
	Latitude      float32        `gorm:"column:latitude" json:"latitude"`
	Longitude     float32        `gorm:"column:longitude" json:"longitude"`
	City          string         `gorm:"column:city" json:"city"`
	CountryCode   string         `gorm:"column:country_code" json:"country_code"`
	Environment   string         `gorm:"column:environment" json:"environment"`
	Provider      string         `gorm:"column:provider;not null;default:OneCloud" json:"provider"`
}

// TableName CloudregionsTbl's table name
func (*CloudregionsTbl) TableName() string {
	return TableNameCloudregionsTbl
}
