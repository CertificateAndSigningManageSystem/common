package conn

import (
	"context"

	"gorm.io/gorm"
)

var mysqlClient *gorm.DB

func InitialMySQL(ctx context.Context) {

}

func GetMySQLConn(ctx context.Context) *gorm.DB {
	return mysqlClient.WithContext(ctx)
}
