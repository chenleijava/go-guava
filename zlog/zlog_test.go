package zlog

import (
	"github.com/chenleijava/go-guava/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestNewZlog(t *testing.T) {
	mqLogger := GetMqLogger()
	r := gin.New()
	r.Use(router.GinRequestInfo(func(req *router.RequestInfo) {
		//log flg  below to which log file
		mqLogger.Info("请求数据详情", zap.Any("request_data_key", req)) //
	}))
	r.GET("/mqlog", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 0})
	})
	log.Printf("http://localhost:8081/mqlog")
	_ = r.Run(":8081")
}

func TestCharts(t *testing.T) {

}

//deliveryRoutingKey=delivery_routing_key
//deliveryExchange=delivery_exchange
//deliveryQueue=delivery_queue
func GetMqLogger() *zap.Logger {
	syncer := GetRabbitMqWriteSyncer()
	syncer.InitRabbitMqWriteSyncerDefault(
		"amqp://chenlei:123@localhost:5672/",
	)
	//pb.log_flg_map[id]=
	return syncer.NewLog2RabbitMq(AddLogFlg("WxTokenLogFlg"))
}

//
func TestNewLog2FileByLumberJackLog(t *testing.T) {
	_log := NewLog2FileByLumberJackLog("./log/jack.log", 1, 0, 0)
	for true {
		time.Sleep(2 * time.Millisecond)
		_, _ = _log.Write([]byte(`//1.除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
//2.除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
//3.如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
//4.如果没有 default，那么 代码块会被阻塞，直到有一个 case 通过评估；否则一直阻塞;`))
	}

}

//1.除 default 外，如果只有一个 case 语句评估通过，那么就执行这个case里的语句；
//2.除 default 外，如果有多个 case 语句评估通过，那么通过伪随机的方式随机选一个；
//3.如果 default 外的 case 语句都没有通过评估，那么执行 default 里的语句；
//4.如果没有 default，那么 代码块会被阻塞，直到有一个 case 通过评估；否则一直阻塞;
func TestSelect(t *testing.T) {

	var c = make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i, ok := <-c //reciver ready ok
	if ok {

	}
	log.Printf("channel value:%d", i)

	//blocking test for chan
	done := make(chan bool)
	go func() {
		log.Printf("and wait 3 seconds>>>>")
		time.Sleep(3 * time.Second)
		log.Printf("notify continue>>>>")
		done <- true
	}()
	log.Printf("blocking>>>>>>>>>")
	<-done //no data ,reciver will blocking... ...
	log.Printf("got notify ,wait done,continue>>>>>>>>>")

	//ticker  loop
	timeOutCurrentTime := time.Tick(time.Second * 3)
	for {
		select {
		case x := <-timeOutCurrentTime:

			log.Printf("too much time: %s", time.Unix(int64(x.Unix()), 0).
				Format("2006-01-02 15:04:05.000000"))
			//if flg == 0 {
			//	//time_out_current_time = time.After(time.Second * 3)
			//	flg = 1
			//}
			//default:
			//	log.Printf("do nothing")
			//	time.Sleep(time.Second * 1)
		}
		//if flg == 1 {
		//	log.Printf("select case break ,reset flg")
		//	flg = 0
		//	break
		//}

	}

}
