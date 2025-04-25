package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hello/greet/internal/logic"
	"hello/greet/internal/svc"
)

func RateLimiterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRateLimiterLogic(r.Context(), svcCtx)
		resp, err := l.RateLimiter()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
