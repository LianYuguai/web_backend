package v1

import (
	"web_backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type LogListReq struct {
	g.Meta `path:"/log/list" method:"get" tags:"LogService" summary:"Get log list"`
}

type LogListRes []*entity.Log
