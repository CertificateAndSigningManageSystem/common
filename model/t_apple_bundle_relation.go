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

// TAppleBundleRelation AppleBundle能力关联表
type TAppleBundleRelation struct {
	Id         uint `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	Capability uint `gorm:"column:capability_id; type:int unsigned; comment:能力项Id，外码" json:"capability,omitempty"`
	BundleId   uint `gorm:"column:bundle_id; type:int unsigned not null; comment:bundle Id，外码; index:idx_bundle_id" json:"bundleId,omitempty"`
}

func (*TAppleBundleRelation) TableName() string {
	return "t_apple_bundle_relation"
}
