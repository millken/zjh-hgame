package dt

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/millken/zjh-hgame/common"
)

func init() {
	DtRegister("/user.gameAuthorize", NewUserGameAuthorize)
}

type UserGameAuthorize struct {
	ap      ActionParam
	session *common.Session
}

func NewUserGameAuthorize(param Param) (Action, error) {
	return &UserGameAuthorize{
		ap:      param.Ap,
		session: param.Session,
	}, nil
}

func (d *UserGameAuthorize) Response() (data gin.H, err error) {
	var id int
	uid, err := d.session.Get("uid")
	if err != nil {
		log.Printf("[ERROR] get session err: %s", err)
	}
	switch uid.(type) {
	case string:
		id = common.StrToInt(uid.(string))
	case int:
		id = uid.(int)
	}
	gameServerToken := fmt.Sprintf("%s-%d", common.Guid(), id)
	if err = d.session.Set("gameServerToken", gameServerToken); err != nil {
		log.Printf("[ERROR] set session[gameServerToken] err: %s", err)
	}
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gameServerToken,
	}
	return
}
