package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/time/rate"
	"net/http"
)

type RateLimiterMiddleware struct {
	limiter *rate.Limiter
}

func NewRateLimiterMiddleware() *RateLimiterMiddleware {
	return &RateLimiterMiddleware{rate.NewLimiter(rate.Limit(1000), 3000)}
}

func (m *RateLimiterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if m.limiter.Allow() {
			next(w, r)
		} else {
			w.Header().Set(httpx.ContentType, httpx.JsonContentType)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"code":555,"msg":"请求火爆，请3秒后再试"}`))
		}
	}
}
