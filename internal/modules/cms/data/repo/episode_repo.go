package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"thmanyah/internal/modules/cms/biz"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type episodeRepo struct {
	db    *pgxpool.Pool
	table string
}

func NewEpisodeRepository(db *pgxpool.Pool) biz.EpisodeRepository {
	return &episodeRepo{
		db:    db,
		table: "episodes",
	}
}

func (r *episodeRepo) Create(ctx context.Context, episode *biz.Episode) error {
	if episode.ID == uuid.Nil {
		episode.ID = uuid.Must(uuid.NewV7())
	}
	episode.CreatedAt = time.Now()
	episode.UpdatedAt = time.Now()

	record := goqu.Record{
		"id":               episode.ID,
		"program_id":       episode.ProgramID,
		"title":            episode.Title,
		"description":      episode.Description,
		"duration_seconds": episode.DurationSecs,
		"episode_number":   episode.EpisodeNumber,
		"season_number":    episode.SeasonNumber,
		"status":           episode.Status,
		"created_at":       episode.CreatedAt,
		"updated_at":       episode.UpdatedAt,
		"published_at":     episode.PublishedAt,
		"scheduled_at":     episode.ScheduledAt,
		"created_by":       episode.CreatedBy,
		"updated_by":       episode.UpdatedBy,
		"media_url":        episode.MediaURL,
		"thumbnail_url":    episode.ThumbnailURL,
		"metadata":         episode.Metadata,
		"view_count":       episode.ViewCount,
		"rating":           episode.Rating,
	}

	if len(episode.Tags) > 0 {
		record["tags"] = pq.Array(episode.Tags)
	} else {
		// Use empty array literal
		record["tags"] = pq.Array([]string{})
	}

	query, args, err := goqu.Insert("episodes").Rows(record).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// Check for unique constraint violation (episode number/season/program) (SQLSTATE 23505)
			if pgErr.Code == "23505" && pgErr.ConstraintName == "episodes_program_id_season_number_episode_number_key" {
				return biz.ErrEpisodeAlreadyExists
			}
			// Check for foreign key constraint violation on program_id (SQLSTATE 23503)
			if pgErr.Code == "23503" && pgErr.ConstraintName != "" && pgErr.ConstraintName == "episodes_program_id_fkey" {
				return biz.ErrProgramNotFound
			}
		}
		return fmt.Errorf("failed to insert episode: %w", err)
	}

	return nil
}

func (r *episodeRepo) Update(ctx context.Context, userID, id uuid.UUID, updates *biz.UpdateEpisodeRequest) (*biz.Episode, error) {
	// Build update record with only safe fields
	updateRecord := goqu.Record{
		"updated_at": time.Now(),
	}

	// Only update fields that are provided and safe to update
	if updates.Title != nil {
		updateRecord["title"] = *updates.Title
	}
	if updates.Description != nil {
		updateRecord["description"] = *updates.Description
	}
	if updates.DurationSecs != nil {
		updateRecord["duration_seconds"] = *updates.DurationSecs
	}
	if updates.EpisodeNumber != nil {
		updateRecord["episode_number"] = *updates.EpisodeNumber
	}
	if updates.SeasonNumber != nil {
		updateRecord["season_number"] = *updates.SeasonNumber
	}
	if updates.Status != nil {
		updateRecord["status"] = *updates.Status
	}
	if updates.PublishedAt != nil {
		updateRecord["published_at"] = *updates.PublishedAt
	}
	if updates.ScheduledAt != nil {
		updateRecord["scheduled_at"] = *updates.ScheduledAt
	}
	if updates.MediaURL != nil {
		updateRecord["media_url"] = *updates.MediaURL
	}
	if updates.ThumbnailURL != nil {
		updateRecord["thumbnail_url"] = *updates.ThumbnailURL
	}
	if updates.Tags != nil {
		if len(*updates.Tags) > 0 {
			updateRecord["tags"] = pq.Array(*updates.Tags)
		} else {
			updateRecord["tags"] = pq.Array([]string{})
		}
	}
	if updates.Metadata != nil {
		updateRecord["metadata"] = *updates.Metadata
	}
	if updates.ViewCount != nil {
		updateRecord["view_count"] = *updates.ViewCount
	}
	if updates.Rating != nil {
		updateRecord["rating"] = *updates.Rating
	}

	query, args, err := goqu.Update("episodes").
		Set(updateRecord).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update episode: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *episodeRepo) Delete(ctx context.Context, userID, id uuid.UUID) error {
	status := biz.EpisodeStatusArchived
	_, err := r.Update(ctx, userID, id, &biz.UpdateEpisodeRequest{
		Status: &status,
	})
	if err != nil {
		return fmt.Errorf("failed to soft delete episode: %w", err)
	}

	return nil
}

