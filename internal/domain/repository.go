package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
	List(ctx context.Context) ([]User, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uint) error
}
