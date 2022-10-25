package dongdian

import (
	"context"
	"fmt"
	"web_backend/internal/dao"
	"web_backend/internal/model/entity"
	"web_backend/internal/service"
)

type (
	sDongdian struct{}
)

func init() {
	service.RegisterDongdian(New())
}

func New() *sDongdian {
	return &sDongdian{}
}

func (s *sDongdian) GetG2g(ctx context.Context) (list []*entity.G2G, err error) {
	var (
		resultLislt []struct {
			Data *entity.G2G
		}
	)
	err = dao.G2G.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetGoldenchamData(ctx context.Context) (list []*entity.GoldenchamData, err error) {
	var (
		resultLislt []struct {
			Data *entity.GoldenchamData
		}
	)
	err = dao.GoldenchamData.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetGoldencham(ctx context.Context) (list []*entity.Goldencham, err error) {
	var (
		resultLislt []struct {
			Data *entity.Goldencham
		}
	)
	err = dao.Goldencham.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetOer(ctx context.Context) (list []*entity.Oer, err error) {
	var (
		resultLislt []struct {
			Data *entity.Oer
		}
	)
	err = dao.Oer.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetOoc(ctx context.Context) (list []*entity.Ooc, err error) {
	var (
		resultLislt []struct {
			Data *entity.Ooc
		}
	)
	err = dao.Ooc.Ctx(ctx).ScanList(&resultLislt, "Data")
	fmt.Println("GetOoc: ", resultLislt)
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetPm(ctx context.Context) (list []*entity.Pm, err error) {
	var (
		resultLislt []struct {
			Data *entity.Pm
		}
	)
	err = dao.Pm.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
func (s *sDongdian) GetSeasoning(ctx context.Context) (list []*entity.Seasoning, err error) {
	var (
		resultLislt []struct {
			Data *entity.Seasoning
		}
	)
	err = dao.Seasoning.Ctx(ctx).ScanList(&resultLislt, "Data")
	for _, v := range resultLislt {
		list = append(list, v.Data)
	}
	return
}
