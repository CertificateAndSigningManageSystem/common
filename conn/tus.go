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
	"net/http"

	"gitee.com/ivfzhou/tus_client"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

var tusClient tus_client.TusClient

// InitialTusClient 初始化tus客户端
func InitialTusClient(ctx context.Context, host string) {
	tusClient = tus_client.NewClient(host, tus_client.WithLogger(&log.TusClientLogger{}))
	options, err := tusClient.Options(ctx)
	if err != nil {
		log.Fatal(ctx, err)
	}
	if options.HTTPStatus != http.StatusOK && options.HTTPStatus != http.StatusNoContent {
		log.Fatal(ctx, "tus options error", options.HTTPStatus)
	}
	log.Info(ctx, "init tus client success")
}

// GetTusClient 获取tus客户端
func GetTusClient(ctx context.Context) tus_client.TusClient {
	return tusClient
}
