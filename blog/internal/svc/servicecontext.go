package svc

import (
	"goblog/blog/internal/config"
	"goblog/blog/internal/db"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := db.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Mysql:  mysql,
	}
}
