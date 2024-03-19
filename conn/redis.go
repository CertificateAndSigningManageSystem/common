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

	"github.com/redis/go-redis/v9"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

var redisClient *redis.Client

func InitialRedis(ctx context.Context, addr, passwd string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf(ctx, "cannot connect redis; %v", err)
	}
	log.Info(ctx, "init redis success")
}

// GetRedisClient 获取Redis客户端
func GetRedisClient(ctx context.Context) *redis.Client {
	return redisClient
}
