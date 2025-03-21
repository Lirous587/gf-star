package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"star/internal/consts"
)

func Auth(r *ghttp.Request) {
	var tokenStr = r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return consts.JwtKey, nil
	})
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}
	r.Middleware.Next()
}
