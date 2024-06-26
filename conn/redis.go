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

package conn

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/redis/go-redis/v9"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

const (
	// CacheKey_LockFmt 分布式锁键格式
	CacheKey_LockFmt = "lock:%s:string"
	// CacheKey_UploadFiles 分片上传文件
	CacheKey_UploadFiles = "upload:files:hash"
	// CacheKey_GenIdFmt 记录唯一id缓存键
	CacheKey_GenIdFmt = "gen:id:%s:set"
	// CacheKey_UploadPartFmt 分片上传信息
	CacheKey_UploadPartFmt = "upload:part:%s:sset"
	// CacheKey_CronRecordFmt 定时任务记录
	CacheKey_CronRecordFmt = "cron:%s:%s"
	// CacheKey_UserSessionFmt 用户会话
	CacheKey_UserSessionFmt = "user:session:%s:%s:string"
	// CacheKey_UserLoginFailTimesFmt 登陆失败次数记录
	CacheKey_UserLoginFailTimesFmt = "user:login:fail:times:%s:string"
)

const (
	LockKey_UserRegisterFmt     = "user:register:%s"
	LockKey_UserChangeAvatarFmt = "user:change:avatar:%d"
	LockKey_UserLoginFmt        = "user:login:%d"
	LockKey_OpenApiCreateFmt    = "openapi:create:%d:%s"
)

var redisClient *redis.Client

// InitialRedis 初始化 Redis 连接
func InitialRedis(ctx context.Context, addr, passwd string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	err := redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatal(ctx, "cannot connect redis", err)
	}
	runtime.SetFinalizer(redisClient, func(redisClient *redis.Client) { CloseRedisClient(ctx) })
	log.Info(ctx, "init redis success")
}

// CloseRedisClient 关闭连接
func CloseRedisClient(ctx context.Context) {
	if redisClient == nil {
		return
	}
	err := redisClient.Close()
	if err != nil {
		log.Error(ctx, err)
	}
}

// GetRedisClient 获取 Redis 客户端
func GetRedisClient(_ context.Context) *redis.Client {
	return redisClient
}

// Lock 加锁
func Lock(ctx context.Context, key string, timeout time.Duration) bool {
	b, err := GetRedisClient(ctx).SetNX(ctx, fmt.Sprintf(CacheKey_LockFmt, key),
		time.Now().Format("20060102150405"), timeout).Result()
	if err != nil {
		log.Error(ctx, err)
	}
	return b
}

// LockWait 加锁
func LockWait(ctx context.Context, key string, wait time.Duration) bool {
	now := time.Now()
	if Lock(ctx, key, 0) {
		return true
	}
	for time.Since(now) < wait {
		runtime.Gosched()
		time.Sleep(time.Millisecond * 100)
		if Lock(ctx, key, 0) {
			return true
		}
	}
	return false
}

// Unlock 解锁
func Unlock(ctx context.Context, key string) {
	err := GetRedisClient(ctx).Del(ctx, fmt.Sprintf(CacheKey_LockFmt, key)).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error(ctx, err)
	}
}

// HasLock 是否加锁了
func HasLock(ctx context.Context, key string) bool {
	result, err := GetRedisClient(ctx).Exists(ctx, fmt.Sprintf(CacheKey_LockFmt, key)).Result()
	if err != nil {
		return true
	}
	return result > 0
}

// WaitUnlock 等待锁释放
func WaitUnlock(ctx context.Context, key string, max time.Duration) bool {
	now := time.Now()
	for HasLock(ctx, key) {
		if max > 0 && time.Since(now) > max {
			return false
		}
		runtime.Gosched()
		time.Sleep(time.Millisecond * 100)
	}

	return true
}
