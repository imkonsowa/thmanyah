package biz

import (
	"context"
	"time"

	cms "thmanyah/internal/modules/cms/biz"
)

type DiscoverRepository interface {
	Search(ctx context.Context, query string, page, pageSize int32) (*SearchResults, int32, error)
	Featured(ctx context.Context) ([]*cms.Program, error)
}

type MemoryCache interface {
	SetWithTTL(key string, value any, cost int64, ttl time.Duration) bool
	Get(key string) (any, bool)
}
