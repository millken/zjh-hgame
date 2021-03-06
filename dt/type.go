package dt

import "github.com/gin-gonic/gin"

type Action interface {
	Response() (data gin.H, err error)
}

type ActionParam struct {
	GameId     int `json:"gameId"`
	RoomId     int `json:"roomId"`
	RequestId  int `json:"requestId"`
	Action     string
	ReturnType string `json:"returnType"`
	Start      int
	Size       int
	Uid        int
}

type Param struct {
	Ap        ActionParam
	HallToken string
}

var Actions = map[string]func(Param) (Action, error){}

func DtRegister(name string, actionFactory func(Param) (Action, error)) {
	if actionFactory == nil {
		panic(" actionFactory is nil")
	}
	if _, dup := Actions[name]; dup {
		panic(" Register called twice for " + name)
	}
	Actions[name] = actionFactory
}
