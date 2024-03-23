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

// TAndroidCert Android证书信息表
type TAndroidCert struct {
	Id           uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	CertId       string    `gorm:"column:cert_id; type:char(32) not null; comment:唯一标识; uniqueIndex:idx_cert_id" json:"certId,omitempty"`
	AppId        uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId       uint      `gorm:"column:user_id; type:int unsigned; comment:创建人，外码" json:"userId,omitempty"`
	Organization string    `gorm:"column:organization; type:varchar(64); comment:组织名" json:"organization,omitempty"`
	Alias        string    `gorm:"column:alias; type:varchar(64); comment:别名" json:"alias,omitempty"`
	Publisher    string    `gorm:"column:publisher; type:varchar(1024); comment:发布者" json:"publisher,omitempty"`
	Owner        string    `gorm:"column:owner; type:varchar(1024); comment:所有者" json:"owner,omitempty"`
	Algorithm    string    `gorm:"column:algorithm; type:varchar(64); comment:证书算法" json:"algorithm,omitempty"`
	Serial       string    `gorm:"column:serial; type:char(8); comment:序列号" json:"serial,omitempty"`
	Content      []byte    `gorm:"column:content; type:blob; comment:证书加密内容" json:"content,omitempty"`
	KeyId        uint      `gorm:"column:key_id; type:int unsigned; comment:证书加密密钥标识，外码" json:"keyId,omitempty"`
	MD5          string    `gorm:"column:md5; type:char(32); comment:md5值" json:"md5,omitempty"`
	SHA1         string    `gorm:"column:sha1; type:char(40); comment:sha1值" json:"sha1,omitempty"`
	SHA256       string    `gorm:"column:sha256; type:char(64); comment:sha256值" json:"sha256,omitempty"`
	Storepass    string    `gorm:"column:storepass; type:varchar(128); comment:密钥库密码" json:"storepass,omitempty"`
	Keypass      string    `gorm:"column:keypass; type:varchar(128); comment:证书密码" json:"keypass,omitempty"`
	EffectTime   time.Time `gorm:"column:effect_time; type:timestamp; comment:生效时间" json:"effectTime,omitempty"`
	ExpireTime   time.Time `gorm:"column:expire_time; type:timestamp; comment:失效时间" json:"expireTime,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
	DeleteTime   time.Time `gorm:"column:delete_time; type:timestamp; comment:删除标识" json:"deleteTime,omitempty"`
}

func (*TAndroidCert) TableName() string {
	return "t_android_cert"
}
