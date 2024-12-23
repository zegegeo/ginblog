package db

import (
	"context"
	"time"
	
	"goblog/blog/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

)

func NewMysql(MysqlConfig config.MysqlConfig) sqlx.SqlConn{
	mysql := sqlx.NewMysql(MysqlConfig.DataSource)
	db, err := mysql.RawDB()
	if err != nil {
		panic(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(MysqlConfig.ConnectTimeout))
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(20)
	return mysql
}