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

package errs

import "net/http"

// ErrUnknownUser 登录时未知用户
var ErrUnknownUser error = &Error{
	HTTPStatus: http.StatusUnauthorized,
	Msg:        "unknown user",
}

// ErrFileNotExists 文件不存在
var ErrFileNotExists error = &Error{
	HTTPStatus: http.StatusNotFound,
	Msg:        "file not exists",
}

// ErrIllegalRequest 非法请求
var ErrIllegalRequest error = &Error{
	HTTPStatus: http.StatusExpectationFailed,
	Msg:        "illegal request",
}

// ErrTooManyRequest 请求频繁
var ErrTooManyRequest error = &Error{
	HTTPStatus: http.StatusTooManyRequests,
	Msg:        "too many request",
}

// ErrNoAuth 无授权或授权非法
var ErrNoAuth error = &Error{
	HTTPStatus: http.StatusForbidden,
	Msg:        "need authorization",
}
