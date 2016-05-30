package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/user.enterGame", NewUserEnterGame)
}

type UserEnterGame struct {
	ap ActionParam
}

func NewUserEnterGame(param Param) (Action, error) {
	return &UserReconn{
		ap: param.Ap,
	}, nil
}

func (d *UserEnterGame) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data": gin.H{
			"gameServerUrl": "ws://123.59.85.45:8010/game",
			"type":          1,
		},
	}
	return
}
