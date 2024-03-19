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

	"github.com/rabbitmq/amqp091-go"

	"gitee.com/CertificateAndSigningManageSystem/common/log"
)

var (
	rabbitmqClient  *amqp091.Connection
	rabbitmqChannel *amqp091.Channel
)

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
	log.Info(ctx, "init rabbitmq success")
}

// GetRabbitMQChannel 获取RabbitMQ通道
func GetRabbitMQChannel(ctx context.Context) *amqp091.Channel {
	return rabbitmqChannel
}
