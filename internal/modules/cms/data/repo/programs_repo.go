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

type programRepo struct {
	db    *pgxpool.Pool
	table string
}

func NewProgramRepository(db *pgxpool.Pool) biz.ProgramRepository {
	return &programRepo{
		db:    db,
		table: "programs",
	}
}

func (r *programRepo) Create(ctx context.Context, program *biz.Program) error {
	if program.ID == uuid.Nil {
		program.ID = uuid.Must(uuid.NewV7())
	}
	program.CreatedAt = time.Now()
	program.UpdatedAt = time.Now()

	record := goqu.Record{
		"id":             program.ID,
		"title":          program.Title,
		"description":    program.Description,
		"category_id":    program.CategoryID,
		"status":         program.Status,
		"created_at":     program.CreatedAt,
		"updated_at":     program.UpdatedAt,
		"published_at":   program.PublishedAt,
		"created_by":     program.CreatedBy,
		"updated_by":     program.UpdatedBy,
		"thumbnail_url":  program.ThumbnailURL,
		"metadata":       program.Metadata,
		"source_url":     program.SourceURL,
		"episodes_count": program.EpisodesCount,
		"is_featured":    program.IsFeatured,
		"view_count":     program.ViewCount,
		"rating":         program.Rating,
	}

	if len(program.Tags) > 0 {
		record["tags"] = pq.Array(program.Tags)
	} else {
		record["tags"] = pq.Array([]string{})
	}

	query, args, err := goqu.Insert("programs").Rows(record).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			// Check for foreign key constraint violation on category_id (SQLSTATE 23503)
			if pgErr.Code == "23503" && pgErr.ConstraintName == "programs_category_id_fkey" {
				return biz.ErrCategoryNotFound
			}
		}
		return fmt.Errorf("failed to insert program: %w", err)
	}

	return nil
}

func (r *programRepo) Update(ctx context.Context, userId, id uuid.UUID, updates *biz.UpdateProgramRequest) (*biz.Program, error) {
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
	if updates.CategoryID != nil {
		updateRecord["category_id"] = *updates.CategoryID
	}
	if updates.Status != nil {
		updateRecord["status"] = *updates.Status
	}
	if updates.PublishedAt != nil {
		updateRecord["published_at"] = *updates.PublishedAt
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
	if updates.SourceURL != nil {
		updateRecord["source_url"] = *updates.SourceURL
	}
	if updates.EpisodesCount != nil {
		updateRecord["episodes_count"] = *updates.EpisodesCount
	}
	if updates.IsFeatured != nil {
		updateRecord["is_featured"] = *updates.IsFeatured
	}
	if updates.ViewCount != nil {
		updateRecord["view_count"] = *updates.ViewCount
	}
	if updates.Rating != nil {
		updateRecord["rating"] = *updates.Rating
	}

	query, args, err := goqu.Update("programs").
		Set(updateRecord).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userId)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update program: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *programRepo) Delete(ctx context.Context, userId, id uuid.UUID) error {
	query, args, err := goqu.Delete("programs").
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userId)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("program not found")
	}

	return nil
}

func (r *programRepo) GetByID(ctx context.Context, id uuid.UUID) (*biz.Program, error) {
	var query string
	var args []interface{}
	var err error

	query, args, err = goqu.Select(
		"id",
		"title",
		"description",
		"category_id",
		"status",
		"created_at",
		"updated_at",
		"published_at",
		"created_by",
		"updated_by",
		"thumbnail_url",
		"tags",
		"metadata",
		"source_url",
		"episodes_count",
		"is_featured",
		"view_count",
		"rating",
	).From("programs").
		Where(goqu.C("id").Eq(id)).
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var program biz.Program

	err = r.db.QueryRow(ctx, query, args...).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("program not found")
		}
		return nil, fmt.Errorf("failed to scan program: %w", err)
	}

	return &program, nil
}

