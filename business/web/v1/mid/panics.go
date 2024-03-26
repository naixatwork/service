package mid

import (
	"context"
	"fmt"
	"github.com/naixatwork/service/business/web/metrics"
	"github.com/naixatwork/service/foundation/web"
	"net/http"
	"runtime/debug"
)

func Panics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("PANIC [%v] TRACE [%s]", rec, string(trace))
					metrics.AddPanics(ctx)
				}
			}()
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
