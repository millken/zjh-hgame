package dt

import "github.com/gin-gonic/gin"

//编辑头像
func init() {
	DtRegister("/room.updateUserIcon", NewRoomUpdateUserIcon)
}

type RoomUpdateUserIcon struct {
	ap ActionParam
}

func NewRoomUpdateUserIcon(param Param) (Action, error) {
	return &RoomUpdateUserIcon{
		ap: param.Ap,
	}, nil
}

func (d *RoomUpdateUserIcon) Response() (data gin.H, err error) {

	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
	}
	return
}
