package logic

import (
	"context"
	"fmt"
	"hello/chinese/chinese"
	"hello/english/english"

	"hello/greet/internal/svc"
	"hello/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	switch req.Name {
	case "english":
		rsp, err := l.svcCtx.English.Ping(l.ctx, &english.Request{Ping: "1"})
		if err != nil {
			return nil, err
		}
		return &types.Response{Message: rsp.Pong}, nil
	case "chinese":
		rsp, err := l.svcCtx.Chinese.Ping(l.ctx, &chinese.Request{Ping: "2"})
		if err != nil {
			return nil, err
		}
		return &types.Response{Message: rsp.Pong}, nil
	default:
		return nil, fmt.Errorf("不存在：%s", req.Name)
	}

	return
}
