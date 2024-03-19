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
	TAndroidSignJob_TaskType_Apk uint8 = 1 + iota
	TAndroidSignJob_TaskType_Aab
	TAndroidSignJob_TaskType_Patch
)

const (
	TAndroidSignJob_Status_Init uint8 = 1 + iota
	TAndroidSignJob_Status_Success
	TAndroidSignJob_Status_Fail
)

const (
	TAndroidSignJob_Source_Web uint8 = 1 + iota
	TAndroidSignJob_Source_API
)

// TAndroidSignJob Android签名任务表
type TAndroidSignJob struct {
	Id           uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	JobId        string    `gorm:"column:job_id; type:char(38) not null; comment:任务Id; uniqueIndex:idx_job_id" json:"jobId,omitempty"`
	AppId        uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId       uint      `gorm:"column:user_id; type:int unsigned; comment:用户Id或凭证Id，外码" json:"userId,omitempty"`
	CertId       uint      `gorm:"column:cert_id; type:int unsigned; comment:证书Id，外码" json:"certId,omitempty"`
	FileId       string    `gorm:"column:file_id; type:char(38); comment:文件Id" json:"fileId,omitempty"`
	SignedFileId string    `gorm:"column:signed_file_id; type:char(38); comment:已签名文件Id" json:"signedFileId,omitempty"`
	TaskType     uint8     `gorm:"column:task_type; type:tinyint unsigned; comment:任务类型，1=apk,2=aab,3=patch" json:"taskType,omitempty"`
	Log          string    `gorm:"column:log; type:text; comment:签名输出日志" json:"log,omitempty"`
	SignSchema   string    `gorm:"column:sign_schema; type:varchar(32); comment:apk签名方案，逗号分割" json:"signSchema,omitempty"`
	MinSDKLevel  uint8     `gorm:"column:min_sdk_level; type:tinyint unsigned; comment:AndroidAPI最低等级" json:"minSDKLevel,omitempty"`
	Source       uint8     `gorm:"column:source; type:tinyint unsigned; comment:来源，1=web,2=api" json:"source,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time; type:timestamp not null; comment:创建时间; index:idx_create_time" json:"createTime,omitempty"`
	FinishTime   time.Time `gorm:"column:finish_time; type:timestamp; comment:结束时间" json:"finishTime,omitempty"`
	Status       uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=待签名,2=成功，3=失败" json:"status,omitempty"`
}

func (*TAndroidSignJob) TableName() string {
	return time.Now().Format("t_android_sign_job_200601")
}
