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

// TAppleBundle AppleBundle信息表
type TAppleBundle struct {
	Id              uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	AppId           uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId          uint      `gorm:"column:user_id; type:int unsigned; comment:创建人，外码" json:"userId,omitempty"`
	AppleBundleName string    `gorm:"column:apple_bundle_name; type:varchar(256); comment:Bundle名" json:"appleBundleName,omitempty"`
	AppleBundleId   string    `gorm:"column:apple_bundle_id; type:char(10); comment:BundleId" json:"appleBundleId,omitempty"`
	Platform        uint8     `gorm:"column:platform; type:tinyint unsigned; comment:平台，1=UNIVERSAL,2=IOS,3=MAC_OS" json:"platform,omitempty"`
	CreateTime      time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
}

func (*TAppleBundle) TableName() string {
	return "t_apple_bundle"
}
