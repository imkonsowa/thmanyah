package biz

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"thmanyah/internal/utils"
	"thmanyah/keys"
)

type UseCase struct {
	logger    *log.Helper
	keysStore *keys.Store

	usersRepo    UsersRepository
	categoryRepo CategoryRepository
	programRepo  ProgramRepository
	episodeRepo  EpisodeRepository
	importRepo   ImportRepository
	s3           S3Client
}

func NewUseCase(
	userRepo UsersRepository,
	categoryRepo CategoryRepository,
	programRepo ProgramRepository,
	episodeRepo EpisodeRepository,
	importRepo ImportRepository,
	keysStore *keys.Store,
	s3 S3Client,
	logger log.Logger,
) *UseCase {
	return &UseCase{
		logger:       log.NewHelper(logger),
		usersRepo:    userRepo,
		categoryRepo: categoryRepo,
		programRepo:  programRepo,
		episodeRepo:  episodeRepo,
		importRepo:   importRepo,
		keysStore:    keysStore,
		s3:           s3,
	}
}

// Program operations

func (uc *UseCase) CreateProgram(ctx context.Context, program *Program) error {
	return uc.programRepo.Create(ctx, program)
}

func (uc *UseCase) UpdateProgram(ctx context.Context, id uuid.UUID, updates *UpdateProgramRequest) (*Program, error) {
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		program, err := uc.programRepo.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}
		if program.CreatedBy != userID {
			return nil, fmt.Errorf("unauthorized: you can only update your own programs")
		}
	}

	return uc.programRepo.Update(ctx, userID, id, updates)
}

func (uc *UseCase) DeleteProgram(ctx context.Context, id uuid.UUID) error {
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.programRepo.Delete(ctx, userID, id)
}

func (uc *UseCase) GetProgram(ctx context.Context, id uuid.UUID) (*Program, error) {
	return uc.programRepo.GetByID(ctx, id)
}

func (uc *UseCase) ListPrograms(ctx context.Context, filter ProgramFilter, pagination PaginationRequest, sort SortRequest) ([]*Program, *PaginationResponse, error) {
	pagination.SetDefaults()

	return uc.programRepo.List(ctx, filter, pagination, sort)
}

func (uc *UseCase) BulkUpdatePrograms(ctx context.Context, ids []uuid.UUID, updates *BulkUpdateProgramsRequest) (int32, error) {
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return 0, fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.programRepo.BulkUpdate(ctx, userID, ids, updates)
}

func (uc *UseCase) BulkDeletePrograms(ctx context.Context, ids []uuid.UUID) error {
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.programRepo.BulkDelete(ctx, userID, ids)
}

// Category operations

func (uc *UseCase) CreateCategory(ctx context.Context, category *Category) error {
	if category.ID == uuid.Nil {
		category.ID = uuid.Must(uuid.NewV7())
	}

	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		category.CreatedBy = userID
	}

	return uc.categoryRepo.Create(ctx, category)
}

func (uc *UseCase) UpdateCategory(ctx context.Context, id uuid.UUID, updates *UpdateCategoryRequest) (*Category, error) {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return nil, fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.categoryRepo.Update(ctx, userID, id, updates)
}

func (uc *UseCase) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.categoryRepo.Delete(ctx, userID, id)
}

func (uc *UseCase) GetCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	return uc.categoryRepo.GetByID(ctx, id)
}

func (uc *UseCase) ListCategories(ctx context.Context, filter CategoryFilter, pagination PaginationRequest, sort SortRequest) ([]*Category, *PaginationResponse, error) {
	pagination.SetDefaults()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		// Filter to only show user's own categories
		filter.CreatedBy = &userID
	}

	return uc.categoryRepo.List(ctx, filter, pagination, sort)
}

// Episode operations

func (uc *UseCase) CreateEpisode(ctx context.Context, episode *Episode) error {
	if episode.ID == uuid.Nil {
		episode.ID = uuid.Must(uuid.NewV7())
	}

	episode.CreatedAt = time.Now()
	episode.UpdatedAt = time.Now()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		episode.CreatedBy = userID
		episode.UpdatedBy = userID
	}

	err := uc.episodeRepo.Create(ctx, episode)
	if err != nil {
		return err
	}

	// Update episodes count for the program
	return uc.programRepo.UpdateEpisodesCount(ctx, episode.ProgramID)
}

func (uc *UseCase) UpdateEpisode(ctx context.Context, id uuid.UUID, updates *UpdateEpisodeRequest) (*Episode, error) {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return nil, fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.episodeRepo.Update(ctx, userID, id, updates)
}

func (uc *UseCase) DeleteEpisode(ctx context.Context, id uuid.UUID) error {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.episodeRepo.Delete(ctx, userID, id)
}

