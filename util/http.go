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

package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
)

var httpClient = http.DefaultClient

// HTTPJsonGet 发送 HTTP GET 请求，并反序列化响应数据
func HTTPJsonGet[T any](ctx context.Context, url string) (*T, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	rid := ctxs.RequestId(ctx)
	if len(rid) > 0 {
		req.Header.Set("X-Csms-Request-Id", rid)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	rspBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ret T
	if err = json.Unmarshal(rspBody, &ret); err != nil {
		return nil, fmt.Errorf("%d %s %v %s", resp.StatusCode, resp.Status, err, string(rspBody))
	}

	return &ret, nil
}
