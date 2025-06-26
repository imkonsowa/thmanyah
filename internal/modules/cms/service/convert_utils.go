package service

import (
	v1 "thmanyah/api/grpc/v1"
	"thmanyah/internal/modules/cms/biz"
)

func convertFullUser(user *biz.User) *v1.User {
	return &v1.User{
		Id:        user.ID.String(),
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		Name:      user.Name,
		Email:     user.Email,
	}
}
