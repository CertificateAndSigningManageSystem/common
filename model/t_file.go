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

// TFile 文件信息表
type TFile struct {
	Id         uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	FileId     string    `gorm:"column:file_id; type:char(38) not null; comment:文件Id，前6为年月，后32为是UUID; uniqueIndex:idx_file_id" json:"fileId,omitempty"`
	UserId     uint      `gorm:"column:user_id; type:int unsigned; comment:上传人，外码" json:"userId,omitempty"`
	TusdId     string    `gorm:"column:tusd_id; type:char(32); comment:tusd文件Id" json:"tusdId,omitempty"`
	Name       string    `gorm:"column:name; type:varchar(256); comment:文件名" json:"name,omitempty"`
	Ext        string    `gorm:"column:ext; type:varchar(16); comment:格式" json:"ext,omitempty"`
	MD5        string    `gorm:"column:md5; type:char(32); comment:散列值" json:"md5,omitempty"`
	SHA1       string    `gorm:"column:sha1; type:char(40); comment:散列值" json:"sha1,omitempty"`
	SHA256     string    `gorm:"column:sha256; type:char(64); comment:散列值" json:"sha256,omitempty"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp not null; comment:上传时间; index:idx_create_time" json:"createTime,omitempty"`
}

func (*TFile) TableName() string {
	return time.Now().Format("t_file_200601")
}
