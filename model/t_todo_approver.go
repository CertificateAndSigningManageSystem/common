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

// TTodoApprover 待办处理人表
type TTodoApprover struct {
	Id     uint `gorm:"column:id; type:int unsigned; primaryKey; autoIncrement; comment:主码" json:"id,omitempty"`
	TodoId uint `gorm:"column:todo_id int unsigned not null; comment:待办Id，外码; index:idx_todo_id"`
	UserId uint `gorm:"column:user_id int unsigned not null; comment:用户Id，外码; index:idx_user_id"`
}

func (TTodoApprover) TableName() string {
	return "t_todo_approver"
}
