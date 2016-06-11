package dt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/millken/zjh-hgame/common"
	"github.com/millken/zjh-hgame/rd"
)

func init() {
	DtRegister("/user.gameAuthorize", NewUserGameAuthorize)
}

type UserGameAuthorize struct {
	ap    ActionParam
	param Param
}

func NewUserGameAuthorize(param Param) (Action, error) {
	return &UserGameAuthorize{
		ap:    param.Ap,
		param: param,
	}, nil
}

func (d *UserGameAuthorize) Response() (data gin.H, err error) {
	var id int
	hallToken := d.param.HallToken
	id = rd.GetInt(fmt.Sprintf("hallToken:%s:uid", hallToken))
	guid := common.Guid()
	gameServerToken := fmt.Sprintf("%s-%d", guid, id)
	if err = rd.Set(fmt.Sprintf("u:%d:gameServerToken", id), gameServerToken); err != nil {
		log.Printf("[ERROR] set session[gameServerToken] err: %s", err)
	}
	rd.Set(fmt.Sprintf("u:%d:roomId", id), d.ap.RoomId)
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gameServerToken,
	}
	return
}
