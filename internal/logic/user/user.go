package user

import (
	"context"
	"errors"
	"fmt"
	"time"
	"web_backend/internal/dao"
	"web_backend/internal/model"
	"web_backend/internal/model/do"
	"web_backend/internal/model/entity"
	"web_backend/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create creates user account.
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (err error) {
	// If Nickname is not specified, it then uses Passport as its default Nickname.
	if in.Nickname == "" {
		in.Nickname = in.Passport
	}
	var (
		available bool
	)
	// Passport checks.
	available, err = s.IsPassportAvailable(ctx, in.Passport)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Passport "%s" is already token by others`, in.Passport)
	}
	// Nickname checks.
	available, err = s.IsNicknameAvailable(ctx, in.Nickname)
	if err != nil {
		return err
	}
	if !available {
		return gerror.Newf(`Nickname "%s" is already token by others`, in.Nickname)
	}
	return dao.User.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.User.Ctx(ctx).Data(do.User{
			Passport: in.Passport,
			Password: in.Password,
			Nickname: in.Nickname,
		}).Insert()
		return err
	})
}

// IsSignedIn checks and returns whether current user is already signed-in.
func (s *sUser) IsSignedIn(ctx context.Context) bool {
	if v := service.BizCtx().Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// SignIn creates session for given user account.
func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).Where(do.User{
		Passport: in.Passport,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return
	}
	if user == nil {
		// return nil, gerror.New(`Passport or Password not correct`)
		// gerror.NewCode(gcode.New(0000001, "", nil), "Passport or Password not correct")
		// gerror.Newf(`Passport or Password not correct`)
		return nil, errors.New("Passport or Password not correct")

	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	var nowTime = time.Now()
	fmt.Println("nowTime: ", nowTime.Format("2006-01-02 15:04:05"))
	err = dao.Log.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Log.Ctx(ctx).Data(do.Log{
			Passport:   user.Passport,
			PassportId: user.Id,
			Time:       gtime.Now(),
		}).Insert()
		return err
	})
	return
}

// SignOut removes the session for current signed-in user.
func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

// IsPassportAvailable checks and returns given passport is available for signing up.
func (s *sUser) IsPassportAvailable(ctx context.Context, passport string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Passport: passport,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// IsNicknameAvailable checks and returns given nickname is available for signing up.
func (s *sUser) IsNicknameAvailable(ctx context.Context, nickname string) (bool, error) {
	count, err := dao.User.Ctx(ctx).Where(do.User{
		Nickname: nickname,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// GetProfile retrieves and returns current user info in session.
func (s *sUser) GetProfile(ctx context.Context) *entity.User {
	return service.Session().GetUser(ctx)
}

func (s *sUser) GetUserList(ctx context.Context) (userList []*entity.User, err error) {

	var (
		resultLislt []struct {
			User *entity.User
		}
	)
	err = dao.User.Ctx(ctx).ScanList(&resultLislt, "User")
	// var user *entity.User
	// err = dao.User.Ctx(ctx).Where(do.User{
	// 	Role: 1,
	// }).Scan(&user)
	fmt.Println("GetUserList: ", resultLislt)
	for _, v := range resultLislt {
		userList = append(userList, v.User)
	}
	return
}
