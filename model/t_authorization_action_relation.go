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

// TAuthorizationActionRelation 凭证授权表
type TAuthorizationActionRelation struct {
	Id       uint `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	AuthId   uint `gorm:"column:auth_id; type:int unsigned not null; comment:凭证Id，外码; index:idx_auth_id"`
	ActionId uint `gorm:"column:action_id; type:int unsigned not null; comment:授权项Id，外码; index:idx_action_id"`
}

func (*TAuthorizationActionRelation) TableName() string {
	return "t_authorization_action_relation"
}
