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

import (
	"errors"
	"net/http"
)

// Error 错误
type Error struct {
	HTTPStatus int
	Msg        string
	WrappedErr error
}

// Msg 获取错误描述
func Msg(err error) string {
	var e *Error
	if errors.As(err, &e) && e != nil {
		return e.Msg
	}
	return ""
}

// Unwrap 获取底层错误
func Unwrap(err error) error {
	var e *Error
	if errors.As(err, &e) && e != nil {
		return e.WrappedErr
	}
	return nil
}

// NewSystemBusyErr 新建系统错误
func NewSystemBusyErr(err error) error {
	return &Error{
		Msg:        "system busy",
		WrappedErr: err,
		HTTPStatus: http.StatusInternalServerError,
	}
}

// NewParamsErr 新建参数错误
func NewParamsErr(err error) error {
	return &Error{
		HTTPStatus: http.StatusExpectationFailed,
		Msg:        "parameters invalid",
		WrappedErr: err,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

func (e *Error) Unwrap() error {
	return e.WrappedErr
}
