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

// TAppleProfile Apple描述文件信息表
type TAppleProfile struct {
	Id            uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	CertId        string `gorm:"column:cert_id; type:char(32) not null; comment:唯一标识，随机生成; uniqueIndex:idx_cert_id" json:"certId,omitempty"`
	UUID          string `gorm:"column:uuid; type:char(32); comment:描述文件标识" json:"uuid,omitempty"`
	AppId         uint   `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId        uint   `gorm:"column:user_id; type:int unsigned; comment:创建人，外码" json:"userId,omitempty"`
	Content       []byte `gorm:"column:content; type:blob; comment:证书加密后的内容" json:"content,omitempty"`
	KeyId         uint   `gorm:"column:key_id; type:int unsigned; comment:加密证书的密钥Id" json:"keyId,omitempty"`
	AppleId       string `gorm:"column:apple_id; type:char(10); comment:Apple账号Id" json:"appleId,omitempty"`
	AppleBundleId string `gorm:"column:apple_bundle_id; type:char(10); comment:绑定的BundleId" json:"appleBundleId,omitempty"`
	AppleCertId   string `gorm:"column:apple_cert_id; type:char(10); comment:对应签名证书Id" json:"appleCertId,omitempty"`
	Text          string `gorm:"column:text; type:text; comment:描述文件中xml文本" json:"text,omitempty"`
	Category      uint8  `gorm:"column:category tinyint unsigned; comment:类型，
		1=IOS_APP_DEVELOPMENT,2=IOS_APP_STORE,3=IOS_APP_ADHOC,
		4=IOS_APP_INHOUSE,5=MAC_APP_DEVELOPMENT,6=MAC_APP_STORE,
		7=MAC_APP_DIRECT,8=TVOS_APP_DEVELOPMENT,9=TVOS_APP_STORE,
		10=TVOS_APP_ADHOC,11=TVOS_APP_INHOUSE,12=MAC_CATALYST_APP_DEVELOPMENT,
		13=MAC_CATALYST_APP_STORE,14=MAC_CATALYST_APP_DIRECT" json:"category,omitempty"`
	Platform    uint8 `gorm:"column:platform; type:tinyint unsigned; comment:Apple证书平台，1=IOS,2=MAC_OS" json:"platform,omitempty"`
	ProfileType uint8 `gorm:"column:profile_type; type:tinyint unsigned; comment:类型，
		1=IOS_APP_ADHOC,2=IOS_APP_STORE,3=MAC_APP_STORE,
		4=IOS_APP_DEVELOPMENT,5=MAC_APP_DEVELOPMENT" json:"profileType,omitempty"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
	ExpireTime time.Time `gorm:"column:expire_time; type:timestamp; comment:过期时间" json:"expireTime,omitempty"`
	DeleteTime time.Time `gorm:"column:delete_time; type:timestamp; comment:删除标识" json:"deleteTime,omitempty"`
}

func (*TAppleProfile) TableName() string {
	return "t_apple_profile"
}
