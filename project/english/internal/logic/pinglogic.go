package logic

import (
	"context"

	"hello/english/english"
	"hello/english/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *english.Request) (*english.Response, error) {
	// todo: add your logic here and delete this line

	return &english.Response{
		Pong: "hello english",
	}, nil
}
