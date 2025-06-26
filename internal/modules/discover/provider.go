package discover

import (
	"github.com/google/wire"
	"thmanyah/internal/modules/discover/biz"
	"thmanyah/internal/modules/discover/data"
	"thmanyah/internal/modules/discover/data/cache"
	"thmanyah/internal/modules/discover/service"
)

var ProviderSet = wire.NewSet(
	biz.NewDiscoverUsecase,
	repo.NewDiscoverRepo,
	cache.NewMemoryCache,
	service.NewDiscoverService,
)
