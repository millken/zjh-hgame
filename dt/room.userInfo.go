package dt

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func init() {
	DtRegister("/room.userInfo", NewDtConf)
}

type UserInfo struct {
	ap ActionParam
	db *sqlx.DB
}

func NewUserInfo(param Param) (Action, error) {
	return &UserInfo{
		ap: param.Ap,
		db: param.Db,
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
