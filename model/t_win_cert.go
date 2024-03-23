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
	TWinCert_Type_OV uint8 = 1 + iota
	TWinCert_Type_EV
)

// TWinCert Windows证书信息表
type TWinCert struct {
	Id           uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	CertId       string    `gorm:"column:cert_id; type:char(32) not null; comment:唯一标识，随机生成; uniqueIndex:idx_cert_id" json:"certId,omitempty"`
	AppId        uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId       uint      `gorm:"column:user_id; type:int unsigned; comment:创建人，外码" json:"userId,omitempty"`
	Fingerprint  string    `gorm:"column:fingerprint; type:char(40); comment:证书指纹" json:"fingerprint,omitempty"`
	Organization string    `gorm:"column:organization; type:varchar(64); comment:组织名" json:"organization,omitempty"`
	Version      string    `gorm:"column:version; type:varchar(16); comment:版本号" json:"version,omitempty"`
	Algorithm    string    `gorm:"column:algorithm; type:varchar(64); comment:证书算法" json:"algorithm,omitempty"`
	Publisher    string    `gorm:"column:publisher; type:varchar(1024); comment:发布者" json:"publisher,omitempty"`
	Serial       string    `gorm:"column:serial; type:char(48); comment:序列号" json:"serial,omitempty"`
	Owner        string    `gorm:"column:owner; type:varchar(1024); comment:所有者" json:"owner,omitempty"`
	Password     string    `gorm:"column:password; type:varchar(64); comment:密码" json:"password,omitempty"`
	Content      []byte    `gorm:"column:content; type:blob; comment:证书加密内容" json:"content,omitempty"`
	KeyId        uint      `gorm:"column:key_id; type:int unsigned; comment:证书加密密钥标识，外码" json:"keyId,omitempty"`
	Type         uint8     `gorm:"column:type; type:int unsigned; comment:证书类型，1=公共证书OV,2=公共证书EV" json:"type,omitempty"`
	EffectTime   time.Time `gorm:"column:effect_time; type:timestamp; comment:生效时间" json:"effectTime,omitempty"`
	ExpireTime   time.Time `gorm:"column:expire_time; type:timestamp; comment:失效时间" json:"expireTime,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
	DeleteTime   time.Time `gorm:"column:delete_time; type:timestamp; comment:删除标识" json:"deleteTime,omitempty"`
}

func (*TWinCert) TableName() string {
	return "t_win_cert"
}
