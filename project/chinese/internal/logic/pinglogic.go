package logic

import (
	"context"

	"hello/chinese/chinese"
	"hello/chinese/internal/svc"

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

func (l *PingLogic) Ping(in *chinese.Request) (*chinese.Response, error) {
	// todo: add your logic here and delete this line

	return &chinese.Response{
		Pong: "你好，中文",
	}, nil
}
