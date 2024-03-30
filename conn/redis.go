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
	"time"

	"github.com/redis/go-redis/v9"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

const (
	// 分布式锁键格式
	CacheKey_LockFmt = "lock:%s:string"
	// 分片上传文件
	CacheKey_UploadFiles = "upload:files:hash"
	// 记录唯一id缓存键
	CacheKey_GenIdFmt = "gen:id:%s:set"
	// 分片上传信息
	CacheKey_UploadPartFmt = "upload:part:%s:sset"
)

var redisClient *redis.Client

// InitialRedis 初始化Redis
func InitialRedis(ctx context.Context, addr, passwd string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal(ctx, "cannot connect redis", err)
	}
	log.Info(ctx, "init redis success")
}

// GetRedisClient 获取Redis客户端
func GetRedisClient(ctx context.Context) *redis.Client {
	return redisClient
}

// Lock 加锁
func Lock(ctx context.Context, key string, timeout time.Duration) bool {
	b, err := GetRedisClient(ctx).SetNX(ctx, fmt.Sprintf(CacheKey_LockFmt, key),
		time.Now().Format("20060102150405"), timeout).Result()
	if err != nil {
		log.Error(ctx, "redis lock error", err)
	}
	return b
}

// Unlock 解锁
func Unlock(ctx context.Context, key string) {
	err := GetRedisClient(ctx).Del(ctx, fmt.Sprintf(CacheKey_LockFmt, key)).Err()
	if !errors.Is(err, redis.Nil) {
		log.Error(ctx, "redis unlock error", err)
	}
}
