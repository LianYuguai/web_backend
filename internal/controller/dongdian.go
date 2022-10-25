package controller

import (
	"context"

	v1 "web_backend/api/v1"
	"web_backend/internal/model"
	"web_backend/internal/model/entity"
	"web_backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Dongdian = cDongdian{}
)

type cDongdian struct{}

func (c *cDongdian) GetG2g(ctx context.Context, req *v1.G2gReq) (res *v1.G2gRes, err error) {
	var list, _ = service.Dongdian().GetG2g(ctx)
	res = (*v1.G2gRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cHello) GetGoldenchamData(ctx context.Context, req *v1.GoldenchamDataReq) (res *v1.GoldenchamDataRes, err error) {
	var list, _ = service.Dongdian().GetGoldenchamData(ctx)
	res = (*v1.GoldenchamDataRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cHello) GetGoldencham(ctx context.Context, req *v1.GoldenchamReq) (res *v1.GoldenchamRes, err error) {
	var list []*entity.Goldencham
	list, err = service.Dongdian().GetGoldencham(ctx)
	res = (*v1.GoldenchamRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cDongdian) GetOer(ctx context.Context, req *v1.OerReq) (res *v1.OerRes, err error) {
	var list, _ = service.Dongdian().GetOer(ctx)
	res = (*v1.OerRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cDongdian) GetOoc(ctx context.Context, req *v1.OocReq) (res *v1.OocRes, err error) {
	var list, _ = service.Dongdian().GetOoc(ctx)
	res = (*v1.OocRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cDongdian) GetPm(ctx context.Context, req *v1.PmReq) (res *v1.PmRes, err error) {
	var list, _ = service.Dongdian().GetPm(ctx)
	res = (*v1.PmRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

func (c *cDongdian) GetSeasoning(ctx context.Context, req *v1.SeasoningReq) (res *v1.SeasoningRes, err error) {
	var list, _ = service.Dongdian().GetSeasoning(ctx)
	res = (*v1.SeasoningRes)(&list)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}
