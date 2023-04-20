package svc

import (
	"blog_template/internal/config"
	"blog_template/models/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel user.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: user.NewUsersModel(conn, c.CacheRedis),
	}
}
