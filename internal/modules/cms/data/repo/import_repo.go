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
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type importRepo struct {
	db *pgxpool.Pool
}

func NewImportRepository(db *pgxpool.Pool) biz.ImportRepository {
	return &importRepo{
		db: db,
	}
}

func (r *importRepo) Create(ctx context.Context, importData *biz.ImportData) error {
	if importData.ID == uuid.Nil {
		importData.ID = uuid.Must(uuid.NewV7())
	}
	importData.CreatedAt = time.Now()
	importData.UpdatedAt = time.Now()

	query, args, err := goqu.Insert("imports").Rows(goqu.Record{
		"id":              importData.ID,
		"source_type":     importData.SourceType,
		"source_url":      importData.SourceURL,
		"source_config":   importData.SourceConfig,
		"category_id":     importData.CategoryID,
		"status":          importData.Status,
		"total_items":     importData.TotalItems,
		"processed_items": importData.ProcessedItems,
		"success_count":   importData.SuccessCount,
		"error_count":     importData.ErrorCount,
		"errors":          pq.Array(importData.Errors),
		"warnings":        pq.Array(importData.Warnings),
		"created_at":      importData.CreatedAt,
		"updated_at":      importData.UpdatedAt,
		"created_by":      importData.CreatedBy,
		"updated_by":      importData.CreatedBy, // Use CreatedBy for UpdatedBy since struct doesn't have UpdatedBy
		"field_mapping":   importData.FieldMapping,
	}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to insert import data: %w", err)
	}

	return nil
}

func (r *importRepo) Update(ctx context.Context, userID, id uuid.UUID, updates *biz.UpdateImportRequest) (*biz.ImportData, error) {
	updateRecord := goqu.Record{
		"updated_at": time.Now(),
	}

	if updates.SourceType != nil {
		updateRecord["source_type"] = *updates.SourceType
	}
	if updates.SourceURL != nil {
		updateRecord["source_url"] = *updates.SourceURL
	}
	if updates.SourceConfig != nil {
		updateRecord["source_config"] = *updates.SourceConfig
	}
	if updates.CategoryID != nil {
		updateRecord["category_id"] = *updates.CategoryID
	}
	if updates.Status != nil {
		updateRecord["status"] = *updates.Status
	}
	if updates.TotalItems != nil {
		updateRecord["total_items"] = *updates.TotalItems
	}
	if updates.ProcessedItems != nil {
		updateRecord["processed_items"] = *updates.ProcessedItems
	}
	if updates.SuccessCount != nil {
		updateRecord["success_count"] = *updates.SuccessCount
	}
	if updates.ErrorCount != nil {
		updateRecord["error_count"] = *updates.ErrorCount
	}
	if updates.Errors != nil {
		updateRecord["errors"] = pq.Array(*updates.Errors)
	}
	if updates.Warnings != nil {
		updateRecord["warnings"] = pq.Array(*updates.Warnings)
	}
	if updates.FieldMapping != nil {
		updateRecord["field_mapping"] = *updates.FieldMapping
	}

	query, args, err := goqu.Update("imports").
		Set(updateRecord).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to update import data: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *importRepo) GetByID(ctx context.Context, id uuid.UUID) (*biz.ImportData, error) {
	query, args, err := goqu.Select(
		"id",
		"source_type",
		"source_url",
		"source_config",
		"category_id",
		"status",
		"total_items",
		"processed_items",
		"success_count",
		"error_count",
		"errors",
		"warnings",
		"created_at",
		"updated_at",
		"created_by",
		"field_mapping",
	).From("imports").
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var importData biz.ImportData
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&importData.ID,
		&importData.SourceType,
		&importData.SourceURL,
		&importData.SourceConfig,
		&importData.CategoryID,
		&importData.Status,
		&importData.TotalItems,
		&importData.ProcessedItems,
		&importData.SuccessCount,
		&importData.ErrorCount,
		&importData.Errors,
		&importData.Warnings,
		&importData.CreatedAt,
		&importData.UpdatedAt,
		&importData.CreatedBy,
		&importData.FieldMapping,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("import data not found")
		}
		return nil, fmt.Errorf("failed to scan import data: %w", err)
	}

	return &importData, nil
}

