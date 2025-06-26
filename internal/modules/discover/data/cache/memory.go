package cache

import (
	"time"

	"github.com/dgraph-io/ristretto/v2"
	"thmanyah/internal/modules/discover/biz"
)

type memory struct {
	cache *ristretto.Cache[string, any]
}

func NewMemoryCache() (biz.MemoryCache, error) {
	cache, err := ristretto.NewCache(&ristretto.Config[string, any]{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})
	if err != nil {
		return nil, err
	}

	return &memory{cache: cache}, nil
}

func (m *memory) SetWithTTL(key string, value any, cost int64, ttl time.Duration) bool {
	return m.cache.SetWithTTL(key, value, cost, ttl)
}

func (m *memory) Get(key string) (any, bool) {
	return m.cache.Get(key)
}
