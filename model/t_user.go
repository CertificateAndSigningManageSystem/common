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

import "time"

const (
	TUser_Status_OK = 1 + iota
	TUser_Status_Locked
)

// TUser 用户信息表
type TUser struct {
	Id            uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	NameEn        string    `gorm:"column:name_en; type:varchar(16) not null unique; comment:英文名; uniqueIndex:idx_name_en"`
	NameZh        string    `gorm:"column:name_zh; type:varchar(32); comment:中文名"`
	Avatar        string    `gorm:"column:avatar; type:char(38); comment:头像文件id，外码"`
	PasswdDigest  string    `gorm:"column:passwd_digest; type:char(32); comment:密码摘要"`
	PasswdSalt    string    `gorm:"column:passwd_salt; type:char(128); comment:密码盐值"`
	RegisterTime  time.Time `gorm:"column:register_time; type:timestamp; comment:注册时间"`
	LastLoginTime time.Time `gorm:"column:last_login_time; type:timestamp; comment:最后登录时间"`
	Status        uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=正常,2=冻结"`
}

func (*TUser) TableName() string {
	return "t_user"
}
