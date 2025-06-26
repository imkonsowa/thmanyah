package service

import (
	"context"

	v1 "thmanyah/api/grpc/v1"
	"thmanyah/internal/modules/cms/biz"
	"thmanyah/internal/utils"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer

	uc *biz.UseCase
}

func NewAuthService(uc *biz.UseCase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) Login(
	ctx context.Context,
	req *v1.LoginRequest,
) (*v1.LoginResponse, error) {
	response, err := s.uc.Login(ctx, &biz.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &v1.LoginResponse{
		AccessToken: response.Token,
		User:        convertFullUser(response.User),
	}, nil
}

func (s *AuthService) Register(
	ctx context.Context,
	req *v1.RegisterRequest,
) (*v1.RegisterResponse, error) {
	err := s.uc.Register(ctx, &biz.RegisterRequest{
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.ConfirmPassword,
		Name:                 req.Name,
	})
	if err != nil {
		return nil, err
	}

	token, err := s.uc.Login(ctx, &biz.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &v1.RegisterResponse{
		AccessToken: token.Token,
		User:        convertFullUser(token.User),
	}, nil
}

func (s *AuthService) GetUserProfile(ctx context.Context, req *emptypb.Empty) (*v1.UserProfileResponse, error) {
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.GetUserProfile(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.UserProfileResponse{
		User: convertFullUser(user),
	}, nil
}

func (s *AuthService) UpdateUserProfile(
	ctx context.Context,
	req *v1.UpdateUserRequest,
) (*v1.UpdateUserResponse, error) {
	userId, err := utils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.uc.UpdateUserProfile(ctx, userId, &biz.UpdateUserRequest{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserResponse{User: convertFullUser(user)}, nil
}
