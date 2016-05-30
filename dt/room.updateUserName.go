package dt

import "github.com/gin-gonic/gin"

//编辑资料
func init() {
	DtRegister("/room.updateUserName", NewRoomUpdateUserName)
}

type RoomUpdateUserName struct {
	ap ActionParam
}

func NewRoomUpdateUserName(param Param) (Action, error) {
	return &RoomUpdateUserName{
		ap: param.Ap,
	}, nil
}

func (d *RoomUpdateUserName) Response() (data gin.H, err error) {

	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
	}
	return
}
