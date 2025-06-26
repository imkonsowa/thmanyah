package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var ErrUserNotFound = errors.NotFound("USER_NOT_FOUND", "user not found")
var ErrInvalidCredentials = errors.Unauthorized("INVALID_CREDENTIALS", "invalid credentials")
var ErrUnauthorized = errors.Unauthorized("UNAUTHORIZED", "unauthorized")
var ErrUserAlreadyExists = errors.BadRequest("USER_ALREADY_EXISTS", "user already exists")
var ErrCategoryAlreadyExists = errors.BadRequest("CATEGORY_ALREADY_EXISTS", "category name already exists")
var ErrCategoryNotFound = errors.NotFound("CATEGORY_NOT_FOUND", "category not found")
var ErrProgramNotFound = errors.NotFound("PROGRAM_NOT_FOUND", "program not found")
var ErrEpisodeNotFound = errors.NotFound("EPISODE_NOT_FOUND", "episode not found")
var ErrEpisodeAlreadyExists = errors.BadRequest("EPISODE_ALREADY_EXISTS", "episode with this number already exists for this program and season")
