package service

import (
	"context"
	"errors"
	"quanfuxia/internal/model/gens"
	"quanfuxia/internal/repository"
)

type UserService interface {
	Register(ctx context.Context, username, password string) error
}

type userServiceImpl struct {
	userRepo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) UserService {
	return &userServiceImpl{userRepo: repo}
}

func (s *userServiceImpl) Register(ctx context.Context, username, password string) error {
	exist, _ := s.userRepo.FindByUsername(ctx, username)
	if exist != nil {
		return errors.New("用户名已存在")
	}
	user := &gens.WaUser{
		Username: username,
		Password: password, // 实际应加密
	}
	return s.userRepo.Create(ctx, user)
}
