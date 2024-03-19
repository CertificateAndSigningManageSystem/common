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
	TWinSignJob_TaskType_PE uint8 = 1 + iota
	TWinSignJob_TaskType_HLKAndWHQL
	TWinSignJob_TaskType_WHQL
	TWinSignJob_TaskType_Sys
)

const (
	TWinSignJob_Source_Web uint8 = 1 + iota
	TWinSignJob_Source_API
)

const (
	TWinSignJob_Status_Init uint8 = 1 + iota
	TWinSignJob_Status_Success
	TWinSignJob_Status_Fail
	TWinSignJob_Status_NeedHLKTest
	TWinSignJob_Status_NeedWHQL
)

// TWinSignJob Windows签名任务表
type TWinSignJob struct {
	Id               uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	JobId            string    `gorm:"column:job_id; type:char(38) not null; comment:任务Id; uniqueIndex:idx_job_id" json:"jobId,omitempty"`
	AppId            uint      `gorm:"column:app_id; type:int unsigned not null; comment:用户Id或凭证Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId           uint      `gorm:"column:user_id; type:int unsigned; comment:用户Id或凭证Id，外码" json:"userId,omitempty"`
	FileId           string    `gorm:"column:file_id; type:char(38); comment:文件Id" json:"fileId,omitempty"`
	TaskType         uint8     `gorm:"column:task_type; type:tinyint unsigned; comment:任务类型，1=pe,2=hlk&whql,3=whql,4=sys" json:"taskType,omitempty"`
	CertOrganization string    `gorm:"column:cert_organization; type:varchar(128); comment:证书组织" json:"certOrganization,omitempty"`
	Log              string    `gorm:"column:log; type:text; comment:签名输出日志" json:"log,omitempty"`
	HLKTestSystem    string    `gorm:"column:hlk_test_system; type:varchar(64); comment:测试系统版本" json:"hlkTestSystem,omitempty"`
	TmpFileId        string    `gorm:"column:tmp_file_id; type:char(38); comment:hlk测试成功后打包文件或者sys签名文件Id" json:"tmpFileId,omitempty"`
	SignedFileId     string    `gorm:"column:signed_file_id; type:char(38); comment:已签名文件Id" json:"signedFileId,omitempty"`
	Source           uint8     `gorm:"column:source; type:tinyint unsigned; comment:来源，1=web,2=api" json:"source,omitempty"`
	CreateTime       time.Time `gorm:"column:create_time; type:timestamp not null; comment:创建时间" json:"createTime,omitempty"`
	FinishTime       time.Time `gorm:"column:finish_time; type:timestamp; comment:结束时间" json:"finishTime,omitempty"`
	Status           uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=待签名,2=成功,3=失败,4=待hlk测试,5=待微软处理" json:"status,omitempty"`
}

func (*TWinSignJob) TableName() string {
	return time.Now().Format("t_win_sign_job_200601")
}
