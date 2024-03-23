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

// TAppleDevice Apple测试设备信息表
type TAppleDevice struct {
	Id            uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	UDID          string    `gorm:"column:udid; type:char(40) not null; comment:设备唯一标识符; uniqueIndex:idx_udid" json:"udid,omitempty"`
	AppId         uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId        uint      `gorm:"column:user_id; type:int unsigned; comment:绑定人，外码" json:"userId,omitempty"`
	Name          string    `gorm:"column:name; type:varchar(128); comment:设备名称" json:"name,omitempty"`
	Model         string    `gorm:"column:model; type:varchar(128); comment:设备类型" json:"model,omitempty"`
	AppleDeviceId string    `gorm:"column:apple_device_id; type:char(8); comment:Apple绑定Id" json:"appleDeviceId,omitempty"`
	BindTime      time.Time `gorm:"column:bind_time; type:timestamp; comment:绑定时间" json:"bindTime,omitempty"`
}

func (*TAppleDevice) TableName() string {
	return "t_apple_device"
}
