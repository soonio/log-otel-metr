package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"hello/chinese/chinese"
	"hello/english/english"
	"hello/greet/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	English english.EnglishClient
	Chinese chinese.ChineseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		English: english.NewEnglishClient(zrpc.MustNewClient(c.English).Conn()),
		Chinese: chinese.NewChineseClient(zrpc.MustNewClient(c.Chinese).Conn()),
	}
}
