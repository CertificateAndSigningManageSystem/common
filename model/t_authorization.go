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

// TAuthorization 凭证信息表
type TAuthorization struct {
	Id         uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	AppId      uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	AuthId     string    `gorm:"column:auth_id; type:varchar(64); comment:凭证Id" json:"authId,omitempty"`
	UserId     uint      `gorm:"column:user_id; type:int unsigned; comment:用户Id，外码" json:"userId,omitempty"`
	IP         string    `gorm:"column:ip; type:varchar(1024); comment:调用IP" json:"ip,omitempty"`
	Frequency  uint      `gorm:"column:frequency; type:int unsigned; comment:每分钟调用最高频率" json:"frequency,omitempty"`
	Secret     string    `gorm:"column:secret; type:char(128); comment:密钥" json:"secret,omitempty"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
	ExpireTime time.Time `gorm:"column:expire_time; type:timestamp; comment:失效时间" json:"expireTime,omitempty"`
}

func (*TAuthorization) TableName() string {
	return "t_authorization"
}
