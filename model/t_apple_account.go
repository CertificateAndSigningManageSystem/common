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
	TAppleAccount_Type_Personal uint8 = 1 + iota
	TAppleAccount_Type_Company
	TAppleAccount_Type_Enterprise
)

// TAppleAccount 苹果账号信息表
type TAppleAccount struct {
	Id      uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	Account string `gorm:"column:account; type:varchar(128); comment:账号" json:"account,omitempty"`
	AppleId string `gorm:"column:team_id; type:char(10) not null; comment:账号Id; uniqueIndex:idx_apple_id"`
	Name    string `gorm:"column:name; type:varchar(128); comment:名称" json:"name,omitempty"`
	Issuer  string `gorm:"column:issuer; type:char(36); comment:Apple凭证iss" json:"issuer,omitempty"`
	Kid     string `gorm:"column:kid; type:char(10); comment:Apple凭证kid" json:"kid,omitempty"`
	Key     string `gorm:"column:key; type:varchar(512); comment:pkcs#8格式的凭证密钥" json:"key,omitempty"`
	Admin   string `gorm:"column:admin; type:varchar(32); comment:账号管理员" json:"admin,omitempty"`
	Type    uint   `gorm:"column:type; type:int unsigned; comment:类型，1=个人,2=公司,3=企业" json:"type,omitempty"`
}

func (*TAppleAccount) TableName() string {
	return "t_apple_account"
}
