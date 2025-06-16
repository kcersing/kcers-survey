package cache

import (
	"github.com/dgraph-io/ristretto"
	"sync"
)

var onceCache sync.Once

var Cache *ristretto.Cache

func InitCache() {
	onceCache.Do(func() {
		Cache = getCache()
	})
}

func getCache() *ristretto.Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters:        1e7,     // Num keys to track frequency of (10M).
		MaxCost:            1 << 30, // Maximum cost of cache (1GB).
		BufferItems:        64,      // Number of keys per Get buffer.
		IgnoreInternalCost: true,
	})
	if err != nil {
		panic(any(err))

	}
	return cache
}
