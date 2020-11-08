package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
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

var c = make(chan int)
var a string
var initWait sync.WaitGroup

func f() {
	a = "hello, world" // write  value
	initWait.Done()
	log.Printf("1111")
	c <- 0 // send, channel no buffer ,wait receiver do it
	log.Printf("receiver got value, chain wait  done ,then close chain ")
	close(c)
}

var done bool
var once sync.Once

func setup() {
	a = "hello, world111111" //
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	log.Printf("%s", a)
}

//https://gin-gonic.com/docs/examples/
func TestStart(t *testing.T) {

	//go doprint()
	//go doprint()

	log.Printf(">>>>>>>>")
	initWait.Add(1)

	go f()

	log.Printf("呃呃呃额额  ")

	//block util wait done
	initWait.Wait()

	d, ok := <-c
	if ok {
		log.Printf("yes receiver get value:%d value:%s", d, a) //race conditions
	}

	//always block , if no data send to chain c
	//block until receive  a value

	//	for {
	//		select {
	//		case d, ok := <-c: // ok  , read value
	//			if ok {
	//				log.Printf("yes receiver get value:%d value:%s", d, a) //race conditions
	//				goto START
	//			}
	//		}
	//	}
	//
	//	//yeah , goto
	//START:

	log.Printf("yeah %s", a)

	user := &User{}
	log.Printf("out ptr:%p", &user)
	testPtr(user)
	log.Printf("out ptr:%p", &user)

	//start  a service
	//Start(443, DebugMode, "", "", func(r *gin.Engine) {
	//	r.GET("/v1/load", func(context *gin.Context) {
	//		context.JSON(http.StatusOK, gin.H{"code": 0})
	//	})
	//})
	var mu sync.Mutex
	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	r := gin.Default()

	//logger
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	// your custom format
	//	return fmt.Sprintf("clientIP:%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format("2006-01-02 15:04:05.000000"),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))

	//graceful restart or stop
	//https://gin-gonic.com/docs/examples/graceful-restart-or-stop/

	//debug print router func
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint: %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	//test post
	r.POST("/v1/load", func(context *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		// in heap - stack
		var user User
		err := context.ShouldBindJSON(&user) //pass by value of user variable address ; copy pointer!
		if err != nil {
			log.Fatalf("bind err:%s", err.Error())
		}
		response(context, user)
	})

	r.GET("/v1/bindQuery", func(context *gin.Context) {
		var u User
		err := context.BindQuery(&u)
		if err != nil {
			log.Fatalf("bind query err:%s", err.Error())
		}
		response(context, &u) // 取地址，copy pointer
	})

	//
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	//run http server
	//https://gin-gonic.com/docs/examples/custom-http-config/
	err := r.Run(":9091")
	if err != nil {
		log.Printf("run gin err:%s", err.Error())
	}
}

//value copy!
type User struct {
	Uid  int32  `form:"uid" json:"uid,omitempty"`
	Name string `form:"name" json:"name,omitempty"`
}

//
func response(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, &BaseResponse{
		Code: 0,
		Msg:  "",
		Data: obj,
	})
}

func responseErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, &BaseResponse{
		Code: 0,
		Msg:  err,
	})
}

type BaseResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func testPtr(user *User) {
	user.Uid = 100
	log.Printf("in ptr:%p", &user)
}
