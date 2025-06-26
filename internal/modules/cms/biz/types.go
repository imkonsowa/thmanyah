package biz

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type RegisterRequest struct {
	Email                string
	Password             string
	PasswordConfirmation string
	Name                 string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`

	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UpdateUserRequest struct {
	Name  string
	Email string
}

type Metadata map[string]string

func (m Metadata) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Metadata) Scan(value interface{}) error {
	if value == nil {
		*m = make(Metadata)
		return nil
	}

	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into Metadata", value)
	}

	if len(data) == 0 || string(data) == "{}" || string(data) == "null" {
		*m = make(Metadata)
		return nil
	}

	return json.Unmarshal(data, m)
}

type CategoryType string
type ProgramStatus string
type EpisodeStatus string
type ImportStatus string

const (
	CategoryTypePodcast       CategoryType = "CATEGORY_TYPE_PODCAST"
	CategoryTypeDocumentary   CategoryType = "CATEGORY_TYPE_DOCUMENTARY"
	CategoryTypeSportsEvent   CategoryType = "CATEGORY_TYPE_SPORTS_EVENT"
	CategoryTypeEducational   CategoryType = "CATEGORY_TYPE_EDUCATIONAL"
	CategoryTypeNews          CategoryType = "CATEGORY_TYPE_NEWS"
	CategoryTypeEntertainment CategoryType = "CATEGORY_TYPE_ENTERTAINMENT"
)

const (
	ProgramStatusDraft     ProgramStatus = "PROGRAM_STATUS_DRAFT"
	ProgramStatusPublished ProgramStatus = "PROGRAM_STATUS_PUBLISHED"
	ProgramStatusArchived  ProgramStatus = "PROGRAM_STATUS_ARCHIVED"
)

const (
	EpisodeStatusDraft     EpisodeStatus = "EPISODE_STATUS_DRAFT"
	EpisodeStatusPublished EpisodeStatus = "EPISODE_STATUS_PUBLISHED"
	EpisodeStatusScheduled EpisodeStatus = "EPISODE_STATUS_SCHEDULED"
	EpisodeStatusArchived  EpisodeStatus = "EPISODE_STATUS_ARCHIVED"
)

const (
	ImportStatusPending    ImportStatus = "IMPORT_STATUS_PENDING"
	ImportStatusProcessing ImportStatus = "IMPORT_STATUS_PROCESSING"
	ImportStatusCompleted  ImportStatus = "IMPORT_STATUS_COMPLETED"
	ImportStatusFailed     ImportStatus = "IMPORT_STATUS_FAILED"
)

type Category struct {
	ID          uuid.UUID    `db:"id"`
	Name        string       `db:"name"`
	Description string       `db:"description"`
	Type        CategoryType `db:"type"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
	CreatedBy   uuid.UUID    `db:"created_by"`
	Metadata    Metadata     `db:"metadata"`
}

type UpdateCategoryRequest struct {
	Name        *string       `json:"name,omitempty"`
	Description *string       `json:"description,omitempty"`
	Type        *CategoryType `json:"type,omitempty"`
	Metadata    *Metadata     `json:"metadata,omitempty"`
}

type Program struct {
	ID            uuid.UUID     `db:"id"`
	Title         string        `db:"title"`
	Description   string        `db:"description"`
	CategoryID    uuid.UUID     `db:"category_id"`
	Status        ProgramStatus `db:"status"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
	PublishedAt   *time.Time    `db:"published_at"`
	CreatedBy     uuid.UUID     `db:"created_by"`
	UpdatedBy     uuid.UUID     `db:"updated_by"`
	ThumbnailURL  string        `db:"thumbnail_url"`
	Tags          []string      `db:"tags"`
	Metadata      Metadata      `db:"metadata"`
	SourceURL     *string       `db:"source_url"`
	EpisodesCount int32         `db:"episodes_count"`
	IsFeatured    bool          `db:"is_featured"`
	ViewCount     int32         `db:"view_count"`
	Rating        float64       `db:"rating"`
}

type UpdateProgramRequest struct {
	Title         *string        `json:"title,omitempty"`
	Description   *string        `json:"description,omitempty"`
	CategoryID    *uuid.UUID     `json:"category_id,omitempty"`
	Status        *ProgramStatus `json:"status,omitempty"`
	PublishedAt   *time.Time     `json:"published_at,omitempty"`
	ThumbnailURL  *string        `json:"thumbnail_url,omitempty"`
	Tags          *[]string      `json:"tags,omitempty"`
	Metadata      *Metadata      `json:"metadata,omitempty"`
	SourceURL     *string        `json:"source_url,omitempty"`
	EpisodesCount *int32         `json:"episodes_count,omitempty"`
	IsFeatured    *bool          `json:"is_featured,omitempty"`
	ViewCount     *int32         `json:"view_count,omitempty"`
	Rating        *float64       `json:"rating,omitempty"`
}

type Episode struct {
	ID            uuid.UUID     `db:"id"`
	ProgramID     uuid.UUID     `db:"program_id"`
	Title         string        `db:"title"`
	Description   string        `db:"description"`
	DurationSecs  int32         `db:"duration_seconds"`
	EpisodeNumber int32         `db:"episode_number"`
	SeasonNumber  int32         `db:"season_number"`
	Status        EpisodeStatus `db:"status"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
	PublishedAt   *time.Time    `db:"published_at"`
	ScheduledAt   *time.Time    `db:"scheduled_at"`
	CreatedBy     uuid.UUID     `db:"created_by"`
	UpdatedBy     uuid.UUID     `db:"updated_by"`
	MediaURL      string        `db:"media_url"`
	ThumbnailURL  string        `db:"thumbnail_url"`
	Tags          []string      `db:"tags"`
	Metadata      Metadata      `db:"metadata"`
	ViewCount     int32         `db:"view_count"`
	Rating        float64       `db:"rating"`
}

