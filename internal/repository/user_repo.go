package repository

import (
	"context"
	"quanfuxia/internal/model/gens"
	"quanfuxia/internal/model/query"
)

type UserRepo interface {
	Create(ctx context.Context, user *gens.WaUser) error
	FindByUsername(ctx context.Context, username string) (*gens.WaUser, error)
}

type userRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &userRepoImpl{}
}

func (r *userRepoImpl) Create(ctx context.Context, user *gens.WaUser) error {
	return query.Q.WaUser.WithContext(ctx).Create(user)
}

func (r *userRepoImpl) FindByUsername(ctx context.Context, username string) (*gens.WaUser, error) {
	return query.Q.WaUser.WithContext(ctx).Where(query.Q.WaUser.Username.Eq(username)).First()
}
