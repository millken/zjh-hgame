package dt

import "github.com/gin-gonic/gin"

//连续登陆领取奖励
func init() {
	DtRegister("/room.getTaskReward", NewRoomGetTaskReward)
}

type RoomGetTaskReward struct {
	ap ActionParam
}

func NewRoomGetTaskReward(param Param) (Action, error) {
	return &RoomGetTaskReward{
		ap: param.Ap,
	}, nil
}

func (d *RoomGetTaskReward) Response() (data gin.H, err error) {
	//再次领取
	if false {
		data = gin.H{
			"requestId": d.ap.RequestId,
			"status":    200,
			"action":    d.ap.Action,
			"data": gin.H{
				"reward": 0,
				"status": -2,
			},
		}

	}
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data": gin.H{
			"total":  11000,
			"reward": 1000,
			"status": 2,
		},
	}
	return
}
