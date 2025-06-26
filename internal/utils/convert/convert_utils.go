package convert

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "thmanyah/api/grpc/v1"
	"thmanyah/internal/modules/cms/biz"
)

var (
	ProtoToBizCategoryType = map[v1.CategoryType]biz.CategoryType{
		v1.CategoryType_CATEGORY_TYPE_PODCAST:       biz.CategoryTypePodcast,
		v1.CategoryType_CATEGORY_TYPE_DOCUMENTARY:   biz.CategoryTypeDocumentary,
		v1.CategoryType_CATEGORY_TYPE_SPORTS_EVENT:  biz.CategoryTypeSportsEvent,
		v1.CategoryType_CATEGORY_TYPE_EDUCATIONAL:   biz.CategoryTypeEducational,
		v1.CategoryType_CATEGORY_TYPE_NEWS:          biz.CategoryTypeNews,
		v1.CategoryType_CATEGORY_TYPE_ENTERTAINMENT: biz.CategoryTypeEntertainment,
	}

	BizToProtoCategoryType = map[biz.CategoryType]v1.CategoryType{
		biz.CategoryTypePodcast:       v1.CategoryType_CATEGORY_TYPE_PODCAST,
		biz.CategoryTypeDocumentary:   v1.CategoryType_CATEGORY_TYPE_DOCUMENTARY,
		biz.CategoryTypeSportsEvent:   v1.CategoryType_CATEGORY_TYPE_SPORTS_EVENT,
		biz.CategoryTypeEducational:   v1.CategoryType_CATEGORY_TYPE_EDUCATIONAL,
		biz.CategoryTypeNews:          v1.CategoryType_CATEGORY_TYPE_NEWS,
		biz.CategoryTypeEntertainment: v1.CategoryType_CATEGORY_TYPE_ENTERTAINMENT,
	}

	ProtoToBizProgramStatus = map[v1.ProgramStatus]biz.ProgramStatus{
		v1.ProgramStatus_PROGRAM_STATUS_DRAFT:     biz.ProgramStatusDraft,
		v1.ProgramStatus_PROGRAM_STATUS_PUBLISHED: biz.ProgramStatusPublished,
		v1.ProgramStatus_PROGRAM_STATUS_ARCHIVED:  biz.ProgramStatusArchived,
	}

	BizToProtoProgramStatus = map[biz.ProgramStatus]v1.ProgramStatus{
		biz.ProgramStatusDraft:     v1.ProgramStatus_PROGRAM_STATUS_DRAFT,
		biz.ProgramStatusPublished: v1.ProgramStatus_PROGRAM_STATUS_PUBLISHED,
		biz.ProgramStatusArchived:  v1.ProgramStatus_PROGRAM_STATUS_ARCHIVED,
	}

	ProtoToBizEpisodeStatus = map[v1.EpisodeStatus]biz.EpisodeStatus{
		v1.EpisodeStatus_EPISODE_STATUS_DRAFT:     biz.EpisodeStatusDraft,
		v1.EpisodeStatus_EPISODE_STATUS_PUBLISHED: biz.EpisodeStatusPublished,
		v1.EpisodeStatus_EPISODE_STATUS_SCHEDULED: biz.EpisodeStatusScheduled,
		v1.EpisodeStatus_EPISODE_STATUS_ARCHIVED:  biz.EpisodeStatusArchived,
	}

	BizToProtoEpisodeStatus = map[biz.EpisodeStatus]v1.EpisodeStatus{
		biz.EpisodeStatusDraft:     v1.EpisodeStatus_EPISODE_STATUS_DRAFT,
		biz.EpisodeStatusPublished: v1.EpisodeStatus_EPISODE_STATUS_PUBLISHED,
		biz.EpisodeStatusScheduled: v1.EpisodeStatus_EPISODE_STATUS_SCHEDULED,
		biz.EpisodeStatusArchived:  v1.EpisodeStatus_EPISODE_STATUS_ARCHIVED,
	}

	ProtoToBizImportStatus = map[v1.ImportStatus]biz.ImportStatus{
		v1.ImportStatus_IMPORT_STATUS_PENDING:    biz.ImportStatusPending,
		v1.ImportStatus_IMPORT_STATUS_PROCESSING: biz.ImportStatusProcessing,
		v1.ImportStatus_IMPORT_STATUS_COMPLETED:  biz.ImportStatusCompleted,
		v1.ImportStatus_IMPORT_STATUS_FAILED:     biz.ImportStatusFailed,
	}

	BizToProtoImportStatus = map[biz.ImportStatus]v1.ImportStatus{
		biz.ImportStatusPending:    v1.ImportStatus_IMPORT_STATUS_PENDING,
		biz.ImportStatusProcessing: v1.ImportStatus_IMPORT_STATUS_PROCESSING,
		biz.ImportStatusCompleted:  v1.ImportStatus_IMPORT_STATUS_COMPLETED,
		biz.ImportStatusFailed:     v1.ImportStatus_IMPORT_STATUS_FAILED,
	}
)

func ConvertFullUser(user *biz.User) *v1.User {
	return &v1.User{
		Id:        user.ID.String(),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Name:      user.Name,
		Email:     user.Email,
	}
}

