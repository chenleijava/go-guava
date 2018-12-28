package zlog

import (
	"github.com/Jeffail/tunny"
	"github.com/chenleijava/go-guava/goroutine-pool"
	"github.com/chenleijava/go-guava/rabbitmq/p"
	"github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"runtime"
	"sync"
	"time"
)

//zap FAQ
//https://github.com/uber-go/zap/blob/master/FAQ.md#does-zap-support-log-rotation
//fileName : outPath file;
//maxSize  : megabytes ;
//MaxBackups: number of log files   default is not removed
//MaxAge: days
//production level :info, warn ,error
//addCaller: call line number
//omitted   key
func NewLog2File(fileName string, addCaller bool, maxSize, maxBackups, days int) *zap.Logger {
	lumberjackLog := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,    // megabytes  M
		MaxBackups: maxBackups, // number of log files   default is not removed
		MaxAge:     days,       // days  default is not removed
		Compress:   true,       // disabled by default
	}
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000000"))
	}
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(lumberjackLog),
		zap.InfoLevel,
	)
	if addCaller {
		return zap.New(core, zap.AddCaller())
	} else {
		return zap.New(core)
	}
}

//std console
//omitted   key
func NewLog2Console() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000000"))
	}
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		os.Stdout,
		zapcore.InfoLevel,
	)
	return zap.New(core, zap.AddCaller())
}

//direct write data to file
func NewLog2FileByLumberJackLog(fileName string, maxSize, maxBackups, days int) *lumberjack.Logger {
	lumberjackLog := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,    // megabytes  M
		MaxBackups: maxBackups, // number of log files   default is not removed
		MaxAge:     days,       // days  default is not removed
		Compress:   true,       // disabled by default
	}
	return lumberjackLog
}

//base on mq
var onceMqWriterSyncerOnce sync.Once
var mqSyncer *RabbitMqWriteSyncer

//Get rabbitMqWriteSysncer
func GetRabbitMqWriteSyncer() *RabbitMqWriteSyncer {
	if mqSyncer == nil {
		onceMqWriterSyncerOnce.Do(func() {
			mqSyncer = new(RabbitMqWriteSyncer)
		})
	}
	return mqSyncer
}

//log flg  below to which log file
const LogFlg = "log_flg"

//Get log flg for consumer
//log.info("",zap.string(LogFlg,"logflg"))=> call  AddLogFlg method
func GetLogFlg(data *[]byte) string {
	return jsoniter.Get(*data, LogFlg).ToString()
}

//add log_flg for producer
func AddLogFlg(logFlg string) *zap.Field {
	f := zap.String(LogFlg, logFlg)
	return &f
}

//
type RabbitMqWriteSyncer struct {
	//
	mqTemplate *p.Producer
	//
	tunnyPool  *tunny.Pool
}

//write to rabbitmq
func (mqSyncer *RabbitMqWriteSyncer) Write(p []byte) (int, error) {
	e := mqSyncer.tunnyPool.Process(&p) // address , pointer type: *[]byte
	l := len(p)
	//set nil for gc
	p = nil
	if e != nil { // error not nill
		return l, e.(error)
	} else {
		return l, nil
	}
}

func (mqSyncer *RabbitMqWriteSyncer) Sync() error {
	// nothing to sync for mq
	return nil
}

//log server  mq default config
const (
	exchangeName = "log_exchange"
	QueueName    = "log_queue"
	routeKey     = "log_route_key"
)

//init mq sendTemplate
func (mqSyncer *RabbitMqWriteSyncer) InitRabbitMqWriteSyncerDefault(url string) {
	mqSyncer.InitRabbitMqWriteSyncer(exchangeName, QueueName, routeKey, url)
}

//init mq sendTemplate
func (mqSyncer *RabbitMqWriteSyncer) InitRabbitMqWriteSyncer(exchangeName, queueName, routeKey, url string) {
	//init sync goroutine pool to handle msg send
	//goroutine pool size equal to cpu number
	//payload process pass arg (
	mqSyncer.tunnyPool =
		goroutine_pool.NewSyncPool(runtime.NumCPU(), func(payload interface{}) interface{} {
			data := (payload).(*[]byte)           //
			return mqSyncer.mqTemplate.Send(data) //
		})

	//init mq syncer template
	mqSyncer.mqTemplate = &p.Producer{
		ExchangeName: exchangeName,
		QueueName:    queueName,
		RouteKey:     routeKey,
	}
	p.Producers = append(p.Producers, mqSyncer.mqTemplate)
	//connect rabbitmq ?!
	p.RabbitmqConn(url)
}

//send data to mq
func (mqSyncer *RabbitMqWriteSyncer) NewLog2RabbitMq(logFlgField *zap.Field) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000000"))
	}
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(mqSyncer),
		zap.InfoLevel,
	)
	if logFlgField != nil {
		return zap.New(core, zap.AddCaller(), zap.Fields(*logFlgField))
	} else {
		return zap.New(core, zap.AddCaller())
	}
}
