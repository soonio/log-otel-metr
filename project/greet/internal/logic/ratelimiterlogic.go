package logic

import (
	"context"
	"time"

	"hello/greet/internal/svc"
	"hello/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RateLimiterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRateLimiterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RateLimiterLogic {
	return &RateLimiterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RateLimiterLogic) RateLimiter() (resp *types.RateLimiterRsp, err error) {
	resp = &types.RateLimiterRsp{
		Datetime: time.Now().Format(time.DateTime),
	}

	return
}
