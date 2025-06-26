package repo

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"thmanyah/internal/modules/cms/biz"
)

type userRepo struct {
	logger *log.Helper
	db     *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool, logger log.Logger) (biz.UsersRepository, error) {
	return &userRepo{
		logger: log.NewHelper(logger),
		db:     db,
	}, nil
}

func (u *userRepo) GetUserWithPassword(ctx context.Context, email string) (*biz.User, error) {
	query := goqu.From("users").
		Select(
			"id",
			"created_at",
			"updated_at",
			"email",
			"name",
			"password",
		).
		Where(goqu.Ex{"email": email})

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	user := &biz.User{}
	err = u.db.QueryRow(ctx, sql, params...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Name,
		&user.Password,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, biz.ErrInvalidCredentials
		}

		return nil, err
	}

	return user, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	var existingUserID string
	err := u.db.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUserID)
	if err == nil {
		return nil, biz.ErrUserAlreadyExists
	}
	if err.Error() != "no rows in result set" {
		return nil, err
	}

	now := time.Now().UTC()
	user.ID = uuid.Must(uuid.NewV7())
	user.CreatedAt = now
	user.UpdatedAt = now

	err = u.db.QueryRow(
		ctx,
		`INSERT INTO users (id, email, password, name, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id, created_at, updated_at`,
		user.ID.String(),
		user.Email,
		user.Password,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	).
		Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) GetUserByIdentifier(ctx context.Context, id string) (*biz.User, error) {
	whereClause := goqu.ExOr{
		"email": id,
	}

	_, err := uuid.Parse(id)
	if err == nil {
		whereClause["id"] = id
	}

	query := goqu.From("users").
		Select(
			"id",
			"created_at",
			"updated_at",
			"email",
			"name",
		).
		Where(whereClause)

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	user := &biz.User{}
	err = u.db.QueryRow(ctx, sql, params...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Name,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, biz.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, userId uuid.UUID, user *biz.UpdateUserRequest) (*biz.User, error) {
	rec := goqu.Record{}
	if user.Name != "" {
		rec["name"] = user.Name
	}
	if user.Email != "" {
		rec["email"] = user.Email
	}

	query := goqu.From("users").
		Where(goqu.Ex{"id": userId}).
		Update().
		Set(rec).
		Returning(
			"id",
			"created_at",
			"updated_at",
			"email",
			"name",
		)

	sql, params, err := query.ToSQL()
	if err != nil {
		return nil, err
	}

	userData := &biz.User{}
	err = u.db.QueryRow(ctx, sql, params...).Scan(
		&userData.ID,
		&userData.CreatedAt,
		&userData.UpdatedAt,
		&userData.Email,
		&userData.Name,
	)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
