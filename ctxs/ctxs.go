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

package ctxs

import (
	"context"

	"gorm.io/gorm"
)

const (
	ctxKey_UserID      ctxKey = "ctxKey_UserID"
	ctxKey_APIAuthID   ctxKey = "ctxKey_APIAuthID"
	ctxKey_Trace       ctxKey = "ctxKey_Trace"
	ctxKey_CallLine    ctxKey = "ctxKey_CallLine"
	ctxKey_RequestID   ctxKey = "ctxKey_RequestID"
	ctxKey_RequestIP   ctxKey = "ctxKey_RequestIP"
	ctxKey_RequestPath ctxKey = "ctxKey_RequestPath"
	ctxKey_Transaction ctxKey = "ctxKey_Transaction"
)

type ctxKey string

// UserID 获取ctx中的用户标识
func UserID(ctx context.Context) uint {
	if ctx == nil {
		return 0
	}
	value := ctx.Value(ctxKey_UserID)
	userID, _ := value.(uint)
	return userID
}

// APIAuthID 获取ctx中的API凭证标识
func APIAuthID(ctx context.Context) uint {
	if ctx == nil {
		return 0
	}
	value := ctx.Value(ctxKey_APIAuthID)
	apiAuthID, _ := value.(uint)
	return apiAuthID
}

// Trace 获取ctx中的错误堆栈信息
func Trace(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value := ctx.Value(ctxKey_Trace)
	trace, _ := value.(string)
	return trace
}

// CallLine 获取log打印行
func CallLine(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value := ctx.Value(ctxKey_CallLine)
	callLine, _ := value.(string)
	return callLine
}

// RequestID 获取上下文中的rid
func RequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	rid, _ := ctx.Value(ctxKey_RequestID).(string)
	return rid
}

// RequestIP 获取上下文中的ip
func RequestIP(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	ip, _ := ctx.Value(ctxKey_RequestIP).(string)
	return ip
}

// RequestPath 获取上下文中的path
func RequestPath(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	path, _ := ctx.Value(ctxKey_RequestPath).(string)
	return path
}

func Transaction(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return nil
	}
	tx, _ := ctx.Value(ctxKey_Transaction).(*gorm.DB)
	return tx
}

// WithUserID 设置上下文中的UserID
func WithUserID(ctx context.Context, userID uint) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_UserID, userID)
}

// WithAPIAuthID 设置上下文中的AuthID
func WithAPIAuthID(ctx context.Context, authID uint) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_APIAuthID, authID)
}

// WithTrace ctx添加堆栈信息
func WithTrace(ctx context.Context, trace string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_Trace, trace)
}

// WithCallLine ctx添加log打印行
func WithCallLine(ctx context.Context, callLine string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_CallLine, callLine)
}

// WithRequestID 上下文附带链路请求标识
func WithRequestID(ctx context.Context, rid string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_RequestID, rid)
}

// WithRequestIP 上下文附带链路请求ip
func WithRequestIP(ctx context.Context, ip string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_RequestIP, ip)
}

// WithRequestPath 上下文附带链路请求path
func WithRequestPath(ctx context.Context, path string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_RequestPath, path)
}

// WithTransaction 上下文添加事务
func WithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_Transaction, tx)
}
