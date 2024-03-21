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
	TTodo_Type_CertExpire uint8 = 1 + iota
	TTodo_Type_ApplyJoinApp
	TTodo_Type_ApplyAppSigning
)

const (
	TTodo_Status_Init uint8 = 1 + iota
	TTodo_Status_Done
	TTodo_Status_Reject
	TTodo_Status_Agree
)

// TTodo 待办信息表
type TTodo struct {
	Id         uint      `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	TodoId     string    `gorm:"column:todo_id; type:char(32) not null; comment:待办Id; uniqueIndex:idx_todo_id" json:"todoId,omitempty"`
	UserId     uint      `gorm:"column:user_id; type:int unsigned not null; comment:申请人; index:idx_user_id" json:"userId,omitempty"`
	Type       uint8     `gorm:"column:type; type:tinyint unsigned; comment:待办类型，1=证书过期,2=申请加入应用,3=申请签名权限" json:"type,omitempty"`
	Status     uint8     `gorm:"column:status; type:tinyint unsigned; comment:状态，1=未处理,2=已处理,3=拒绝,4=同意" json:"status,omitempty"`
	ApplyMsg   string    `gorm:"column:apply_msg; type:varchar(1024); comment:申请原因" json:"applyMsg,omitempty"`
	RejectMsg  string    `gorm:"column:reject_msg; type:varchar(1024); comment:拒绝原因" json:"rejectMsg,omitempty"`
	RefId      uint      `gorm:"column:ref_id; type:int unsigned; comment:证书Id/Apple账号Id/测试设备Id，外码" json:"refId,omitempty"`
	CreateTime time.Time `gorm:"column:create_time; type:timestamp not null; comment:创建时间" json:"createTime,omitempty"`
	FinishTime time.Time `gorm:"column:finish_time; type:timestamp; comment:结束时间" json:"finishTime,omitempty"`
}

func (*TTodo) TableName() string {
	return "t_todo"
}
