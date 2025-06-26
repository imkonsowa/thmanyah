package service

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	khttp "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/errors"
	v1 "thmanyah/api/grpc/v1"
	"thmanyah/internal/modules/cms/biz"
	"thmanyah/internal/utils"
	"thmanyah/internal/utils/convert"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CmsService struct {
	v1.UnimplementedCmsServiceServer

	uc *biz.UseCase
}

func NewCmsService(uc *biz.UseCase) *CmsService {
	return &CmsService{uc: uc}
}

// Program operations

func (s *CmsService) CreateProgram(ctx context.Context, req *v1.CreateProgramRequest) (*v1.CreateProgramResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	program := &biz.Program{
		Title:        req.Title,
		Description:  req.Description,
		CategoryID:   categoryID,
		Status:       biz.ProgramStatusDraft,
		ThumbnailURL: req.ThumbnailUrl,
		Tags:         req.Tags,
		Metadata:     convert.ConvertMetadata(req.Metadata),
		IsFeatured:   req.IsFeatured,
		CreatedBy:    userID,
		UpdatedBy:    userID,
	}

	if req.SourceUrl != "" {
		sourceUrl := req.SourceUrl
		program.SourceURL = &sourceUrl
	}

	err = s.uc.CreateProgram(ctx, program)
	if err != nil {
		return nil, err
	}

	// Fetch the created program to return it with all fields populated
	createdProgram, err := s.uc.GetProgram(ctx, program.ID)
	if err != nil {
		return nil, err
	}

	return &v1.CreateProgramResponse{
		Program: convert.ConvertProgram(createdProgram),
	}, nil
}

func (s *CmsService) UpdateProgram(ctx context.Context, req *v1.UpdateProgramRequest) (*v1.UpdateProgramResponse, error) {
	programID, err := uuid.Parse(req.ProgramId)
	if err != nil {
		return nil, err
	}

	// Build safe update request
	updates := &biz.UpdateProgramRequest{}

	if req.Title != nil {
		updates.Title = req.Title
	}
	if req.Description != nil {
		updates.Description = req.Description
	}
	if req.CategoryId != nil {
		categoryID, err := uuid.Parse(*req.CategoryId)
		if err != nil {
			return nil, err
		}
		updates.CategoryID = &categoryID
	}
	if req.Status != nil {
		if bizStatus, exists := convert.ProtoToBizProgramStatus[*req.Status]; exists {
			updates.Status = &bizStatus
		} else {
			return nil, fmt.Errorf("invalid program status: %v", *req.Status)
		}
	}
	if req.ThumbnailUrl != nil {
		updates.ThumbnailURL = req.ThumbnailUrl
	}
	if len(req.Tags) > 0 {
		updates.Tags = &req.Tags
	}
	if len(req.Metadata) > 0 {
		metadata := convert.ConvertMetadata(req.Metadata)
		updates.Metadata = &metadata
	}
	if req.SourceUrl != nil {
		updates.SourceURL = req.SourceUrl
	}
	if req.IsFeatured != nil {
		updates.IsFeatured = req.IsFeatured
	}

	program, err := s.uc.UpdateProgram(ctx, programID, updates)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateProgramResponse{
		Program: convert.ConvertProgram(program),
	}, nil
}

