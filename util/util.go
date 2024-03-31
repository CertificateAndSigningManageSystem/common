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
	"io"
	"net"
	"time"

	"gitee.com/ivfzhou/gotools/v4"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

// CloseIO 关闭流
func CloseIO(ctx context.Context, io io.Closer) {
	if io == nil {
		return
	}
	log.ErrorIf(ctx, io.Close())
}

// DoThreeTimesIfErr 最多执行三次 fn，若 err 非 nil。
func DoThreeTimesIfErr(fn func() error) error {
	var err error
	for i := 0; i < 3; i++ {
		err = fn()
		if err == nil {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	return err
}

// GetLocalNetIP 获取本地IP
func GetLocalNetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && !gotools.IsIntranet(ipnet.IP.String()) {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetLocalNoLoopIP 获取本地IP
func GetLocalNoLoopIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
