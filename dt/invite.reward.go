package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/invite.reward", NewInviteReward)
}

type InviteReward struct {
	ap ActionParam
}

func NewInviteReward(param Param) (Action, error) {
	return &InviteReward{
		ap: param.Ap,
	}, nil
}

func (d *InviteReward) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
