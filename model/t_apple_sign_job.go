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

// TAppleSignJob Apple签名任务表
type TAppleSignJob struct {
	Id                  uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	JobId               string    `gorm:"column:job_id; type:char(38) not null; comment:任务Id; uniqueIndex:idx_job_id" json:"jobId,omitempty"`
	AppId               uint      `gorm:"column:app_id; type:int unsigned not null; comment:应用Id，外码; index:idx_app_id" json:"appId,omitempty"`
	UserId              uint      `gorm:"column:user_id; type:int unsigned; comment; 用户Id或凭证Id，外码" json:"userId,omitempty"`
	FileId              string    `gorm:"column:file_id; type:char(38); comment:文件Id" json:"fileId,omitempty"`
	TaskType            uint8     `gorm:"column:task_type; type:tinyint unsigned; comment:任务类型，1=ipa,2=pkg,3=dmg" json:"taskType,omitempty"`
	CertId              uint      `gorm:"column:cert_id; type:int unsigned; comment:证书Id，外码" json:"certId,omitempty"`
	Log                 string    `gorm:"column:log; type:text; comment:签名输出日志" json:"log,omitempty"`
	SignedFileId        string    `gorm:"column:signed_file_id; type:char(38); comment:已签名文件Id" json:"signedFileId,omitempty"`
	ULDomains           string    `gorm:"column:ul_domains; type:varchar(256); comment:UL域名，多个逗号分割" json:"ulDomains,omitempty"`
	KeychainAccessGroup string    `gorm:"column:keychain-access-group; type:varchar(256); comment:钥匙串访问组，多个分号分割" json:"keychainAccessGroup,omitempty"`
	AppexInfo           string    `gorm:"column:appex_info; type:varchar(1024); comment:appex信息，JSON对象数组的形式存储" json:"appexInfo,omitempty"`
	Source              uint8     `gorm:"column:source; type:tinyint unsigned; comment:来源，1=web,2=api" json:"source,omitempty"`
	CreateTime          time.Time `gorm:"column:create_time; type:timestamp not null; comment:创建时间; index:idx_create_time" json:"createTime,omitempty"`
	FinishTime          time.Time `gorm:"column:finish_time; type:timestamp; comment:结束时间" json:"finishTime,omitempty"`
	Status              uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=待签名,2=成功,3=失败" json:"status,omitempty"`
}

func (t *TAppleSignJob) TableName() string {
	if !t.CreateTime.IsZero() {
		return t.CreateTime.Format("t_apple_sign_job_200601")
	}
	return time.Now().Format("t_apple_sign_job_200601")
}
