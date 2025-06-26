package cms

import (
	"github.com/google/wire"
	"thmanyah/internal/modules/cms/biz"
	"thmanyah/internal/modules/cms/data/repo"
	"thmanyah/internal/modules/cms/data/s3"
	"thmanyah/internal/modules/cms/service"
)

var ProviderSet = wire.NewSet(
	// data layer dependencies
	repo.NewUsersRepo,
	repo.NewCategoryRepository,
	repo.NewProgramRepository,
	repo.NewEpisodeRepository,
	repo.NewImportRepository,
	s3.NewS3Client,

	// biz layer dependencies
	biz.NewUseCase,

	// network layer dependencies
	service.NewAuthService,
	service.NewCmsService,
)
