package go2cache


type Cache interface {
	//从缓存中获取数据
	Get(key string) interface{}
	//将数据放入缓存
	//key :键  value：存储数据值
	Put(key string, value interface{}) error
	//删除
	Delete(key string) error
	//当前键值是否存在
	IsExist(key string) bool
}

//// cache 缓存元数据
//type CacheObject struct {
//	//存储值
//	Value interface{}
//	//value值过期时间
//	DeadLineTime time.Time
//	//expired
//	Seconds float64
//}
//
////当前对象是否已经过期
//func (object *CacheObject) IsExpired() bool {
//	return time.Now().After(object.DeadLineTime)
//}
