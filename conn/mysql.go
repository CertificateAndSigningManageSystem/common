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
	"fmt"
	"runtime"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
	"gitee.com/CertificateAndSigningManageSystem/common/log"
	"gitee.com/CertificateAndSigningManageSystem/common/model"
)

var mysqlClient *gorm.DB

// InitialMySQL 初始化 MySQL 连接
func InitialMySQL(ctx context.Context, user, pass, host, port, db string, maxIdea, maxOpen int) {
	obj, err := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pass, host, port, db),
	),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
			Logger: logger.New(&log.GormLogFormatter{},
				logger.Config{
					IgnoreRecordNotFoundError: true,
					LogLevel:                  logger.Info,
					SlowThreshold:             3 * time.Second,
				}),
		},
	)
	if err != nil {
		log.Fatal(ctx, "init mysql error", err)
	}
	sql, err := obj.DB()
	if err != nil {
		log.Fatal(ctx, err)
	}
	sql.SetMaxIdleConns(maxIdea)
	sql.SetMaxOpenConns(maxOpen)
	// obj = obj.Debug()
	mysqlClient = obj
	runtime.SetFinalizer(mysqlClient, func(obj *gorm.DB) { CloseMysqlClient(ctx) })
	log.Info(ctx, "init mysql success")
}

// GetMySQLClient 获取 MySQL 客户端
func GetMySQLClient(ctx context.Context) *gorm.DB {
	// 若上下文有事务则使用之
	tx := ctxs.Transaction(ctx)
	if tx != nil {
		return tx
	}
	return mysqlClient.WithContext(ctx)
}

// CloseMysqlClient 断开连接
func CloseMysqlClient(ctx context.Context) {
	if mysqlClient == nil {
		return
	}
	db, err := mysqlClient.DB()
	if err != nil {
		log.Error(ctx, err)
		return
	}
	if err = db.Close(); err != nil {
		log.Error(ctx, err)
	}
}

// AutoMigrateAllTable 创建数据库表结构
func AutoMigrateAllTable(ctx context.Context) error {
	return GetMySQLClient(ctx).AutoMigrate(
		&model.TApp{},
		&model.TAuthorizationAction{},
		&model.TAuthorizationActionRelation{},
		&model.TUser{},
		&model.TUserRole{},
		&model.TAndroidSignJob{},
		&model.TAuthorization{},
		&model.TEvent{},
		&model.TFile{},
		&model.TWinSignJob{},
		&model.TAppleSignJob{},
		&model.TTodo{},
		&model.TAppleAccount{},
		&model.TTodoApprover{},
		&model.TAppleDevice{},
		&model.TAndroidCertOrg{},
		&model.TNotice{},
		&model.TAndroidCert{},
		&model.TWinCert{},
		&model.TAppleCert{},
		&model.TAppleProfile{},
		&model.TAppleBundle{},
		&model.TAppleBundleCapability{},
		&model.TAppleBundleRelation{},
	)
}
