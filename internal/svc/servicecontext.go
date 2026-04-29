package svc

import (
	"database/sql"
	"fmt"
	"time"
	"user_rpc/internal/config"
	"user_rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel // 注入用户模型
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println(c.DataSource)
	sqlConn := sqlx.NewMysql(c.DataSource)
	//  获取底层的 *sql.DB 对象，以便配置连接池参数
	if rawDB, ok := sqlConn.(interface{ RawDB() *sql.DB }); ok {
		db := rawDB.RawDB()
		db.SetMaxOpenConns(c.MaxOpenConns)                                    // 设置最大打开连接数
		db.SetMaxIdleConns(c.MaxIdleConns)                                    // 设置最大空闲连接数
		db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second) // 设置连接最大生命周期
		db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second) // 设置空闲连接超时时间
	}

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn),
	}
}
