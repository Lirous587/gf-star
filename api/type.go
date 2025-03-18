package api

type AuthHeader struct {
	Authorization string `in:"header" v:"required" dc:"Bearer令牌" example:"{{token}}"`
}
