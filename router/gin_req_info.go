package router

import (
	"bytes"
	"github.com/chenleijava/go-guava/bufpool"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/log"
	"time"
)

//base request info
type RequestInfo struct {
	Status       int    `json:"status"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	Query        string `json:"query"`
	Ip           string `json:"ip"`
	UserAgent    string `json:"userAgent"`
	RequestTime  string `json:"requestTime"`
	ResponseTime string `json:"responseTime"`
	Latency      string `json:"cost"`
	ResponseData string `json:"responseData"`
}

//read response body
type bodyWriter struct {
	gin.ResponseWriter
	bodyBuffer *bytes.Buffer
}

//interceptor writer
//copy response data other buffer
//write data to response
func (w *bodyWriter) Write(b []byte) (int, error) {
	if w.bodyBuffer != nil {
		w.bodyBuffer.Write(b)
		return w.ResponseWriter.Write(b)
	}
	return 0, nil
}

//Get the basic information of the gin request
//f hand RequestInfo
func GinRequestInfo(f func(req *RequestInfo)) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middle wares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery //query string
		//replace writer
		bodyBuf := bufpool.GetStringBuffer()
		defer bufpool.PutStringBuffer(bodyBuf)

		bodyWriter := &bodyWriter{ResponseWriter: c.Writer, bodyBuffer: bodyBuf}
		c.Writer = bodyWriter

		//call logic controller
		c.Next()

		//done
		end := time.Now()
		latency := end.Sub(start) //cost time
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				log.Errorf("Append error field if this is an erroneous request,error:%s", e)
			}
		} else {
			//build request and response detail
			req := &RequestInfo{
				Status:       c.Writer.Status(),                          //status
				Method:       c.Request.Method,                           // request method
				Path:         path,                                       //uri
				Query:        query,                                      //query string
				Ip:           c.ClientIP(),                               // client ip
				UserAgent:    c.Request.UserAgent(),                      //user-agent
				RequestTime:  start.Format("2006-01-02 15:04:05.000000"), // requestTime
				ResponseTime: end.Format("2006-01-02 15:04:05.000000"),   //  response time
				Latency:      latency.String(),                           //  cost
				ResponseData: bodyWriter.bodyBuffer.String(),             //get copy data form body buffer
			}
			//pass to hand function/
			f(req)
		}

	}
}
