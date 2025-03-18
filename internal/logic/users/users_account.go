package users

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"star/internal/consts"
	"star/internal/dao"
	"star/internal/model/entity"
	"time"
)

type jwtClaims struct {
	Id       uint
	Username string
	jwt.RegisteredClaims
}

func (u *Users) Login(ctx context.Context, username, pwd string) (tokenStr string, err error) {
	var user entity.Users
	err = dao.Users.Ctx(ctx).Where("username", username).Scan(&user)
	if err != nil {
		return "", gerror.New("用户名或密码错误")
	}
	if user.Id == 0 {
		return "", gerror.New("用户不存在")
	}
	if user.Password != u.encryptPwd(pwd) {
		return "", gerror.New("用户名或密码错误")
	}
	uc := &jwtClaims{
		Id:       user.Id,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(6 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString(consts.JwtKey)
}

func (u *Users) Info(ctx context.Context) (user *entity.Users, err error) {
	tokenStr := g.RequestFromCtx(ctx).Request.Header.Get("Authorization")
	tokenClaims, _ := jwt.ParseWithClaims(tokenStr, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return consts.JwtKey, nil
	})
	if claims, ok := tokenClaims.Claims.(*jwtClaims); ok && tokenClaims.Valid {
		err = dao.Users.Ctx(ctx).Where("id", claims.Id).Scan(&user)
	}
	return
}
