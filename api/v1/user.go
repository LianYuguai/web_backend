package v1

import (
	"web_backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"UserService" summary:"Get user list"`
}

type User struct {
	Id       uint        `json:"id"       description:"User ID"`
	Passport string      `json:"passport" description:"User Passport"`
	Nickname string      `json:"nickname" description:"User Nickname"`
	CreateAt *gtime.Time `json:"createAt" description:"Created Time"`
	UpdateAt *gtime.Time `json:"updateAt" description:"Updated Time"`
	Role     int         `json:"role"     description:"User role"`
}

type UserListRes []*User

type UserSignInReq struct {
	g.Meta   `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Passport string `v:"required"`
	Password string `v:"required"`
}
type UserSignInRes *entity.User
