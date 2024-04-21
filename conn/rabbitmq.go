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
	"runtime"

	"github.com/rabbitmq/amqp091-go"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

var (
	rabbitmqClient  *amqp091.Connection
	rabbitmqChannel *amqp091.Channel
)

// InitialRabbitMQ 初始化 RabbitMQ 连接
func InitialRabbitMQ(ctx context.Context, uri string) {
	connection, err := amqp091.Dial(uri)
	if err != nil {
		log.Fatal(ctx, "initial rabbitmq error", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(ctx, "initial rabbitmq error", err)
	}
	rabbitmqClient = connection
	rabbitmqChannel = channel
	runtime.SetFinalizer(rabbitmqClient, func(rabbitmqClient *amqp091.Connection) { GetRabbitMQChannel(ctx) })
	log.Info(ctx, "init rabbitmq success")
}

// GetRabbitMQChannel 获取 RabbitMQ 通道
func GetRabbitMQChannel(_ context.Context) *amqp091.Channel {
	return rabbitmqChannel
}

// CloseRabbitMQClient 关闭连接
func CloseRabbitMQClient(ctx context.Context) {
	if rabbitmqClient == nil {
		return
	}
	err := rabbitmqChannel.Close()
	if err != nil {
		log.Error(ctx, err)
	}
	if err = rabbitmqClient.Close(); err != nil {
		log.Error(ctx, err)
	}
}
