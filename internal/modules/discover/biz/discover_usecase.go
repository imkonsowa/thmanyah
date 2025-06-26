package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	cms "thmanyah/internal/modules/cms/biz"
)

type DiscoverUsecase struct {
	searchRepo DiscoverRepository
	cache      MemoryCache
	logger     *log.Helper
}

// CachedSearchResult holds search results with total count for caching
type CachedSearchResult struct {
	Results    *SearchResults
	TotalCount int32
}

func NewDiscoverUsecase(searchRepo DiscoverRepository, cache MemoryCache, logger log.Logger) *DiscoverUsecase {
	return &DiscoverUsecase{
		searchRepo: searchRepo,
		cache:      cache,
		logger:     log.NewHelper(logger),
	}
}

func (d *DiscoverUsecase) Search(ctx context.Context, query string, page, pageSize int32) (*SearchResults, int32, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// Generate cache key based on query parameters
	cacheKey := fmt.Sprintf("search:%s:page:%d:size:%d", query, page, pageSize)
	const cacheTTL = 5 * time.Minute // Cache for 5 minutes

	if cached, found := d.cache.Get(cacheKey); found {
		if cachedResult, ok := cached.(*CachedSearchResult); ok {
			return cachedResult.Results, cachedResult.TotalCount, nil
		}
		d.logger.Warn("Invalid cached data type for search results")
	}

	results, totalCount, err := d.searchRepo.Search(ctx, query, page, pageSize)
	if err != nil {
		d.logger.Errorf("Failed to search: %v", err)
		return nil, 0, err
	}

	cachedResult := &CachedSearchResult{
		Results:    results,
		TotalCount: totalCount,
	}
	if !d.cache.SetWithTTL(cacheKey, cachedResult, 1, cacheTTL) {
		d.logger.Warnf("Failed to cache search results for query: %s", query)
	}

	return results, totalCount, nil
}

func (d *DiscoverUsecase) Featured(ctx context.Context) ([]*cms.Program, error) {
	const cacheKey = "featured_programs"
	const cacheTTL = 15 * time.Minute // Cache for 15 minutes

	if cached, found := d.cache.Get(cacheKey); found {
		if programs, ok := cached.([]*cms.Program); ok {
			return programs, nil
		}

		d.logger.Warn("Invalid cached data type for featured programs")
	}

	programs, err := d.searchRepo.Featured(ctx)
	if err != nil {
		d.logger.Errorf("Failed to get featured programs: %v", err)
		return nil, err
	}

	if !d.cache.SetWithTTL(cacheKey, programs, 1, cacheTTL) {
		d.logger.Warn("Failed to cache featured programs")
	}

	return programs, nil
}
