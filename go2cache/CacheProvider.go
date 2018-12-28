package go2cache

const (
	// memory cache
	Level1 = iota
	//redis  cache
	Level2
)

//cache provider interface
type CacheProvider interface {
	//cacheProvider name
	Name() string
	//cache levek
	Level() int
	//build cache
	BuildCache(region string) (interface{}, error)
	// region list
	GetRegions() []Region
}
