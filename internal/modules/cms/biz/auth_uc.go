package biz

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"thmanyah/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (uc *UseCase) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	user, err := uc.usersRepo.GetUserWithPassword(ctx, request.Email)
	if err != nil {
		uc.logger.Errorw("msg", "find user db error", "err", err)

		return nil, err
	}

	if user == nil {
		return nil, ErrInvalidCredentials
	}

	password := request.Password
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	if user.Password != hashedPassword {
		return nil, ErrInvalidCredentials
	}

	claimsMap := utils.NewClaimsBuilder().
		WithUserID(user.ID.String()).
		WithExpiry(time.Now().Add(time.Hour * 72).Unix()).
		Build()

	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, claimsMap)

	signedString, err := claims.SignedString(uc.keysStore.PrivateKey())

	if err != nil {
		return nil, fmt.Errorf("generate token failed: %s", err.Error())
	}

	return &LoginResponse{
		Token: signedString,
		User:  user,
	}, nil
}

func (uc *UseCase) Register(ctx context.Context, req *RegisterRequest) error {
	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	_, err := uc.usersRepo.CreateUser(ctx, &User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetUserProfile(ctx context.Context, userId uuid.UUID) (*User, error) {
	user, err := uc.usersRepo.GetUserByIdentifier(ctx, userId.String())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) UpdateUserProfile(ctx context.Context, userId uuid.UUID, req *UpdateUserRequest) (*User, error) {
	user, err := uc.usersRepo.GetUserByIdentifier(ctx, userId.String())
	if err != nil {
		return nil, err
	}

	user, err = uc.usersRepo.UpdateUser(ctx, user.ID, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
