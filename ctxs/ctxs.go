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
	ctxKey_UserId      ctxKey = "ctxKey_UserId"
	ctxKey_APIAuthId   ctxKey = "ctxKey_APIAuthId"
	ctxKey_Trace       ctxKey = "ctxKey_Trace"
	ctxKey_CallLine    ctxKey = "ctxKey_CallLine"
	ctxKey_RequestId   ctxKey = "ctxKey_RequestId"
	ctxKey_RequestIP   ctxKey = "ctxKey_RequestIP"
	ctxKey_RequestPath ctxKey = "ctxKey_RequestPath"
	ctxKey_Transaction ctxKey = "ctxKey_Transaction"
	ctxKey_Func        ctxKey = "ctxKey_Func"
)

type ctxKey string

// UserId 获取ctx中的用户Id
func UserId(ctx context.Context) uint {
	if ctx == nil {
		return 0
	}
	value := ctx.Value(ctxKey_UserId)
	userID, _ := value.(uint)
	return userID
}

// APIAuthId 获取ctx中的API凭证Id
func APIAuthId(ctx context.Context) uint {
	if ctx == nil {
		return 0
	}
	value := ctx.Value(ctxKey_APIAuthId)
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

// CallLine 获取log打印代码行
func CallLine(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	value := ctx.Value(ctxKey_CallLine)
	callLine, _ := value.(string)
	return callLine
}

// RequestId 获取上下文中的链路Id
func RequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	rid, _ := ctx.Value(ctxKey_RequestId).(string)
	return rid
}

// RequestIP 获取上下文中的请求IP
func RequestIP(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	ip, _ := ctx.Value(ctxKey_RequestIP).(string)
	return ip
}

// RequestPath 获取上下文中的请求Path
func RequestPath(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	path, _ := ctx.Value(ctxKey_RequestPath).(string)
	return path
}

// Transaction 获取上下文中的数据库事务对象
func Transaction(ctx context.Context) *gorm.DB {
	if ctx == nil {
		return nil
	}
	tx, _ := ctx.Value(ctxKey_Transaction).(*gorm.DB)
	return tx
}

// Func 获取上下文中的函数信息
func Func(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	fn, _ := ctx.Value(ctxKey_Func).(string)
	return fn
}

// WithUserId 设置上下文中的UserId
func WithUserId(ctx context.Context, userId uint) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_UserId, userId)
}

// WithAPIAuthId 设置上下文中的AuthId
func WithAPIAuthId(ctx context.Context, authId uint) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_APIAuthId, authId)
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

// WithRequestId 上下文附带链路请求标识
func WithRequestId(ctx context.Context, rid string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_RequestId, rid)
}

// WithRequestIP 上下文附带链路请求IP
func WithRequestIP(ctx context.Context, ip string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_RequestIP, ip)
}

// WithRequestPath 上下文附带链路请求Path
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

// WithFunc 设置上下文函数信息
func WithFunc(ctx context.Context, fn string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxKey_Func, fn)
}

// NewCtx 新建上下文对象
func NewCtx(fn string) context.Context {
	ctx := context.Background()
	ctx = WithFunc(ctx, fn)
	return ctx
}

// CloneCtx 克隆上下文对象，复制请求Id，请求IP、函数和用户凭证信息。
func CloneCtx(ctx context.Context) context.Context {
	newCtx := NewCtx(Func(ctx))
	newCtx = WithRequestId(newCtx, RequestId(ctx))
	newCtx = WithRequestIP(newCtx, RequestIP(ctx))
	newCtx = WithAPIAuthId(newCtx, APIAuthId(ctx))
	newCtx = WithUserId(newCtx, UserId(ctx))
	return newCtx
}
