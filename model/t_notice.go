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

// TNotice 系统公告表
type TNotice struct {
	Id           uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty" json:"id,omitempty"`
	Content      string    `gorm:"column:content; type:text; comment:公告内容" json:"content,omitempty"`
	ActiveTime   time.Time `gorm:"column:active_time; type:timestamp; comment:生效时间" json:"activeTime,omitempty"`
	InactiveTime time.Time `gorm:"column:inactive_time; type:timestamp; comment:失效时间" json:"inactiveTime,omitempty"`
	Creator      uint      `gorm:"column:creator; type:int unsigned; comment:创建人" json:"creator,omitempty"`
	CreateTime   time.Time `gorm:"column:create_time; type:timestamp; comment:创建时间" json:"createTime,omitempty"`
}

func (*TNotice) TableName() string {
	return "t_notice"
}