func (uc *UseCase) GetEpisode(ctx context.Context, id uuid.UUID) (*Episode, error) {
	return uc.episodeRepo.GetByID(ctx, id)
}

func (uc *UseCase) ListEpisodes(ctx context.Context, filter EpisodeFilter, pagination PaginationRequest, sort SortRequest) ([]*Episode, *PaginationResponse, error) {
	pagination.SetDefaults()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		// Filter to only show user's own episodes
		filter.CreatedBy = &userID
	}

	return uc.episodeRepo.List(ctx, filter, pagination, sort)
}

func (uc *UseCase) ListEpisodesByProgram(ctx context.Context, programID uuid.UUID, pagination PaginationRequest, sort SortRequest) ([]*Episode, *PaginationResponse, error) {
	pagination.SetDefaults()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		// Check if user owns this program
		program, err := uc.programRepo.GetByID(ctx, programID)
		if err != nil {
			return nil, nil, err
		}
		if program.CreatedBy != userID {
			return nil, nil, fmt.Errorf("unauthorized: you can only view episodes of your own programs")
		}
	}

	return uc.episodeRepo.ListByProgram(ctx, programID, pagination, sort)
}

func (uc *UseCase) IncrementEpisodeViewCount(ctx context.Context, id uuid.UUID) error {
	return uc.episodeRepo.IncrementViewCount(ctx, id)
}

func (uc *UseCase) IncrementProgramViewCount(ctx context.Context, id uuid.UUID) error {
	return uc.programRepo.IncrementViewCount(ctx, id)
}

// Import operations

func (uc *UseCase) ImportData(ctx context.Context, importData *ImportData) (*ImportData, error) {
	if importData.ID == uuid.Nil {
		importData.ID = uuid.Must(uuid.NewV7())
	}

	importData.CreatedAt = time.Now()
	importData.UpdatedAt = time.Now()
	importData.Status = ImportStatusPending

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		importData.CreatedBy = userID
	}

	err := uc.importRepo.Create(ctx, importData)
	if err != nil {
		return nil, err
	}

	return importData, nil
}

func (uc *UseCase) UpdateImport(ctx context.Context, id uuid.UUID, updates *UpdateImportRequest) (*ImportData, error) {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return nil, fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.importRepo.Update(ctx, userID, id, updates)
}

func (uc *UseCase) GetImport(ctx context.Context, id uuid.UUID) (*ImportData, error) {
	return uc.importRepo.GetByID(ctx, id)
}

func (uc *UseCase) ListImports(ctx context.Context, pagination PaginationRequest, sort SortRequest) ([]*ImportData, *PaginationResponse, error) {
	pagination.SetDefaults()

	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID != uuid.Nil {
		// Filter to only show user's own imports
		filter := ImportFilter{
			CreatedBy: &userID,
		}
		return uc.importRepo.List(ctx, filter, pagination, sort)
	}

	// If no user ID, return empty list
	return uc.importRepo.List(ctx, ImportFilter{}, pagination, sort)
}

func (uc *UseCase) UpdateImportProgress(ctx context.Context, id uuid.UUID, updates *UpdateImportProgressRequest) error {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.importRepo.UpdateProgress(ctx, userID, id, updates)
}

func (uc *UseCase) AddImportError(ctx context.Context, id uuid.UUID, errorMsg string) error {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.importRepo.AddError(ctx, userID, id, errorMsg)
}

func (uc *UseCase) AddImportWarning(ctx context.Context, id uuid.UUID, warningMsg string) error {
	// Get user ID from context if available
	userID, _ := utils.GetUserID(ctx)
	if userID == uuid.Nil {
		return fmt.Errorf("unauthorized: user ID not found in context")
	}

	return uc.importRepo.AddWarning(ctx, userID, id, warningMsg)
}

func (uc *UseCase) UpdateEpisodeFile(ctx context.Context, userId, episodeId uuid.UUID, request *UpdateEpisodeFileRequest) (string, error) {
	episode, err := uc.episodeRepo.GetByID(ctx, episodeId)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("episodes/%s/%s%s", episode.ID.String(), request.Target, filepath.Ext(request.Header.Filename))
	err = uc.s3.PutObject(ctx, "thmanyah", key, request.File)
	if err != nil {
		return "", err
	}

	fileUrl := "/thmanyah/" + key
	updateEpisodeRequest := &UpdateEpisodeRequest{}
	if request.Target == "thumbnail" {
		updateEpisodeRequest.ThumbnailURL = &fileUrl
	}
	if request.Target == "media" {
		updateEpisodeRequest.MediaURL = &fileUrl
	}

	_, err = uc.episodeRepo.Update(ctx, userId, episodeId, updateEpisodeRequest)
	if err != nil {
		return "", err
	}

	return uc.s3.GetObjectPublicURL(ctx, "thmanyah", key), nil
}