func BizProgramStatusToProto(status biz.ProgramStatus) v1.ProgramStatus {
	if protoStatus, exists := BizToProtoProgramStatus[status]; exists {
		return protoStatus
	}
	return v1.ProgramStatus_PROGRAM_STATUS_DRAFT // Default fallback
}

func BizEpisodeStatusToProto(status biz.EpisodeStatus) v1.EpisodeStatus {
	if protoStatus, exists := BizToProtoEpisodeStatus[status]; exists {
		return protoStatus
	}
	return v1.EpisodeStatus_EPISODE_STATUS_DRAFT // Default fallback
}

func BizCategoryTypeToProto(catType biz.CategoryType) v1.CategoryType {
	if protoType, exists := BizToProtoCategoryType[catType]; exists {
		return protoType
	}
	return v1.CategoryType_CATEGORY_TYPE_PODCAST // Default fallback
}

func BizImportStatusToProto(status biz.ImportStatus) v1.ImportStatus {
	if protoStatus, exists := BizToProtoImportStatus[status]; exists {
		return protoStatus
	}
	return v1.ImportStatus_IMPORT_STATUS_PENDING // Default fallback
}

func ConvertMetadata(m map[string]string) biz.Metadata {
	if m == nil {
		return nil
	}
	metadata := biz.Metadata{}
	for k, v := range m {
		metadata[k] = v
	}
	return metadata
}

func ConvertProgram(p *biz.Program) *v1.Program {
	if p == nil {
		return nil
	}

	program := &v1.Program{
		Id:            p.ID.String(),
		Title:         p.Title,
		Description:   p.Description,
		CategoryId:    p.CategoryID.String(),
		Status:        BizProgramStatusToProto(p.Status),
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.UpdatedAt),
		CreatedBy:     p.CreatedBy.String(),
		UpdatedBy:     p.UpdatedBy.String(),
		ThumbnailUrl:  p.ThumbnailURL,
		Tags:          p.Tags,
		Metadata:      make(map[string]string),
		IsFeatured:    p.IsFeatured,
		EpisodesCount: p.EpisodesCount,
		ViewCount:     p.ViewCount,
		Rating:        p.Rating,
	}

	if p.PublishedAt != nil {
		program.PublishedAt = timestamppb.New(*p.PublishedAt)
	}

	if p.SourceURL != nil {
		program.SourceUrl = p.SourceURL
	}

	if p.Metadata != nil {
		for k, v := range p.Metadata {
			program.Metadata[k] = v
		}
	}

	return program
}

func ConvertPrograms(programs []*biz.Program) []*v1.Program {
	if programs == nil {
		return nil
	}

	result := make([]*v1.Program, 0, len(programs))
	for _, p := range programs {
		result = append(result, ConvertProgram(p))
	}
	return result
}

func ConvertCategory(c *biz.Category) *v1.Category {
	if c == nil {
		return nil
	}

	category := &v1.Category{
		Id:          c.ID.String(),
		Name:        c.Name,
		Description: c.Description,
		Type:        BizCategoryTypeToProto(c.Type),
		CreatedAt:   timestamppb.New(c.CreatedAt),
		UpdatedAt:   timestamppb.New(c.UpdatedAt),
		CreatedBy:   c.CreatedBy.String(),
		Metadata:    make(map[string]string),
	}

	if c.Metadata != nil {
		for k, v := range c.Metadata {
			category.Metadata[k] = v
		}
	}

	return category
}

func ConvertCategories(categories []*biz.Category) []*v1.Category {
	if categories == nil {
		return nil
	}

	result := make([]*v1.Category, 0, len(categories))
	for _, c := range categories {
		result = append(result, ConvertCategory(c))
	}
	return result
}

func ConvertEpisode(e *biz.Episode) *v1.Episode {
	if e == nil {
		return nil
	}

	episode := &v1.Episode{
		Id:              e.ID.String(),
		ProgramId:       e.ProgramID.String(),
		Title:           e.Title,
		Description:     e.Description,
		DurationSeconds: e.DurationSecs,
		EpisodeNumber:   e.EpisodeNumber,
		SeasonNumber:    e.SeasonNumber,
		Status:          BizEpisodeStatusToProto(e.Status),
		CreatedAt:       timestamppb.New(e.CreatedAt),
		UpdatedAt:       timestamppb.New(e.UpdatedAt),
		CreatedBy:       e.CreatedBy.String(),
		UpdatedBy:       e.UpdatedBy.String(),
		MediaUrl:        e.MediaURL,
		ThumbnailUrl:    e.ThumbnailURL,
		Tags:            e.Tags,
		Metadata:        make(map[string]string),
		ViewCount:       e.ViewCount,
		Rating:          e.Rating,
	}

	if e.PublishedAt != nil {
		episode.PublishedAt = timestamppb.New(*e.PublishedAt)
	}

	if e.ScheduledAt != nil {
		episode.ScheduledAt = timestamppb.New(*e.ScheduledAt)
	}

	if e.Metadata != nil {
		for k, v := range e.Metadata {
			episode.Metadata[k] = v
		}
	}

	return episode
}

func ConvertEpisodes(episodes []*biz.Episode) []*v1.Episode {
	if episodes == nil {
		return nil
	}

	result := make([]*v1.Episode, 0, len(episodes))
	for _, e := range episodes {
		result = append(result, ConvertEpisode(e))
	}
	return result
}
