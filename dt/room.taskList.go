package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.taskList", NewRoomTaskList)
}

type RoomTaskList struct {
	ap ActionParam
}

func NewRoomTaskList(param Param) (Action, error) {
	return &RoomTaskList{
		ap: param.Ap,
	}, nil
}

func (d *RoomTaskList) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