func (s *CmsService) DeleteProgram(ctx context.Context, req *v1.DeleteProgramRequest) (*emptypb.Empty, error) {
	programID, err := uuid.Parse(req.ProgramId)
	if err != nil {
		return nil, err
	}

	err = s.uc.DeleteProgram(ctx, programID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CmsService) GetProgram(ctx context.Context, req *v1.GetProgramRequest) (*v1.GetProgramResponse, error) {
	programID, err := uuid.Parse(req.ProgramId)
	if err != nil {
		return nil, err
	}

	program, err := s.uc.GetProgram(ctx, programID)
	if err != nil {
		return nil, err
	}

	return &v1.GetProgramResponse{
		Program: convert.ConvertProgram(program),
	}, nil
}

func (s *CmsService) ListPrograms(ctx context.Context, req *v1.ListProgramsRequest) (*v1.ListProgramsResponse, error) {
	filter := biz.ProgramFilter{}

	if req.CategoryId != "" {
		categoryID, err := uuid.Parse(req.CategoryId)
		if err != nil {
			return nil, err
		}
		filter.CategoryID = &categoryID
	}

	if req.Status != 0 {
		if bizStatus, exists := convert.ProtoToBizProgramStatus[req.Status]; exists {
			filter.Status = &bizStatus
		} else {
			return nil, fmt.Errorf("invalid program status: %v", req.Status)
		}
	}

	if req.SearchQuery != "" {
		filter.SearchQuery = &req.SearchQuery
	}

	if len(req.Tags) > 0 {
		filter.Tags = req.Tags
	}

	if req.FeaturedOnly {
		filter.FeaturedOnly = &req.FeaturedOnly
	}

	pagination := biz.PaginationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	sort := biz.SortRequest{
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}

	programs, paginationResp, err := s.uc.ListPrograms(ctx, filter, pagination, sort)
	if err != nil {
		return nil, err
	}

	return &v1.ListProgramsResponse{
		Programs:   convert.ConvertPrograms(programs),
		TotalCount: paginationResp.TotalCount,
		Page:       paginationResp.Page,
		PageSize:   paginationResp.PageSize,
	}, nil
}

func (s *CmsService) BulkUpdatePrograms(ctx context.Context, req *v1.BulkUpdateProgramsRequest) (*v1.BulkUpdateProgramsResponse, error) {
	programIDs := make([]uuid.UUID, len(req.ProgramIds))
	for i, idStr := range req.ProgramIds {
		programID, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		programIDs[i] = programID
	}

	updates := &biz.BulkUpdateProgramsRequest{}

	if req.Status != 0 {
		if bizStatus, exists := convert.ProtoToBizProgramStatus[req.Status]; exists {
			updates.Status = &bizStatus
		} else {
			return nil, fmt.Errorf("invalid program status: %v", req.Status)
		}
	}
	if req.CategoryId != "" {
		categoryID, err := uuid.Parse(req.CategoryId)
		if err != nil {
			return nil, err
		}
		updates.CategoryID = &categoryID
	}
	if len(req.Tags) > 0 {
		updates.Tags = &req.Tags
	}
	if len(req.Metadata) > 0 {
		metadata := convert.ConvertMetadata(req.Metadata)
		updates.Metadata = &metadata
	}
	updates.IsFeatured = &req.IsFeatured

	updatedCount, err := s.uc.BulkUpdatePrograms(ctx, programIDs, updates)
	if err != nil {
		return nil, err
	}

	return &v1.BulkUpdateProgramsResponse{
		UpdatedCount: updatedCount,
		FailedCount:  int32(len(req.ProgramIds)) - updatedCount,
	}, nil
}

func (s *CmsService) BulkDeletePrograms(ctx context.Context, req *v1.BulkDeleteProgramsRequest) (*emptypb.Empty, error) {
	programIDs := make([]uuid.UUID, len(req.ProgramIds))
	for i, idStr := range req.ProgramIds {
		programID, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		programIDs[i] = programID
	}

	err := s.uc.BulkDeletePrograms(ctx, programIDs)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// Category operations

func (s *CmsService) CreateCategory(ctx context.Context, req *v1.CreateCategoryRequest) (*v1.CreateCategoryResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	categoryType := biz.CategoryTypePodcast // Default to PODCAST
	if req.Type != 0 {
		if bizType, exists := convert.ProtoToBizCategoryType[req.Type]; exists {
			categoryType = bizType
		} else {
			return nil, fmt.Errorf("invalid category type: %v", req.Type)
		}
	}

	category := &biz.Category{
		Name:        req.Name,
		Description: req.Description,
		Type:        categoryType,
		Metadata:    convert.ConvertMetadata(req.Metadata),
		CreatedBy:   userID,
	}

	err = s.uc.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	// Fetch the created category to return it with all fields populated
	createdCategory, err := s.uc.GetCategory(ctx, category.ID)
	if err != nil {
		return nil, err
	}

	return &v1.CreateCategoryResponse{
		Category: convert.ConvertCategory(createdCategory),
	}, nil
}

func (s *CmsService) UpdateCategory(ctx context.Context, req *v1.UpdateCategoryRequest) (*v1.UpdateCategoryResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	// Build safe update request
	updates := &biz.UpdateCategoryRequest{}

	if req.Name != nil {
		updates.Name = req.Name
	}
	if req.Description != nil {
		updates.Description = req.Description
	}
	if req.Type != nil {
		if bizType, exists := convert.ProtoToBizCategoryType[*req.Type]; exists {
			updates.Type = &bizType
		} else {
			return nil, fmt.Errorf("invalid category type: %v", *req.Type)
		}
	}
	if len(req.Metadata) > 0 {
		metadata := convert.ConvertMetadata(req.Metadata)
		updates.Metadata = &metadata
	}

	category, err := s.uc.UpdateCategory(ctx, categoryID, updates)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateCategoryResponse{
		Category: convert.ConvertCategory(category),
	}, nil
}

func (s *CmsService) DeleteCategory(ctx context.Context, req *v1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	err = s.uc.DeleteCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CmsService) GetCategory(ctx context.Context, req *v1.GetCategoryRequest) (*v1.GetCategoryResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	category, err := s.uc.GetCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return &v1.GetCategoryResponse{
		Category: convert.ConvertCategory(category),
	}, nil
}

func (s *CmsService) ListCategories(ctx context.Context, req *v1.ListCategoriesRequest) (*v1.ListCategoriesResponse, error) {
	filter := biz.CategoryFilter{}

	// Only apply type filter if explicitly provided in the request
	if req.Type != 0 {
		if bizType, exists := convert.ProtoToBizCategoryType[req.Type]; exists {
			filter.Type = &bizType
		} else {
			return nil, fmt.Errorf("invalid category type: %v", req.Type)
		}
	}

	if req.SearchQuery != "" {
		filter.SearchQuery = &req.SearchQuery
	}

	pagination := biz.PaginationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	sort := biz.SortRequest{
		SortBy:    "name",
		SortOrder: "asc",
	}

	categories, paginationResp, err := s.uc.ListCategories(ctx, filter, pagination, sort)
	if err != nil {
		return nil, err
	}

	return &v1.ListCategoriesResponse{
		Categories: convert.ConvertCategories(categories),
		TotalCount: paginationResp.TotalCount,
		Page:       paginationResp.Page,
		PageSize:   paginationResp.PageSize,
	}, nil
}

// Episode operations

func (s *CmsService) CreateEpisode(ctx context.Context, req *v1.CreateEpisodeRequest) (*v1.CreateEpisodeResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	programID, err := uuid.Parse(req.ProgramId)
	if err != nil {
		return nil, err
	}

	episode := &biz.Episode{
		ProgramID:     programID,
		Title:         req.Title,
		Description:   req.Description,
		DurationSecs:  req.DurationSeconds,
		EpisodeNumber: req.EpisodeNumber,
		SeasonNumber:  req.SeasonNumber,
		MediaURL:      req.MediaUrl,
		ThumbnailURL:  req.ThumbnailUrl,
		Tags:          req.Tags,
		Metadata:      convert.ConvertMetadata(req.Metadata),
		CreatedBy:     userID,
		UpdatedBy:     userID,
		Status:        biz.EpisodeStatusDraft,
	}

	err = s.uc.CreateEpisode(ctx, episode)
	if err != nil {
		return nil, err
	}

	// Fetch the created episode to return it with all fields populated
	createdEpisode, err := s.uc.GetEpisode(ctx, episode.ID)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEpisodeResponse{
		Episode: convert.ConvertEpisode(createdEpisode),
	}, nil
}

func (s *CmsService) UpdateEpisode(ctx context.Context, req *v1.UpdateEpisodeRequest) (*v1.UpdateEpisodeResponse, error) {
	episodeID, err := uuid.Parse(req.EpisodeId)
	if err != nil {
		return nil, err
	}

	// Build safe update request
	updates := &biz.UpdateEpisodeRequest{}

	if req.Title != nil {
		updates.Title = req.Title
	}
	if req.Description != nil {
		updates.Description = req.Description
	}
	if req.DurationSeconds != nil {
		updates.DurationSecs = req.DurationSeconds
	}
	if req.EpisodeNumber != nil {
		updates.EpisodeNumber = req.EpisodeNumber
	}
	if req.SeasonNumber != nil {
		updates.SeasonNumber = req.SeasonNumber
	}
	if req.Status != nil {
		if bizStatus, exists := convert.ProtoToBizEpisodeStatus[*req.Status]; exists {
			updates.Status = &bizStatus
		} else {
			return nil, fmt.Errorf("invalid episode status: %v", *req.Status)
		}
	}
	if req.MediaUrl != nil {
		updates.MediaURL = req.MediaUrl
	}
	if req.ThumbnailUrl != nil {
		updates.ThumbnailURL = req.ThumbnailUrl
	}
	if len(req.Tags) > 0 {
		updates.Tags = &req.Tags
	}
	if len(req.Metadata) > 0 {
		metadata := convert.ConvertMetadata(req.Metadata)
		updates.Metadata = &metadata
	}
	if req.ScheduledAt != nil {
		scheduledAt := req.ScheduledAt.AsTime()
		updates.ScheduledAt = &scheduledAt
	}

	episode, err := s.uc.UpdateEpisode(ctx, episodeID, updates)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateEpisodeResponse{
		Episode: convert.ConvertEpisode(episode),
	}, nil
}

func (s *CmsService) DeleteEpisode(ctx context.Context, req *v1.DeleteEpisodeRequest) (*emptypb.Empty, error) {
	episodeID, err := uuid.Parse(req.EpisodeId)
	if err != nil {
		return nil, err
	}

	err = s.uc.DeleteEpisode(ctx, episodeID)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CmsService) GetEpisode(ctx context.Context, req *v1.GetEpisodeRequest) (*v1.GetEpisodeResponse, error) {
	episodeID, err := uuid.Parse(req.EpisodeId)
	if err != nil {
		return nil, err
	}

	episode, err := s.uc.GetEpisode(ctx, episodeID)
	if err != nil {
		return nil, err
	}

	return &v1.GetEpisodeResponse{
		Episode: convert.ConvertEpisode(episode),
	}, nil
}

func (s *CmsService) ListEpisodes(ctx context.Context, req *v1.ListEpisodesRequest) (*v1.ListEpisodesResponse, error) {
	programID, err := uuid.Parse(req.ProgramId)
	if err != nil {
		return nil, err
	}

	filter := biz.EpisodeFilter{
		ProgramID: &programID,
	}

	// Only apply status filter if explicitly provided in the request
	if req.Status != 0 {
		if bizStatus, exists := convert.ProtoToBizEpisodeStatus[req.Status]; exists {
			filter.Status = &bizStatus
		} else {
			return nil, fmt.Errorf("invalid episode status: %v", req.Status)
		}
	}

	if req.SearchQuery != "" {
		filter.SearchQuery = &req.SearchQuery
	}

	pagination := biz.PaginationRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	sort := biz.SortRequest{
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}

	episodes, paginationResp, err := s.uc.ListEpisodes(ctx, filter, pagination, sort)
	if err != nil {
		return nil, err
	}

	return &v1.ListEpisodesResponse{
		Episodes:   convert.ConvertEpisodes(episodes),
		TotalCount: paginationResp.TotalCount,
		Page:       paginationResp.Page,
		PageSize:   paginationResp.PageSize,
	}, nil
}

// Import operations

func (s *CmsService) ImportData(ctx context.Context, req *v1.ImportDataRequest) (*v1.ImportDataResponse, error) {
	userID, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	categoryID, err := uuid.Parse(req.DefaultCategoryId)
	if err != nil {
		return nil, err
	}

	importData := &biz.ImportData{
		SourceType:   req.SourceType,
		SourceURL:    req.SourceUrl,
		SourceConfig: convert.ConvertMetadata(req.SourceConfig),
		CategoryID:   categoryID,
		FieldMapping: convert.ConvertMetadata(req.FieldMapping),
		CreatedBy:    userID,
	}

	result, err := s.uc.ImportData(ctx, importData)
	if err != nil {
		return nil, err
	}

	return &v1.ImportDataResponse{
		ImportId:       result.ID.String(),
		Status:         convert.BizImportStatusToProto(result.Status),
		TotalItems:     result.TotalItems,
		ProcessedItems: result.ProcessedItems,
		SuccessCount:   result.SuccessCount,
		ErrorCount:     result.ErrorCount,
		Warnings:       result.Warnings,
	}, nil
}

func (s *AuthService) UpdateAvatar(ctx khttp.Context, userId uuid.UUID) (*v1.EpisodeFileUpdateResponse, error) {
	file, header, err := ctx.Request().FormFile("file")
	if err != nil {
		return nil, err
	}

	target := ctx.Request().FormValue("target")
	episodeId, err := uuid.Parse(ctx.Request().FormValue("episode_id"))
	if err != nil {
		return nil, err
	}

	if !slices.Contains([]string{"thumbnail", "media"}, target) {
		return nil, errors.Newf(http.StatusBadRequest, "INVALID_TARGET", "INVALID_TARGET")
	}

	fileUrl, err := s.uc.UpdateEpisodeFile(ctx, userId, episodeId, &biz.UpdateEpisodeFileRequest{
		Header: header,
		File:   file,
		Target: target,
	})
	if err != nil {
		return nil, err
	}

	return &v1.EpisodeFileUpdateResponse{
		FileUrl: fileUrl,
	}, nil
}
