/*
 * Copyright (c) 2023 ivfzhou
 * common is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package model

const (
	// 超管
	TUserRole_Role_SuperAdmin uint8 = 1 + iota
	// 系统管理员
	TUserRole_Role_Admin
	// 应用管理员
	TUserRole_Role_AppAdmin
	// 应用成员
	TUserRole_Role_AppMember
	// 可使用签名服务
	TUserRole_Role_AppSigner
)

// TUserRole 用户角色表
type TUserRole struct {
	ID     uint  `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	UserID uint  `gorm:"column:user_id; type:int unsigned not null; comment:用户ID，外码" json:"userId,omitempty"`
	AppID  uint  `gorm:"column:app_id; type:int unsigned not null; comment:应用ID，外码" json:"appId,omitempty"`
	Role   uint8 `gorm:"column:role; type:tinyint unsigned not null; comment:角色，1=超管,2=系统管理员,3=应用管理员,4=应用成员,5=可使用签名服务" json:"role,omitempty"`
}

func (*TUserRole) TableName() string {
	return "t_user_role"
}
