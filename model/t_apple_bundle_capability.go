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

// TAppleBundleCapability AppleBundle能力表
type TAppleBundleCapability struct {
	Id       uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	Name     string `gorm:"column:name; type:varchar(64); comment:能力名" json:"name,omitempty"`
	Platform uint8  `gorm:"column:platform; type:tinyint unsigned; comment:平台，1=IOS,2=MAC_OS,3=UNIVERSAL" json:"platform,omitempty"`
}

func (*TAppleBundleCapability) TableName() string {
	return "t_apple_bundle_capability"
}
