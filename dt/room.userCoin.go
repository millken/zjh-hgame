package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.userCoin", NewRoomUserCoin)
}

type RoomUserCoin struct {
	ap ActionParam
}

func NewRoomUserCoin(param Param) (Action, error) {
	return &RoomUserCoin{
		ap: param.Ap,
	}, nil
}

func (d *RoomUserCoin) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
