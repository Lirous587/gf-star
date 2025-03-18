package users

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"star/internal/dao"
	"star/internal/model/do"
)

type RegisterInput struct {
	Username string
	Password string
	Email    string
}

func (u *Users) Register(ctx context.Context, in RegisterInput) error {
	if err := u.checkUser(ctx, in.Username); err != nil {
		return err
	}

	_, err := dao.Users.Ctx(ctx).Data(do.Users{
		Username: in.Username,
		Password: u.encryptPwd(in.Password),
		Email:    in.Email,
	}).Insert()
	return err
}

func (u *Users) checkUser(ctx context.Context, username string) error {
	count, err := dao.Users.Ctx(ctx).Where("username", username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户已存在")
	}
	return nil
}

func (u *Users) encryptPwd(pwd string) string {
	return gmd5.MustEncryptString(pwd)
}
