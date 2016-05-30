package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.giftRecordRead", NewRoomGiftRecordRead)
}

type RoomGiftRecordRead struct {
	ap ActionParam
}

func NewRoomGiftRecordRead(param Param) (Action, error) {
	return &RoomGiftRecordRead{
		ap: param.Ap,
	}, nil
}

func (d *RoomGiftRecordRead) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
	}
	return
}
