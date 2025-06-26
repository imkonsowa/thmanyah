package repo

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
	cms "thmanyah/internal/modules/cms/biz"
	"thmanyah/internal/modules/discover/biz"
)

type discoverRepo struct {
	logger *log.Helper
	db     *pgxpool.Pool
}

func NewDiscoverRepo(db *pgxpool.Pool, logger log.Logger) (biz.DiscoverRepository, error) {
	return &discoverRepo{
		logger: log.NewHelper(logger),
		db:     db,
	}, nil
}

func (s *discoverRepo) Search(ctx context.Context, query string, page, pageSize int32) (*biz.SearchResults, int32, error) {
	if query == "" {
		return &biz.SearchResults{
			Categories: []*cms.Category{},
			Programs:   []*cms.Program{},
			Episodes:   []*cms.Episode{},
		}, 0, nil
	}

	searchQuery := strings.TrimSpace(query)
	offset := (page - 1) * pageSize

	results := &biz.SearchResults{
		Categories: []*cms.Category{},
		Programs:   []*cms.Program{},
		Episodes:   []*cms.Episode{},
	}

	// Search categories
	categories, err := s.searchCategories(ctx, searchQuery)
	if err != nil {
		s.logger.Errorf("Failed to search categories: %v", err)
		return nil, 0, err
	}
	results.Categories = categories

	// Search programs
	programs, err := s.searchPrograms(ctx, searchQuery, pageSize, offset)
	if err != nil {
		s.logger.Errorf("Failed to search programs: %v", err)
		return nil, 0, err
	}
	results.Programs = programs

	// Search episodes
	episodes, err := s.searchEpisodes(ctx, searchQuery, pageSize, offset)
	if err != nil {
		s.logger.Errorf("Failed to search episodes: %v", err)
		return nil, 0, err
	}
	results.Episodes = episodes

	// Count total results
	totalCount := int32(len(categories) + len(programs) + len(episodes))

	return results, totalCount, nil
}

func (s *discoverRepo) searchCategories(ctx context.Context, query string) ([]*cms.Category, error) {
	sql := `
		SELECT id, name, description, type, created_at, updated_at, created_by, metadata
		FROM categories 
		WHERE to_tsvector('simple', unaccent(name || ' ' || coalesce(description, ''))) @@ plainto_tsquery('simple', unaccent($1))
		ORDER BY name
		LIMIT 10
	`

	rows, err := s.db.Query(ctx, sql, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*cms.Category
	for rows.Next() {
		category := &cms.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.Type,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.CreatedBy,
			&category.Metadata,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, rows.Err()
}

func (s *discoverRepo) searchPrograms(ctx context.Context, query string, limit, offset int32) ([]*cms.Program, error) {
	sql := `
		SELECT id, title, description, category_id, status, created_at, updated_at, 
		       published_at, created_by, updated_by, thumbnail_url, tags, metadata, 
		       source_url, episodes_count, is_featured, view_count, rating
		FROM programs 
		WHERE search_vector @@ plainto_tsquery('simple', unaccent($1))
		ORDER BY ts_rank(search_vector, plainto_tsquery('simple', unaccent($1))) DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := s.db.Query(ctx, sql, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []*cms.Program
	for rows.Next() {
		program := &cms.Program{}
		err := rows.Scan(
			&program.ID,
			&program.Title,
			&program.Description,
			&program.CategoryID,
			&program.Status,
			&program.CreatedAt,
			&program.UpdatedAt,
			&program.PublishedAt,
			&program.CreatedBy,
			&program.UpdatedBy,
			&program.ThumbnailURL,
			&program.Tags,
			&program.Metadata,
			&program.SourceURL,
			&program.EpisodesCount,
			&program.IsFeatured,
			&program.ViewCount,
			&program.Rating,
		)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	return programs, rows.Err()
}

func (s *discoverRepo) searchEpisodes(ctx context.Context, query string, limit, offset int32) ([]*cms.Episode, error) {
	sql := `
		SELECT id, program_id, title, description, duration_seconds, episode_number, season_number, 
		       status, created_at, updated_at, published_at, scheduled_at, created_by, updated_by, 
		       media_url, thumbnail_url, tags, metadata, view_count, rating
		FROM episodes 
		WHERE search_vector @@ plainto_tsquery('simple', unaccent($1))
		ORDER BY ts_rank(search_vector, plainto_tsquery('simple', unaccent($1))) DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := s.db.Query(ctx, sql, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var episodes []*cms.Episode
	for rows.Next() {
		episode := &cms.Episode{}
		err := rows.Scan(
			&episode.ID,
			&episode.ProgramID,
			&episode.Title,
			&episode.Description,
			&episode.DurationSecs,
			&episode.EpisodeNumber,
			&episode.SeasonNumber,
			&episode.Status,
			&episode.CreatedAt,
			&episode.UpdatedAt,
			&episode.PublishedAt,
			&episode.ScheduledAt,
			&episode.CreatedBy,
			&episode.UpdatedBy,
			&episode.MediaURL,
			&episode.ThumbnailURL,
			&episode.Tags,
			&episode.Metadata,
			&episode.ViewCount,
			&episode.Rating,
		)
		if err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}

	return episodes, rows.Err()
}

func (s *discoverRepo) Featured(ctx context.Context) ([]*cms.Program, error) {
	sql := `
		SELECT id, title, description, category_id, status, created_at, updated_at, 
		       published_at, created_by, updated_by, thumbnail_url, tags, metadata, 
		       source_url, episodes_count, is_featured, view_count, rating
		FROM programs 
		WHERE is_featured = true
		ORDER BY created_at DESC
	`

	rows, err := s.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []*cms.Program
	for rows.Next() {
		program := &cms.Program{}
		err := rows.Scan(
			&program.ID,
			&program.Title,
			&program.Description,
			&program.CategoryID,
			&program.Status,
			&program.CreatedAt,
			&program.UpdatedAt,
			&program.PublishedAt,
			&program.CreatedBy,
			&program.UpdatedBy,
			&program.ThumbnailURL,
			&program.Tags,
			&program.Metadata,
			&program.SourceURL,
			&program.EpisodesCount,
			&program.IsFeatured,
			&program.ViewCount,
			&program.Rating,
		)
		if err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	return programs, rows.Err()
}
