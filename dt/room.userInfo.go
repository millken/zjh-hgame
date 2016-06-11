package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/room.userInfo", NewUserInfo)
}

type UserInfo struct {
	ap ActionParam
}

func NewUserInfo(param Param) (Action, error) {
	return &UserInfo{
		ap: param.Ap,
	}, nil
}

func (d *UserInfo) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data": gin.H{
			"payPolicy": []gin.H{
				gin.H{
					"types": "sdk",
				},
			},
		},
	}
	return
}
