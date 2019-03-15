package http_client

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"testing"
)

var logger, _ = zap.NewDevelopment()
var sugar = logger.Sugar()

//
type Message struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//
func (msg *Message) ToString() string {
	return fmt.Sprintf("code:%d msg:%s", msg.Code, msg.Msg)
}

//
func TestGetBytes(t *testing.T) {
	data, err := GetBytes("http://localhost:9527/test/get")
	if err != nil {
		logger.DPanic("get func error", zap.String("errkey", err.Error()))
	}
	msg := &Message{}
	json.Unmarshal(data, msg)
	sugar.Debugf("data%s", string(data))
}
func TestGetObject(t *testing.T) {
	msg := &Message{}
	GetObject("http://localhost:9527/test/get", msg)
}
func TestPost(t *testing.T) {
	msg := &Message{Code: 12, Msg: "来自客户端"}
	data, _ := Post("http://localhost:9527/test/post", msg)
	json.Unmarshal(data, msg)
	sugar.Debugf("data%s", string(data))
}
func TestPostGetObject(t *testing.T) {
	msg := &Message{Code: 13, Msg: "来自客户端"}
	obj := &Message{}
	PostGetObject("http://localhost:9527/test/post", msg, obj)
	sugar.Debugf("data%s", obj.ToString())
}

func TestHead(t *testing.T) {
	v := Head("http://cdn.atmob.com/pic/gnu.png")
	log.Printf("%d", v) // status code
}
