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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net"
	"time"
	"unicode"

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

// IsAllHanCharacters 字符串是否是汉字
func IsAllHanCharacters(str string) bool {
	if len(str) <= 0 {
		return false
	}
	for _, v := range []rune(str) {
		if !unicode.Is(unicode.Han, v) {
			return false
		}
	}
	return true
}

// IsAllLetterCharacters 字符串是否是字母
func IsAllLetterCharacters(str string) bool {
	if len(str) <= 0 {
		return false
	}
	for _, v := range []rune(str) {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true
}

// CalcSum 计算摘要
func CalcSum(bs []byte) (_md5, _sha1, _sha256 string) {
	sum := md5.Sum(bs)
	_md5 = hex.EncodeToString(sum[:])
	sum1 := sha1.Sum(bs)
	_sha1 = hex.EncodeToString(sum1[:])
	sum2 := sha256.Sum256(bs)
	_sha256 = hex.EncodeToString(sum2[:])
	return
}
