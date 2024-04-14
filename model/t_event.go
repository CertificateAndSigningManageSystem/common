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
	TEvent_Type_Register uint8 = 1 + iota
	TEvent_Type_Login
	TEvent_Type_LockAccount
	TEvent_Type_UnlockAccount
	TEvent_Type_ChangePasswd
	TEvent_Type_ResetPasswd
	TEvent_Type_ModifyUserInfo
	TEvent_Type_CreateApp
	TEvent_Type_UpdateApp
	TEvent_Type_LockApp
	TEvent_Type_ApplyOpenAPIToken
	TEvent_Type_UpdateOpenAPIToken
	TEvent_Type_RenewalOpenAPIToken
	TEvent_Type_DeleteOpenAPIToken
	TEvent_Type_ApplyAndroidCert
	TEvent_Type_ApplyAppleProfile
	TEvent_Type_DeleteCert
	TEvent_Type_ResetLoginFailTimes
	TEvent_Type_ApplyPushCert
	TEvent_Type_RegisterAppleDevice
	TEvent_Type_UnregisterAppleDevice
	TEvent_Type_BindAppleAccount
	TEvent_Type_ExportAppInfo
	TEvent_Type_DownloadGooglePlayCert
	TEvent_Type_ObtainFacebookDigest
)

// TEvent 事件表
type TEvent struct {
	Id   uint  `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	Type uint8 `gorm:"column:type; type:tinyint unsigned; comment:类型，
		1=注册,2=登陆,3=冻结用户,4=解冻用户,5=密码修改,6=密码重置,
		7=个人信息修改,8=应用创建,9=应用更新,10=应用注销,
		11=申请OpenAPI凭证,12=更新OpenAPI凭证,13=续期OpenAPI凭证,
		14=删除OpenAPI凭证,15=申请Android证书,16=申请描述文件,17=删除证书,
		18=重置登录连续失败次数,19=申请Push证书,20=注册测试设备,21=解绑测试设备,
		22=绑定Apple账号,23=导出应用数据,24=GooglePlay证书下载,25=获取证书Facebook散列" json:"type,omitempty"`
	OccurTime time.Time `gorm:"column:occur_time; type:timestamp not null; comment:发生时间; index:idx_occur_time" json:"occurTime,omitempty"`
	UserId    uint      `gorm:"column:user_id; type:int unsigned not null; comment:关联人，外码; index:idx_user_id" json:"userId,omitempty"`
	AppId     uint      `gorm:"column:app_id; type:int unsigned; comment:关联应用，外码; index:id_app_id" json:"appId,omitempty"`
	Content   string    `gorm:"column:content; type:text; comment:内容"`
}

func (t *TEvent) TableName() string {
	if !t.OccurTime.IsZero() {
		return t.OccurTime.Format("t_event_200601")
	}
	return time.Now().Format("t_event_200601")
}
