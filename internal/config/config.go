package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DataSource      string `json:"DataSource,optional"`      // 数据库连接字符串
	MaxOpenConns    int    `json:"MaxOpenConns,optional"`    // 最大打开连接数
	MaxIdleConns    int    `json:"MaxIdleConns,optional"`    // 最大空闲连接数
	ConnMaxLifetime int    `json:"ConnMaxLifetime,optional"` // 连接最大生命周期，单位秒
	ConnMaxIdleTime int    `json:"ConnMaxIdleTime,optional"` // 连接最大空闲时间，单位秒
}
