package v1

import (
	"web_backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type G2gReq struct {
	g.Meta `path:"/g2g" method:"get" tags:"G2gService" summary:"G2gService"`
}
type G2gRes []*entity.G2G

type GoldenchamDataReq struct {
	g.Meta `path:"/goldencham-data" method:"get" tags:"GoldenchamDataService" summary:"GoldenchamDataService"`
}
type GoldenchamDataRes []*entity.GoldenchamData

type GoldenchamReq struct {
	g.Meta `path:"/goldencham" method:"get" tags:"GoldenchamService" summary:"GoldenchamService"`
}
type GoldenchamRes []*entity.Goldencham

type OerReq struct {
	g.Meta `path:"/oer" method:"get" tags:"OerService" summary:"OerService"`
}
type OerRes []*entity.Oer

type OocReq struct {
	g.Meta `path:"/ooc" method:"get" tags:"OocService" summary:"OocService"`
}
type OocRes []*entity.Ooc

type PmReq struct {
	g.Meta `path:"/pm" method:"get" tags:"PmService" summary:"PmService"`
}
type PmRes []*entity.Pm

type SeasoningReq struct {
	g.Meta `path:"/seasoning" method:"get" tags:"SeasoningService" summary:"SeasoningService"`
}
type SeasoningRes []*entity.Seasoning
