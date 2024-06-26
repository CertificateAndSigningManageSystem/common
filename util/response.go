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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
	"gitee.com/CertificateAndSigningManageSystem/common/errs"
	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

// Fail 响应失败
func Fail(c *gin.Context, status int, msg string) {
	ctx := c.Request.Context()
	rid := ctxs.RequestId(ctx)
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf8")
	c.Writer.Header().Set("X-Csms-Error-Message", url.QueryEscape(msg))
	c.Writer.Header().Set("X-Csms-Request-Id", rid)
	c.Writer.WriteHeader(status)
	c.Request = c.Request.WithContext(ctxs.AppendErrMsg(ctx, msg))
}

// FailByErr 响应失败
func FailByErr(c *gin.Context, err error) {
	ctx := c.Request.Context()
	var e *errs.Error
	if !errors.As(err, &e) {
		e = &errs.Error{
			HTTPStatus: http.StatusInternalServerError,
		}
		log.Error(ctx, "unknown error", err)
	}
	rid := ctxs.RequestId(ctx)
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf8")
	c.Writer.Header().Set("X-Csms-Error-Message", url.QueryEscape(e.Msg))
	c.Writer.Header().Set("X-Csms-Request-Id", rid)
	c.Writer.WriteHeader(e.HTTPStatus)
	c.Request = c.Request.WithContext(ctxs.AppendErrMsg(ctx, err.Error()))
}

// Success 响应成功
func Success(c *gin.Context, v any) {
	ctx := c.Request.Context()
	rid := ctxs.RequestId(ctx)
	var rspBody []byte
	if v != nil {
		rspBody, _ = json.Marshal(v)
		if len(rspBody) < 5*1024 {
			log.Info(ctx, "rsqBody is", string(rspBody))
		}
	}
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf8")
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(rspBody)))
	c.Writer.Header().Set("X-Csms-Request-Id", rid)
	c.Writer.WriteHeader(http.StatusOK)
	n, err := c.Writer.Write(rspBody)
	log.ErrorIf(ctx, err)
	if n != len(rspBody) {
		log.Errorf(ctx, "write length not equal body length %d!=%d", n, len(rspBody))
	}
}

// SuccessMsg 响应成功
func SuccessMsg(c *gin.Context, v string) {
	ctx := c.Request.Context()
	rid := ctxs.RequestId(ctx)
	c.Writer.Header().Set("X-Csms-Request-Id", rid)
	c.Writer.Header().Set("X-Csms-Message", url.QueryEscape(v))
	c.Writer.WriteHeader(http.StatusOK)
}

// VendFile 响应文件流
func VendFile(c *gin.Context, fileSize int64, fileName string, fileObj io.Reader) {
	ctx := c.Request.Context()
	rid := ctxs.RequestId(ctx)
	c.Writer.Header().Set("X-Csms-Request-Id", rid)
	c.DataFromReader(
		http.StatusOK,
		fileSize,
		"application/octet-stream",
		fileObj,
		map[string]string{"Content-Disposition": fmt.Sprintf("attachment; filename=%s", url.QueryEscape(fileName))},
	)
}
