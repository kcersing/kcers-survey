package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers-survey/biz/pkg/errno"
	"net/http"
	"time"
)

// Response 返回的结果：
type Response struct {
	Time      string      `json:"time"`
	Code      int64       `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     int64       `json:"total"`
	CacheTime string      `json:"cacheTime"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}, Total int64, cacheTime string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		Data:      data,
		Total:     Total,
		Time:      time.Now().Format(time.DateTime),
		CacheTime: cacheTime,
	})
}
func SendErr(c *app.RequestContext, err error, data interface{}, cacheTime string) {
	hlog.Info(err)
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		Code:      Err.ErrCode,
		Message:   Err.ErrMsg,
		Data:      data,
		Time:      time.Now().Format(time.DateTime),
		CacheTime: cacheTime,
	})
}
