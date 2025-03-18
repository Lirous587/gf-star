package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"star/api"
)

type RandListReq struct {
	g.Meta `path:"words/rand" method:"get" sm:"随机获取单词列表" tags:"单词"`
	api.AuthHeader
	Limit uint `json:"limit" v:"between:1,300" dc:"限制个数,默认50"`
}

type RandListRes struct {
	List []List `json:"list"`
}

type SetLevelReq struct {
	g.Meta `path:"words/{id}/level" method:"patch" sm:"更新掌握水平" tags:"单词"`
	api.AuthHeader
	Id    uint             `json:"id" v:"required"`
	Level ProficiencyLevel `json:"level" v:"required|between:1,5"`
}

type SetLevelRes struct {
}