func (r *importRepo) List(ctx context.Context, filter biz.ImportFilter, pagination biz.PaginationRequest, sort biz.SortRequest) ([]*biz.ImportData, *biz.PaginationResponse, error) {
	pagination.SetDefaults()

	var conditions []exp.Expression

	if filter.Status != nil {
		conditions = append(conditions, goqu.C("status").Eq(*filter.Status))
	}

	if filter.SearchQuery != nil && *filter.SearchQuery != "" {
		searchPattern := "%" + *filter.SearchQuery + "%"
		conditions = append(conditions, goqu.Or(
			goqu.C("source_type").ILike(searchPattern),
			goqu.C("source_url").ILike(searchPattern),
		))
	}

	if filter.CreatedBy != nil {
		conditions = append(conditions, goqu.C("created_by").Eq(*filter.CreatedBy))
	}

	// Count total records
	countQuery, countArgs, err := goqu.Select(goqu.COUNT("*")).
		From("imports").
		Where(conditions...).
		ToSQL()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	var totalCount int32
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count imports: %w", err)
	}

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

	offset := (pagination.Page - 1) * pagination.PageSize
	limit := pagination.PageSize

	selectQuery, selectArgs, err := goqu.Select(
		"id",
		"source_type",
		"source_url",
		"source_config",
		"category_id",
		"status",
		"total_items",
		"processed_items",
		"success_count",
		"error_count",
		"errors",
		"warnings",
		"created_at",
		"updated_at",
		"created_by",
		"field_mapping",
	).From("imports").
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
		return nil, nil, fmt.Errorf("failed to query imports: %w", err)
	}
	defer rows.Close()

	var imports []*biz.ImportData
	for rows.Next() {
		var importData biz.ImportData
		err := rows.Scan(
			&importData.ID,
			&importData.SourceType,
			&importData.SourceURL,
			&importData.SourceConfig,
			&importData.CategoryID,
			&importData.Status,
			&importData.TotalItems,
			&importData.ProcessedItems,
			&importData.SuccessCount,
			&importData.ErrorCount,
			&importData.Errors,
			&importData.Warnings,
			&importData.CreatedAt,
			&importData.UpdatedAt,
			&importData.CreatedBy,
			&importData.FieldMapping,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to scan import data: %w", err)
		}
		imports = append(imports, &importData)
	}

	// Calculate pagination response
	totalPages := (totalCount + pagination.PageSize - 1) / pagination.PageSize
	paginationResponse := &biz.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return imports, paginationResponse, nil
}

func (r *importRepo) UpdateProgress(ctx context.Context, userID, id uuid.UUID, updates *biz.UpdateImportProgressRequest) error {
	updateRecord := goqu.Record{
		"updated_at": time.Now(),
	}

	if updates.ProcessedItems != nil {
		updateRecord["processed_items"] = *updates.ProcessedItems
	}
	if updates.SuccessCount != nil {
		updateRecord["success_count"] = *updates.SuccessCount
	}
	if updates.ErrorCount != nil {
		updateRecord["error_count"] = *updates.ErrorCount
	}

	query, args, err := goqu.Update("imports").
		Set(updateRecord).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build update progress query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update import progress: %w", err)
	}

	return nil
}

func (r *importRepo) AddError(ctx context.Context, userID, id uuid.UUID, errorMsg string) error {
	// First get the current errors
	importData, err := r.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get import data: %w", err)
	}

	// Check if user owns this import
	if importData.CreatedBy != userID {
		return fmt.Errorf("import not found")
	}

	// Add the new error
	errors := append(importData.Errors, errorMsg)
	errorCount := int32(len(errors))

	query, args, err := goqu.Update("imports").
		Set(goqu.Record{
			"errors":      pq.Array(errors),
			"error_count": errorCount,
			"updated_at":  time.Now(),
		}).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build add error query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to add error to import: %w", err)
	}

	return nil
}

func (r *importRepo) AddWarning(ctx context.Context, userID, id uuid.UUID, warningMsg string) error {
	// First get the current warnings
	importData, err := r.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get import data: %w", err)
	}

	// Check if user owns this import
	if importData.CreatedBy != userID {
		return fmt.Errorf("import not found")
	}

	// Add the new warning
	warnings := append(importData.Warnings, warningMsg)

	query, args, err := goqu.Update("imports").
		Set(goqu.Record{
			"warnings":   pq.Array(warnings),
			"updated_at": time.Now(),
		}).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build add warning query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to add warning to import: %w", err)
	}

	return nil
}
