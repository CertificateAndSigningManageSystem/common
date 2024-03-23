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

// TAppleCert Apple证书信息表
type TAppleCert struct {
	Id           uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	CertId       string `gorm:"column:cert_id; type:char(32) not null; comment:唯一标识，随机生成; uniqueIndex:idx_cert_id" json:"certId,omitempty"`
	AppleCertId  string `gorm:"column:apple_cert_id; type:char(10); comment:证书Id" json:"appleCertId,omitempty"`
	AppId        uint   `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId       uint   `gorm:"column:user_id; type:int unsigned; comment:创建人，外码" json:"userId,omitempty"`
	AppleId      string `gorm:"column:apple_id; type:char(10); comment:Apple账号Id" json:"appleId,omitempty"`
	Organization string `gorm:"column:organization; type:varchar(64); comment:组织名" json:"organization,omitempty"`
	Content      []byte `gorm:"column:content; type:blob; comment:证书加密后的内容" json:"content,omitempty"`
	KeyId        uint   `gorm:"column:key_id; type:int unsigned; comment:加密证书的密钥Id" json:"keyId,omitempty"`
	Category     uint8  `gorm:"column:category; type:tinyint unsigned; comment:类型，
		1=DISTRIBUTION,2=DEVELOPMENT,3=IOS_DISTRIBUTION,
		4=IOS_DEVELOPMENT,5=MAC_APP_DISTRIBUTION,
		6=MAC_INSTALL_DISTRIBUTION,7=MAC_APP_DEVELOPMENT,
		8=DEVELOPER_ID_KEXT,9=DEVELOPER_ID_APPLICATION,
		10=DEVELOPER_ID_INSTALL" json:"category,omitempty"`
	SerialNumber string    `gorm:"column:serial_number; type:char(32); comment:序列号" json:"serialNumber,omitempty"`
	Platform     uint8     `gorm:"column:platform; type:tinyint unsigned; comment:Apple证书平台，1=IOS,2=MAC_OS" json:"platform,omitempty"`
	Passwd       string    `gorm:"column:passwd; type:varchar(64); comment:证书密码" json:"passwd,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
	ExpireTim    time.Time `gorm:"column:expire_time; type:timestamp; comment:过期时间" json:"expireTim,omitempty"`
	DeleteTime   time.Time `gorm:"column:delete_time; type:timestamp; comment:删除标识" json:"deleteTime,omitempty"`
}

func (*TAppleCert) TableName() string {
	return "t_apple_cert"
}
