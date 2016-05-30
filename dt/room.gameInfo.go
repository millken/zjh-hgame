package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.gameInfo", NewRoomGameInfo)
}

type RoomGameInfo struct {
	ap ActionParam
}

func NewRoomGameInfo(param Param) (Action, error) {
	return &RoomGameInfo{
		ap: param.Ap,
	}, nil
}

func (d *RoomGameInfo) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
