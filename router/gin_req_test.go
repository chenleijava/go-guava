package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

//test get request
//func TestGinRequestInfo(t *testing.T) {
//	r := gin.New()
//	//register middle ware
//	//get request info ，save logs or send to mq ?
//	r.Use(GinRequestInfo(func(req *RequestInfo) {
//		d, _ := json.Marshal(req)
//		log.Printf("%s", string(d))
//	}))
//	r.GET("/v1/load", func(context *gin.Context) {
//		context.JSON(http.StatusOK, gin.H{"code": 0})
//	})
//	r.Run(":7777")
//}

func TestStart(t *testing.T) {
	//start  a service
	Start(443, DebugMode, "", "", func(r *gin.Engine) {
		r.GET("/v1/load", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"code": 0})
		})
	})
}
