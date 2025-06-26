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
)

type categoryRepo struct {
	db    *pgxpool.Pool
	table string
}

func NewCategoryRepository(db *pgxpool.Pool) biz.CategoryRepository {
	return &categoryRepo{
		db:    db,
		table: "categories",
	}
}

func (r *categoryRepo) Create(ctx context.Context, category *biz.Category) error {
	now := time.Now()
	if category.ID == uuid.Nil {
		category.ID = uuid.Must(uuid.NewV7())
	}
	category.CreatedAt = now
	category.UpdatedAt = now

	query, args, err := goqu.Insert("categories").Rows(goqu.Record{
		"id":          category.ID,
		"name":        category.Name,
		"description": category.Description,
		"type":        category.Type,
		"created_at":  category.CreatedAt,
		"updated_at":  category.UpdatedAt,
		"created_by":  category.CreatedBy,
		"metadata":    category.Metadata,
	}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "categories_name_key" {
				return biz.ErrCategoryAlreadyExists
			}
		}
		return fmt.Errorf("failed to insert category: %w", err)
	}

	return nil
}

func (r *categoryRepo) Update(ctx context.Context, userID, id uuid.UUID, updates *biz.UpdateCategoryRequest) (*biz.Category, error) {
	updateRecord := goqu.Record{
		"updated_at": time.Now(),
	}

	if updates.Name != nil {
		updateRecord["name"] = *updates.Name
	}
	if updates.Description != nil {
		updateRecord["description"] = *updates.Description
	}
	if updates.Type != nil {
		updateRecord["type"] = *updates.Type
	}
	if updates.Metadata != nil {
		updateRecord["metadata"] = *updates.Metadata
	}

	query, args, err := goqu.Update("categories").
		Set(updateRecord).
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build update query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "categories_name_key" {
				return nil, biz.ErrCategoryAlreadyExists
			}
		}
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return r.GetByID(ctx, id)
}

func (r *categoryRepo) Delete(ctx context.Context, userID, id uuid.UUID) error {
	query, args, err := goqu.Delete("categories").
		Where(goqu.C("id").Eq(id)).
		Where(goqu.C("created_by").Eq(userID)).
		ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build delete query: %w", err)
	}

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("category not found")
	}

	return nil
}

func (r *categoryRepo) GetByID(ctx context.Context, id uuid.UUID) (*biz.Category, error) {
	query, args, err := goqu.Select(
		"id",
		"name",
		"description",
		"type",
		"created_at",
		"updated_at",
		"created_by",
		"metadata",
	).From("categories").
		Where(goqu.C("id").Eq(id)).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}

	var category biz.Category
	err = r.db.QueryRow(ctx, query, args...).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, biz.ErrCategoryNotFound
		}
		return nil, fmt.Errorf("failed to scan category: %w", err)
	}

	return &category, nil
}

func (r *categoryRepo) List(ctx context.Context, filter biz.CategoryFilter, pagination biz.PaginationRequest, sort biz.SortRequest) ([]*biz.Category, *biz.PaginationResponse, error) {
	pagination.SetDefaults()

	var conditions []exp.Expression

	if filter.Type != nil {
		conditions = append(conditions, goqu.C("type").Eq(*filter.Type))
	}

	if filter.SearchQuery != nil && *filter.SearchQuery != "" {
		searchPattern := "%" + *filter.SearchQuery + "%"
		conditions = append(conditions, goqu.Or(
			goqu.C("name").ILike(searchPattern),
			goqu.C("description").ILike(searchPattern),
		))
	}

	if filter.CreatedBy != nil {
		conditions = append(conditions, goqu.C("created_by").Eq(*filter.CreatedBy))
	}

	// Count total records
	countQuery, countArgs, err := goqu.Select(goqu.COUNT("*")).
		From("categories").
		Where(conditions...).
		ToSQL()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build count query: %w", err)
	}

	var totalCount int32
	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&totalCount)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to count categories: %w", err)
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
		"name",
		"description",
		"type",
		"created_at",
		"updated_at",
		"created_by",
		"metadata",
	).From("categories").
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
		return nil, nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()

	var categories []*biz.Category
	for rows.Next() {
		var category biz.Category
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
			return nil, nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, &category)
	}

	// Calculate pagination response
	totalPages := (totalCount + pagination.PageSize - 1) / pagination.PageSize
	paginationResponse := &biz.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalCount: totalCount,
		TotalPages: totalPages,
	}

	return categories, paginationResponse, nil
}
