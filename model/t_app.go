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
	TApp_Platform_Windows uint8 = 1 + iota
	TApp_Platform_Android
	TApp_Platform_Apple
)

const (
	TApp_Status_OK uint8 = 1 + iota
	TApp_Status_Locked
)

// TApp 应用信息表
type TApp struct {
	Id             uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	AppId          string    `gorm:"column:app_id; type:char(32) not null; comment:应用Id; uniqueIndex:idx_app_id" json:"appId,omitempty"`
	Name           string    `gorm:"column:name; type:varchar(64); comment:名称" json:"name,omitempty"`
	UserId         uint      `gorm:"column:user_id; type:int unsigned; comment:创建人，外码; index:idx_user_id" json:"userId,omitempty"`
	Avatar         string    `gorm:"column:avatar; type:char(38); comment:头像文件Id" json:"avatar,omitempty"`
	Platform       uint8     `gorm:"column:platform; type:tinyint unsigned; comment:平台，1=Windows,2=Android,3=Apple" json:"platform,omitempty"`
	CreateTime     time.Time `gorm:"column:create_time; type:timestamp not null; comment:创建时间; index:idx_create_time" json:"createTime,omitempty"`
	AppleAccountId uint      `gorm:"column:apple_account_id; type:int unsigned; comment:Apple账号Id，外码" json:"appleAccountId,omitempty"`
	Status         uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=正常,2=注销" json:"status,omitempty"`
}

func (*TApp) TableName() string {
	return "t_app"
}
