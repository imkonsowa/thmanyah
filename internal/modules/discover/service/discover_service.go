package service

import (
	"context"

	pb "thmanyah/api/grpc/v1"
	cms "thmanyah/internal/modules/cms/biz"
	"thmanyah/internal/modules/discover/biz"
	"thmanyah/internal/utils/convert"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DiscoverService struct {
	pb.UnimplementedDiscoverServiceServer

	discoverUc *biz.DiscoverUsecase
	logger     *log.Helper
}

func NewDiscoverService(discoverUc *biz.DiscoverUsecase, logger log.Logger) *DiscoverService {
	return &DiscoverService{
		discoverUc: discoverUc,
		logger:     log.NewHelper(logger),
	}
}

func (s *DiscoverService) Featured(ctx context.Context, req *pb.FeaturedRequest) (*pb.FeaturedResponse, error) {
	programs, err := s.discoverUc.Featured(ctx)
	if err != nil {
		s.logger.Errorf("Failed to get featured programs: %v", err)
		return nil, err
	}

	// Convert biz programs to protobuf messages
	pbPrograms := s.convertProgramsToPB(programs)

	return &pb.FeaturedResponse{
		Programs: pbPrograms,
	}, nil
}

func (s *DiscoverService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	if req.Query == "" {
		return &pb.SearchResponse{
			Categories: []*pb.Category{},
			Programs:   []*pb.Program{},
			Episodes:   []*pb.Episode{},
			TotalCount: 0,
			Page:       req.Page,
			PageSize:   req.PageSize,
			TotalPages: 0,
		}, nil
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	results, totalCount, err := s.discoverUc.Search(ctx, req.Query, page, pageSize)
	if err != nil {
		s.logger.Errorf("Failed to search: %v", err)
		return nil, err
	}

	// Convert biz results to protobuf messages
	pbCategories := s.convertCategoriesToPB(results.Categories)
	pbPrograms := s.convertProgramsToPB(results.Programs)
	pbEpisodes := s.convertEpisodesToPB(results.Episodes)

	totalPages := (totalCount + pageSize - 1) / pageSize

	return &pb.SearchResponse{
		Categories: pbCategories,
		Programs:   pbPrograms,
		Episodes:   pbEpisodes,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

func (s *DiscoverService) convertCategoriesToPB(categories []*cms.Category) []*pb.Category {
	pbCategories := make([]*pb.Category, len(categories))
	for i, category := range categories {
		pbCategories[i] = &pb.Category{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
			Type:        convert.BizCategoryTypeToProto(category.Type),
			CreatedAt:   timestamppb.New(category.CreatedAt),
			UpdatedAt:   timestamppb.New(category.UpdatedAt),
			CreatedBy:   category.CreatedBy.String(),
			Metadata:    category.Metadata,
		}
	}
	return pbCategories
}

func (s *DiscoverService) convertProgramsToPB(programs []*cms.Program) []*pb.Program {
	pbPrograms := make([]*pb.Program, len(programs))
	for i, program := range programs {
		pbProgram := &pb.Program{
			Id:            program.ID.String(),
			Title:         program.Title,
			Description:   program.Description,
			CategoryId:    program.CategoryID.String(),
			Status:        convert.BizProgramStatusToProto(program.Status),
			CreatedAt:     timestamppb.New(program.CreatedAt),
			UpdatedAt:     timestamppb.New(program.UpdatedAt),
			CreatedBy:     program.CreatedBy.String(),
			UpdatedBy:     program.UpdatedBy.String(),
			ThumbnailUrl:  program.ThumbnailURL,
			Tags:          program.Tags,
			Metadata:      program.Metadata,
			EpisodesCount: program.EpisodesCount,
			IsFeatured:    program.IsFeatured,
			ViewCount:     program.ViewCount,
			Rating:        program.Rating,
		}

		if program.PublishedAt != nil {
			pbProgram.PublishedAt = timestamppb.New(*program.PublishedAt)
		}
		if program.SourceURL != nil {
			pbProgram.SourceUrl = program.SourceURL
		}

		pbPrograms[i] = pbProgram
	}
	return pbPrograms
}

func (s *DiscoverService) convertEpisodesToPB(episodes []*cms.Episode) []*pb.Episode {
	pbEpisodes := make([]*pb.Episode, len(episodes))
	for i, episode := range episodes {
		pbEpisode := &pb.Episode{
			Id:              episode.ID.String(),
			ProgramId:       episode.ProgramID.String(),
			Title:           episode.Title,
			Description:     episode.Description,
			DurationSeconds: episode.DurationSecs,
			EpisodeNumber:   episode.EpisodeNumber,
			SeasonNumber:    episode.SeasonNumber,
			Status:          convert.BizEpisodeStatusToProto(episode.Status),
			CreatedAt:       timestamppb.New(episode.CreatedAt),
			UpdatedAt:       timestamppb.New(episode.UpdatedAt),
			CreatedBy:       episode.CreatedBy.String(),
			UpdatedBy:       episode.UpdatedBy.String(),
			MediaUrl:        episode.MediaURL,
			ThumbnailUrl:    episode.ThumbnailURL,
			Tags:            episode.Tags,
			Metadata:        episode.Metadata,
			ViewCount:       episode.ViewCount,
			Rating:          episode.Rating,
		}

		if episode.PublishedAt != nil {
			pbEpisode.PublishedAt = timestamppb.New(*episode.PublishedAt)
		}
		if episode.ScheduledAt != nil {
			pbEpisode.ScheduledAt = timestamppb.New(*episode.ScheduledAt)
		}

		pbEpisodes[i] = pbEpisode
	}
	return pbEpisodes
}
