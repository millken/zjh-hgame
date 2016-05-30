package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/user.gameAuthorize", NewUserGameAuthorize)
}

type UserGameAuthorize struct {
	ap ActionParam
}

func NewUserGameAuthorize(param Param) (Action, error) {
	return &UserGameAuthorize{
		ap: param.Ap,
	}, nil
}

func (d *UserGameAuthorize) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
	}
	return
}
