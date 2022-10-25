package controller

import (
	"context"

	v1 "web_backend/api/v1"
	"web_backend/internal/model"
	"web_backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Log = cLog{}
)

type cLog struct{}

func (c *cHello) GetLogList(ctx context.Context, req *v1.LogListReq) (res *v1.LogListRes, err error) {
	var logList, _ = service.Log().GetLogList(ctx)
	res = (*v1.LogListRes)(&logList)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}