// UpdateEpisodeRequest contains only the fields that are safe to update
type UpdateEpisodeRequest struct {
	Title         *string        `json:"title,omitempty"`
	Description   *string        `json:"description,omitempty"`
	DurationSecs  *int32         `json:"duration_seconds,omitempty"`
	EpisodeNumber *int32         `json:"episode_number,omitempty"`
	SeasonNumber  *int32         `json:"season_number,omitempty"`
	Status        *EpisodeStatus `json:"status,omitempty"`
	PublishedAt   *time.Time     `json:"published_at,omitempty"`
	ScheduledAt   *time.Time     `json:"scheduled_at,omitempty"`
	MediaURL      *string        `json:"media_url,omitempty"`
	ThumbnailURL  *string        `json:"thumbnail_url,omitempty"`
	Tags          *[]string      `json:"tags,omitempty"`
	Metadata      *Metadata      `json:"metadata,omitempty"`
	ViewCount     *int32         `json:"view_count,omitempty"`
	Rating        *float64       `json:"rating,omitempty"`
}

// BulkUpdateProgramsRequest contains only the fields that are safe to update in bulk operations
type BulkUpdateProgramsRequest struct {
	Status     *ProgramStatus `json:"status,omitempty"`
	CategoryID *uuid.UUID     `json:"category_id,omitempty"`
	Tags       *[]string      `json:"tags,omitempty"`
	Metadata   *Metadata      `json:"metadata,omitempty"`
	IsFeatured *bool          `json:"is_featured,omitempty"`
}

type PaginationRequest struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

func (p *PaginationRequest) SetDefaults() {
	if p.PageSize == 0 {
		p.PageSize = 20
	}
	if p.Page == 0 {
		p.Page = 1
	}
}

type SortRequest struct {
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

type PaginationResponse struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"page_size"`
	TotalCount int32 `json:"total_count"`
	TotalPages int32 `json:"total_pages"`
}

type ProgramFilter struct {
	CategoryID   *uuid.UUID     `json:"category_id"`
	Status       *ProgramStatus `json:"status"`
	SearchQuery  *string        `json:"search_query"`
	Tags         []string       `json:"tags"`
	FeaturedOnly *bool          `json:"featured_only"`
	CreatedBy    *uuid.UUID     `json:"created_by"`
}

type CategoryFilter struct {
	Type        *CategoryType `json:"type"`
	SearchQuery *string       `json:"search_query"`
	CreatedBy   *uuid.UUID    `json:"created_by"`
}

type EpisodeFilter struct {
	ProgramID   *uuid.UUID     `json:"program_id"`
	Status      *EpisodeStatus `json:"status"`
	SearchQuery *string        `json:"search_query"`
	CreatedBy   *uuid.UUID     `json:"created_by"`
}

type ImportFilter struct {
	Status      *ImportStatus `json:"status"`
	SearchQuery *string       `json:"search_query"`
	CreatedBy   *uuid.UUID    `json:"created_by"`
}

type ImportData struct {
	ID             uuid.UUID    `db:"id"`
	SourceType     string       `db:"source_type"`
	SourceURL      string       `db:"source_url"`
	SourceConfig   Metadata     `db:"source_config"`
	CategoryID     uuid.UUID    `db:"category_id"`
	Status         ImportStatus `db:"status"`
	TotalItems     int32        `db:"total_items"`
	ProcessedItems int32        `db:"processed_items"`
	SuccessCount   int32        `db:"success_count"`
	ErrorCount     int32        `db:"error_count"`
	Errors         []string     `db:"errors"`
	Warnings       []string     `db:"warnings"`
	CreatedAt      time.Time    `db:"created_at"`
	UpdatedAt      time.Time    `db:"updated_at"`
	CreatedBy      uuid.UUID    `db:"created_by"`
	FieldMapping   Metadata     `db:"field_mapping"`
}

// UpdateImportRequest contains only the fields that are safe to update
type UpdateImportRequest struct {
	SourceType     *string       `json:"source_type,omitempty"`
	SourceURL      *string       `json:"source_url,omitempty"`
	SourceConfig   *Metadata     `json:"source_config,omitempty"`
	CategoryID     *uuid.UUID    `json:"category_id,omitempty"`
	Status         *ImportStatus `json:"status,omitempty"`
	TotalItems     *int32        `json:"total_items,omitempty"`
	ProcessedItems *int32        `json:"processed_items,omitempty"`
	SuccessCount   *int32        `json:"success_count,omitempty"`
	ErrorCount     *int32        `json:"error_count,omitempty"`
	Errors         *[]string     `json:"errors,omitempty"`
	Warnings       *[]string     `json:"warnings,omitempty"`
	FieldMapping   *Metadata     `json:"field_mapping,omitempty"`
}

type UpdateImportProgressRequest struct {
	ProcessedItems *int32 `json:"processed_items,omitempty"`
	SuccessCount   *int32 `json:"success_count,omitempty"`
	ErrorCount     *int32 `json:"error_count,omitempty"`
}

type UpdateEpisodeFileRequest struct {
	Target string
	File   multipart.File
	Header *multipart.FileHeader
}
