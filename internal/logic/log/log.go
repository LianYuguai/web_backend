package log

import (
	"context"
	"web_backend/internal/dao"
	"web_backend/internal/model/entity"
	"web_backend/internal/service"
)

type (
	sLog struct{}
)

func init() {
	service.RegisterLog(New())
}

func New() *sLog {
	return &sLog{}
}

func (s *sLog) GetLogList(ctx context.Context) (userList []*entity.Log, err error) {

	var (
		resultLislt []struct {
			Log *entity.Log
		}
	)
	err = dao.Log.Ctx(ctx).ScanList(&resultLislt, "Log")
	// var user *entity.User
	// err = dao.User.Ctx(ctx).Where(do.User{
	// 	Role: 1,
	// }).Scan(&user)
	for _, v := range resultLislt {
		userList = append(userList, v.Log)
	}
	return
}
