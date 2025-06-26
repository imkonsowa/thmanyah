package biz

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/google/uuid"
)

type UsersRepository interface {
	GetUserWithPassword(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUserByIdentifier(ctx context.Context, identifier string) (*User, error)
	UpdateUser(ctx context.Context, userId uuid.UUID, user *UpdateUserRequest) (*User, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	Update(ctx context.Context, userID, id uuid.UUID, updates *UpdateCategoryRequest) (*Category, error)
	Delete(ctx context.Context, userID, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Category, error)
	List(ctx context.Context, filter CategoryFilter, pagination PaginationRequest, sort SortRequest) ([]*Category, *PaginationResponse, error)
}

type ProgramRepository interface {
	Create(ctx context.Context, program *Program) error
	Update(ctx context.Context, userID, id uuid.UUID, updates *UpdateProgramRequest) (*Program, error)
	Delete(ctx context.Context, userID, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Program, error)
	List(ctx context.Context, filter ProgramFilter, pagination PaginationRequest, sort SortRequest) ([]*Program, *PaginationResponse, error)
	BulkUpdate(ctx context.Context, userID uuid.UUID, ids []uuid.UUID, updates *BulkUpdateProgramsRequest) (int32, error)
	BulkDelete(ctx context.Context, userID uuid.UUID, ids []uuid.UUID) error
	IncrementViewCount(ctx context.Context, id uuid.UUID) error
	UpdateEpisodesCount(ctx context.Context, programID uuid.UUID) error
}

type EpisodeRepository interface {
	Create(ctx context.Context, episode *Episode) error
	Update(ctx context.Context, userID, id uuid.UUID, updates *UpdateEpisodeRequest) (*Episode, error)
	Delete(ctx context.Context, userID, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*Episode, error)
	List(ctx context.Context, filter EpisodeFilter, pagination PaginationRequest, sort SortRequest) ([]*Episode, *PaginationResponse, error)
	ListByProgram(ctx context.Context, programID uuid.UUID, pagination PaginationRequest, sort SortRequest) ([]*Episode, *PaginationResponse, error)
	IncrementViewCount(ctx context.Context, id uuid.UUID) error
}

type ImportRepository interface {
	Create(ctx context.Context, importData *ImportData) error
	Update(ctx context.Context, userID, id uuid.UUID, updates *UpdateImportRequest) (*ImportData, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ImportData, error)
	List(ctx context.Context, filter ImportFilter, pagination PaginationRequest, sort SortRequest) ([]*ImportData, *PaginationResponse, error)
	UpdateProgress(ctx context.Context, userID, id uuid.UUID, updates *UpdateImportProgressRequest) error
	AddError(ctx context.Context, userID, id uuid.UUID, errorMsg string) error
	AddWarning(ctx context.Context, userID, id uuid.UUID, warningMsg string) error
}

type S3Client interface {
	GetObject(ctx context.Context, bucket, key string) (io.Reader, error)
	PutObject(ctx context.Context, bucket, key string, file multipart.File) error
	DeleteObject(ctx context.Context, bucket, key string) error
	GetObjectSignedURL(ctx context.Context, bucket, key string) (string, error)
	GetObjectPublicURL(ctx context.Context, bucket, key string) string
}
