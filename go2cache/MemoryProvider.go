package go2cache

import (
	"sync"
)

var mmp = &MemoryProvider{memoryCacheRegion: make(map[string]*MemoryCache)}

//基于redis缓存提供者
type MemoryProvider struct {
	//同步 MemoryCache map
	mu sync.RWMutex
	//基于map的region 和cache的封装
	memoryCacheRegion map[string]*MemoryCache
	//regins
	regions []Region
}

//构建region cache
func (p *MemoryProvider) BuildCache(region string) (interface{}, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	cache := mmp.memoryCacheRegion[region]
	if cache == nil {
		cache = BuildMemoryCache()
		mmp.memoryCacheRegion[region] = cache
		p.regions = append(p.regions, Region{Name: region})
	}
	return cache, nil
}

//获取region 列表
func (p *MemoryProvider) GetRegions() []Region {
	return p.regions
}

func (p *MemoryProvider) Name() string {
	return "memory_provider"
}

func (p *MemoryProvider) Level() int {
	return Level1
}
