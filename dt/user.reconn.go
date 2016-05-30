package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/user.reconn", NewUserReconn)
}

type UserReconn struct {
	ap ActionParam
}

func NewUserReconn(param Param) (Action, error) {
	return &UserReconn{
		ap: param.Ap,
	}, nil
}

func (d *UserReconn) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
		"data":      gin.H{},
	}
	return
}
