package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"kcers-survey/biz/pkg/errno"
)

// ServiceFunc is a service call that takes a validated request and returns data/total or error.
type ServiceFunc[Req any] func(ctx context.Context, c *app.RequestContext, req Req) (data any, total int64, err error)

// Handle wraps the common handler pattern: bind → validate → call service → respond.
// Returns 204 No Content for void operations (data=nil, total=0, err=nil).
func Handle[Req any](fn ServiceFunc[Req]) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req Req
		if err := c.BindAndValidate(&req); err != nil {
			SendResponse(c, errno.ConvertErr(err), nil, 0, "")
			return
		}
		data, total, err := fn(ctx, c, req)
		if err != nil {
			SendResponse(c, errno.ConvertErr(err), nil, 0, "")
			return
		}
		SendResponse(c, errno.Success, data, total, "")
	}
}
