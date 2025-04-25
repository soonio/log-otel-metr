package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"hello/chinese/chinese"
	"hello/english/english"
	"hello/greet/internal/config"
	"hello/greet/internal/middleware"
)

type ServiceContext struct {
	Config      config.Config
	English     english.EnglishClient
	Chinese     chinese.ChineseClient
	RateLimiter rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RateLimiter: middleware.NewRateLimiterMiddleware().Handle,
		English:     english.NewEnglishClient(zrpc.MustNewClient(c.English).Conn()),
		Chinese:     chinese.NewChineseClient(zrpc.MustNewClient(c.Chinese).Conn()),
	}
}
