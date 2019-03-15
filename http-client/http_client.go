package http_client

import (
	"github.com/astaxie/beego/httplib"
	"time"
)

const (
	connectTimeOut   = 30 * time.Second
	readWriteTimeOut = 30 * time.Second
)


func Head(url string) int  {
	r,_:=httplib.Head(url).Response()
	return r.StatusCode
}

//Get bytes
func GetBytes(url string) ([]byte, error) {
	return httplib.Get(url).SetTimeout(connectTimeOut, readWriteTimeOut).Bytes()
}

//Get string
func GetString(url string) (string, error) {
	return httplib.Get(url).SetTimeout(connectTimeOut, readWriteTimeOut).String()
}

//Get JSON
//v , ser json struck, ref  address
func GetObject(url string, target interface{}) error {
	return httplib.Get(url).SetTimeout(connectTimeOut, readWriteTimeOut).ToJSON(target)
}

//POST object data by json
//content-type application/json
//return json data --- bytes
func Post(url string, serializationObj interface{}) ([]byte, error) {
	req, e := httplib.Post(url).SetTimeout(connectTimeOut, readWriteTimeOut).JSONBody(serializationObj)
	if e != nil {
		return nil, e
	}
	return req.Bytes()
}

//Post
func PostBytes(url string) ([]byte, error) {
	req, e := httplib.Post(url).SetTimeout(connectTimeOut, readWriteTimeOut).Bytes()
	if e != nil {
		return nil, e
	}
	return req, nil
}

//POST object data by json
//content-type application/json
//return json data --- target
func PostGetObject(url string, serializationObj, target interface{}) error {
	req, e := httplib.Post(url).SetTimeout(connectTimeOut, readWriteTimeOut).JSONBody(serializationObj)
	if e != nil {
		return e
	}
	return req.ToJSON(target)
}
