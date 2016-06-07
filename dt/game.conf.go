package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/game.conf", NewGameConf)
}

type GameConf struct {
	ap ActionParam
}

func NewGameConf(param Param) (Action, error) {
	return &GameConf{
		ap: param.Ap,
	}, nil
}

func (d *GameConf) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data": gin.H{
			"speakerServerUrl": "ws://192.168.0.190:9030/speaker",
		},
	}
	return
}
