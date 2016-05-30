package dt

import "github.com/gin-gonic/gin"

func init() {
	DtRegister("/pay.payOrderList", NewPayPayOrderList)
}

type PayPayOrderList struct {
	ap ActionParam
}

func NewPayPayOrderList(param Param) (Action, error) {
	return &PayPayOrderList{
		ap: param.Ap,
	}, nil
}

func (d *PayPayOrderList) Response() (data gin.H, err error) {
	data = gin.H{
		"requestId": d.ap.RequestId,
		"status":    200,
		"action":    d.ap.Action,
	}
	return
}
