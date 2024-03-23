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

// TAndroidCertOrg Android证书主体表
type TAndroidCertOrg struct {
	Id         uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	Name       string    `gorm:"column:name; type:varchar(256); comment:主体名" json:"name,omitempty"`
	UserId     uint      `gorm:"column:user_id; type:int unsigned; comment:创建人" json:"userId,omitempty"`
	Dname      string    `gorm:"column:dname; type:varchar(1024); comment:组织信息，keytools的dname参数值" json:"dname,omitempty"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
}

func (*TAndroidCertOrg) TableName() string {
	return "t_android_cert_org"
}
