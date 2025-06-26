package utils

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	jwt2 "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ClaimsBuilder struct {
	userID string
	expiry int64
}

func NewClaimsBuilder() *ClaimsBuilder {
	return &ClaimsBuilder{}
}

func (c *ClaimsBuilder) WithUserID(userID string) *ClaimsBuilder {
	c.userID = userID
	return c
}

func (c *ClaimsBuilder) WithExpiry(expiry int64) *ClaimsBuilder {
	c.expiry = expiry
	return c
}

func (c *ClaimsBuilder) Build() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": c.userID,
	}
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	claims, ok := jwt2.FromContext(ctx)
	if !ok {
		return uuid.Nil, errors.Unauthorized("unauthorized", "unauthorized")
	}

	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.Unauthorized("unauthorized", "unauthorized")
	}

	if claimsMap == nil {
		return uuid.Nil, errors.Unauthorized("unauthorized", "unauthorized")
	}

	userId, ok := claimsMap["user_id"]
	if !ok {
		return uuid.Nil, errors.Unauthorized("unauthorized", "unauthorized")
	}

	if userId == "" {
		return uuid.Nil, errors.Unauthorized("unauthorized", "unauthorized")
	}

	return uuid.Parse(userId.(string))
}