func (r *programRepo) List(ctx context.Context, filter biz.ProgramFilter, pagination biz.PaginationRequest, sort biz.SortRequest) ([]*biz.Program, *biz.PaginationResponse, error) {
	pagination.SetDefaults()

	// Build WHERE conditions
	var conditions []exp.Expression

	if filter.CategoryID != nil {
		conditions = append(conditions, goqu.C("category_id").Eq(*filter.CategoryID))
	}

	if filter.Status != nil {
		conditions = append(conditions, goqu.C("status").Eq(*filter.Status))
	}

	if filter.FeaturedOnly != nil && *filter.FeaturedOnly {
		conditions = append(conditions, goqu.C("is_featured").Eq(true))
	}

	if filter.SearchQuery != nil && *filter.SearchQuery != "" {
		searchPattern := "%" + *filter.SearchQuery + "%"
		conditions = append(conditions, goqu.Or(
			goqu.C("title").ILike(searchPattern),
			goqu.C("description").ILike(searchPattern),
		))
	}

	if len(filter.Tags) > 0 {
		tagConditions := make([]exp.Expression, len(filter.Tags))
		for i, tag := range filter.Tags {
			tagConditions[i] = goqu.L("? = ANY(tags)", tag)
		}
		conditions = append(conditions, goqu.Or(tagConditions...))
	}

	if filter.CreatedBy != nil {
		conditions = append(conditions, goqu.C("created_by").Eq(*filter.CreatedBy))
	}

	// Count total records
	countQuery, countArgs, err := goqu.Select(goqu.COUNT("*")).
		From("programs").
		Where(conditions...).
		ToSQL()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	var totalCount int32
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count programs: %w", err)
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
		orderBy = goqu.C("created_at").Desc()
	}

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	limit := pagination.PageSize

	// Build final query
	selectQuery, selectArgs, err := goqu.Select(
		"id",
		"title",
		"description",
		"category_id",
		"status",
		"created_at",
		"updated_at",
		"published_at",
		"created_by",
		"updated_by",
		"thumbnail_url",
		"tags",
		"metadata",
		"source_url",
		"episodes_count",
		"is_featured",
		"view_count",
		"rating",
	).From("programs").
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
		return nil, nil, fmt.Errorf("failed to query programs: %w", err)
	}
	defer rows.Close()

	var programs []*biz.Program
	for rows.Next() {
		var program biz.Program
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
			return nil, nil, fmt.Errorf("failed to scan program: %w", err)
		}
		programs = append(programs, &program)
	}

	// Calculate pagination response
	totalPages := (totalCount + pagination.PageSize - 1) / pagination.PageSize
	paginationResponse := &biz.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return programs, paginationResponse, nil
}

func (r *programRepo) BulkUpdate(ctx context.Context, userId uuid.UUID, ids []uuid.UUID, updates *biz.BulkUpdateProgramsRequest) (int32, error) {
	if len(ids) == 0 {
		return 0, nil
	}

	// Build update record with only safe fields
	updateRecord := goqu.Record{
		"updated_at": time.Now(),
	}

	// Only update fields that are provided and safe to update
	if updates.Status != nil {
		updateRecord["status"] = *updates.Status
	}
	if updates.CategoryID != nil {
		updateRecord["category_id"] = *updates.CategoryID
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
	if updates.IsFeatured != nil {
		updateRecord["is_featured"] = *updates.IsFeatured
	}

	query, args, err := goqu.Update("programs").
		Set(updateRecord).
		Where(goqu.C("id").In(ids)).
		Where(goqu.C("created_by").Eq(userId)).
		ToSQL()
	if err != nil {
		return 0, fmt.Errorf("failed to build bulk update query: %w", err)
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, fmt.Errorf("failed to bulk update programs: %w", err)
	}

	return int32(result.RowsAffected()), nil
}

func (r *programRepo) BulkDelete(ctx context.Context, userId uuid.UUID, ids []uuid.UUID) error {
	if len(ids) == 0 {
		return nil
	}

	query, args, err := goqu.Delete("programs").
		Where(goqu.C("id").In(ids)).
		Where(goqu.C("created_by").Eq(userId)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build bulk delete query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to bulk delete programs: %w", err)
	}

	return nil
}

func (r *programRepo) IncrementViewCount(ctx context.Context, id uuid.UUID) error {
	query, args, err := goqu.Update("programs").
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
		return fmt.Errorf("program not found")
	}

	return nil
}

func (r *programRepo) UpdateEpisodesCount(ctx context.Context, programID uuid.UUID) error {
	query, args, err := goqu.Update("programs").
		Set(goqu.Record{
			"episodes_count": goqu.Select(goqu.COUNT("*")).
				From("episodes").
				Where(goqu.And(
					goqu.C("program_id").Eq(programID),
					goqu.C("status").Neq(biz.EpisodeStatusArchived),
				)),
		}).
		Where(goqu.C("id").Eq(programID)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build update episodes count query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update episodes count: %w", err)
	}

	return nil
}
