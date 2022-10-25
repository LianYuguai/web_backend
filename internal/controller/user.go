package controller

import (
	"context"
	"encoding/json"
	"math/rand"
	"strings"

	v1 "web_backend/api/v1"
	"web_backend/internal/model"
	"web_backend/internal/model/entity"
	"web_backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	var user *entity.User
	user, err = service.User().SignIn(ctx, model.UserSignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
			Code:    "0000001",
			Message: "用户不存在",
			Data:    nil,
		})
		return
	}
	var data = make(map[string]interface{})
	data["user"] = user
	data["token"] = RandAllString(16)

	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    data,
	})
	return
}

func (c *cHello) GetUserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	var userList, _ = service.User().GetUserList(ctx)
	var listData []byte
	listData, err = json.Marshal(&userList)
	res = &v1.UserListRes{}
	err = json.Unmarshal(listData, res)
	g.RequestFromCtx(ctx).Response.WriteJson(model.CommonRes{
		Code:    "0",
		Message: "",
		Data:    res,
	})
	return
}

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

/*
RandAllString  生成随机字符串([a~zA~Z0~9])

	lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}
