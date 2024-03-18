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
	// 获取应用信息
	TAuthorizationAction_ID_GetAppInfo uint = 1 + iota
	// 获取签名信息信息
	TAuthorizationAction_ID_GetSignJobInfo
	// 使用签名服务
	TAuthorizationAction_ID_SignJob
)

// TAuthorizationAction 凭证授权项表
type TAuthorizationAction struct {
	Id       uint   `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	Name     string `gorm:"column:name; type:varchar(64); comment:名称" json:"name,omitempty"`
	Platform uint8  `gorm:"column:platform; type:tinyint unsigned; comment:适用应用平台，1=Windows,2=Android,3=Apple,4=所有" json:"platform,omitempty"`
}

func (*TAuthorizationAction) TableName() string {
	return "t_authorization_action"
}
