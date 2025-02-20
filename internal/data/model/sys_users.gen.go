// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUser = "sys_users"

// SysUser mapped from table <sys_users>
type SysUser struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	UUID      string         `gorm:"column:uuid;not null" json:"uuid"`                  // 用户UUID
	Username  string         `gorm:"column:username;not null" json:"username"`          // 用户名(登入)
	NickName  string         `gorm:"column:nick_name;not null" json:"nick_name"`        // 昵称
	Password  string         `gorm:"column:password;not null" json:"password"`          // 密码
	Phone     string         `gorm:"column:phone;not null" json:"phone"`                // 手机
	RoleID    int64          `gorm:"column:role_id;not null" json:"role_id"`            // 角色id
	Salt      string         `gorm:"column:salt;not null" json:"salt"`                  // 盐
	Avatar    string         `gorm:"column:avatar;not null" json:"avatar"`              // 头像
	Sex       int32          `gorm:"column:sex;not null" json:"sex"`                    // 性别 0-未知 1-男 2-女
	Email     string         `gorm:"column:email;not null" json:"email"`                // 邮箱
	DeptID    int64          `gorm:"column:dept_id;not null" json:"dept_id"`            // 部门id
	PostID    int64          `gorm:"column:post_id;not null" json:"post_id"`            // 岗位id
	Remark    string         `gorm:"column:remark;not null" json:"remark"`              // 备注
	Status    int32          `gorm:"column:status;not null;default:1" json:"status"`    // 1=正常 2=异常
	RoleIds   string         `gorm:"column:role_ids;not null" json:"role_ids"`          // 多角色
	PostIds   string         `gorm:"column:post_ids;not null" json:"post_ids"`          // 多岗位
	CreateBy  string         `gorm:"column:create_by;not null" json:"create_by"`        // 创建人
	UpdateBy  string         `gorm:"column:update_by;not null" json:"update_by"`        // 更新人
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`               // 删除时间
	Secret    string         `gorm:"column:secret;not null" json:"secret"`              // google密钥
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
