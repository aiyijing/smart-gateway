package httputil

import (
	"github.com/gin-gonic/gin"
)

type Message struct {
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func MsgWithJsonString(g *gin.Context, data string) {
	g.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	g.Writer.WriteString(data)
	g.Done()
}

func MsgWithData(g *gin.Context, data interface{}) {
	g.JSON(200, Message{
		Msg:  "success",
		Code: 200,
		Data: data,
	})
}

func Msg(g *gin.Context, err error) {
	if err != nil {
		MsgErr(g, err)
		return
	}
	MsgSuccess(g)
}

func MsgSuccess(g *gin.Context) {
	g.JSON(200, Message{
		Msg:  "success",
		Code: 200,
	})
}

func MsgErr(g *gin.Context, err error) {
	g.JSON(500, Message{
		Msg:  err.Error(),
		Code: 500,
	})
}
