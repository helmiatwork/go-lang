package service

import (
	"belajar-go/internal/domain"
	"context"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, u *domain.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.repo.Create(ctx, u)
}

func (s *UserService) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) List(ctx context.Context) ([]domain.User, error) {
	return s.repo.List(ctx)
}

func (s *UserService) Update(ctx context.Context, u *domain.User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	return s.repo.Update(ctx, u)
}

func (s *UserService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