func (r *episodeRepo) GetByID(ctx context.Context, id uuid.UUID) (*biz.Episode, error) {
	var query string
	var args []interface{}
	var err error

	query, args, err = goqu.Select(
		"id",
		"program_id",
		"title",
		"description",
		"duration_seconds",
		"episode_number",
		"season_number",
		"status",
		"created_at",
		"updated_at",
		"published_at",
		"scheduled_at",
		"created_by",
		"updated_by",
		"media_url",
		"thumbnail_url",
		"tags",
		"metadata",
		"view_count",
		"rating",
	).From("episodes").
		Where(goqu.C("id").Eq(id)).
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var episode biz.Episode

	err = r.db.QueryRow(ctx, query, args...).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("episode not found")
		}
		return nil, fmt.Errorf("failed to scan episode: %w", err)
	}

	return &episode, nil
}

func (r *episodeRepo) List(ctx context.Context, filter biz.EpisodeFilter, pagination biz.PaginationRequest, sort biz.SortRequest) ([]*biz.Episode, *biz.PaginationResponse, error) {
	pagination.SetDefaults()

	// Build WHERE conditions
	var conditions []exp.Expression

	// Filter by program_id if provided
	if filter.ProgramID != nil {
		conditions = append(conditions, goqu.C("program_id").Eq(*filter.ProgramID))
	}

	if filter.Status != nil {
		conditions = append(conditions, goqu.C("status").Eq(*filter.Status))
	}

	if filter.SearchQuery != nil && *filter.SearchQuery != "" {
		searchPattern := "%" + *filter.SearchQuery + "%"
		conditions = append(conditions, goqu.Or(
			goqu.C("title").ILike(searchPattern),
			goqu.C("description").ILike(searchPattern),
		))
	}

	if filter.CreatedBy != nil {
		conditions = append(conditions, goqu.C("created_by").Eq(*filter.CreatedBy))
	}

	// Count total records
	countQuery, countArgs, err := goqu.Select(goqu.COUNT("*")).
		From("episodes").
		Where(conditions...).
		ToSQL()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	var totalCount int32
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count episodes: %w", err)
	}

	// Build ORDER BY clause
	var orderBy exp.OrderedExpression
	if sort.SortBy != "" {
		if sort.SortOrder == "desc" {
			orderBy = goqu.C(sort.SortBy).Desc()
		} else {
			orderBy = goqu.C(sort.SortBy).Asc()
		}
	} else {
		// Default ordering: season_number ASC, episode_number ASC
		orderBy = goqu.C("created_at").Asc()
	}

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	limit := pagination.PageSize

	// Build final query
	selectQuery, selectArgs, err := goqu.Select(
		"id",
		"program_id",
		"title",
		"description",
		"duration_seconds",
		"episode_number",
		"season_number",
		"status",
		"created_at",
		"updated_at",
		"published_at",
		"scheduled_at",
		"created_by",
		"updated_by",
		"media_url",
		"thumbnail_url",
		"tags",
		"metadata",
		"view_count",
		"rating",
	).From("episodes").
		Where(conditions...).
		Order(orderBy).
		Limit(uint(limit)).
		Offset(uint(offset)).
		ToSQL()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build select query: %w", err)
	}

	rows, err := r.db.Query(ctx, selectQuery, selectArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query episodes: %w", err)
	}
	defer rows.Close()

	var episodes []*biz.Episode
	for rows.Next() {
		var episode biz.Episode
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
			return nil, nil, fmt.Errorf("failed to scan episode: %w", err)
		}
		episodes = append(episodes, &episode)
	}

	// Calculate pagination response
	totalPages := (totalCount + pagination.PageSize - 1) / pagination.PageSize
	paginationResponse := &biz.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return episodes, paginationResponse, nil
}

func (r *episodeRepo) ListByProgram(ctx context.Context, programID uuid.UUID, pagination biz.PaginationRequest, sort biz.SortRequest) ([]*biz.Episode, *biz.PaginationResponse, error) {
	pagination.SetDefaults()

	filter := biz.EpisodeFilter{
		ProgramID: &programID,
	}
	return r.List(ctx, filter, pagination, sort)
}

func (r *episodeRepo) IncrementViewCount(ctx context.Context, id uuid.UUID) error {
	query, args, err := goqu.Update("episodes").
		Set(goqu.Record{
			"view_count": goqu.L("view_count + 1"),
		}).
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build increment view count query: %w", err)
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to increment view count: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("episode not found")
	}

	return nil
}
