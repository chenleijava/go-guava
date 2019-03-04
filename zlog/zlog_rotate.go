package zlog

import (
	"github.com/robfig/cron"
	"gopkg.in/natefinch/lumberjack.v2"
	"sync"
)

//log rotate
type LogRotate struct {
	logger    *lumberjack.Logger
	c         *cron.Cron
	writeLock sync.Mutex
}

//get log rotate instance
func NewLogRotate(fileName string, maxSize, maxBackups, days int, spec string) *LogRotate {
	instance := new(LogRotate)
	instance.logger = NewLog2FileByLumberJackLog(fileName, maxSize, maxBackups, days)
	instance.c = cron.New()
	if spec == "" {
		_ = instance.c.AddFunc("58 59 23 * * ?", instance.rotate) //default 23:59:58 rotate log file
	} else {
		_ = instance.c.AddFunc(spec, instance.rotate)
	}
	instance.c.Start()
	return instance
}

//rote
func (l *LogRotate) rotate() {
	l.writeLock.Lock()    //lock
	_ = l.logger.Rotate() //rotate
	l.writeLock.Unlock()  //  un lock
}

//write data
func (l *LogRotate) Write(d []byte) {
	l.writeLock.Lock()                  //lock
	_, _ = l.logger.Write(d)            // write logger data
	_, _ = l.logger.Write([]byte("\n")) // write file '\n'
	l.writeLock.Unlock()                //  un lock
}
