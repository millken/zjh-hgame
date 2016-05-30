package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.giftRecord", NewRoomGiftRecord)
}

type RoomGiftRecord struct {
	ap ActionParam
}

func NewRoomGiftRecord(param Param) (Action, error) {
	return &RoomGiftRecord{
		ap: param.Ap,
	}, nil
}

func (d *RoomGiftRecord) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
