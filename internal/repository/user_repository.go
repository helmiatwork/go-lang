package repository

import (
	"belajar-go/internal/domain"
	"context"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		if isDuplicate(err) {
			return domain.ErrDuplicateEmail
		}

		return err
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	res := r.db.WithContext(ctx).Delete(&domain.User{}, id)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r *userRepository) Update(ctx context.Context, u *domain.User) error {
	res := r.db.WithContext(ctx).Model(u).Where("id = ?", u.ID).Updates(map[string]any{"name": u.Name, "email": u.Email})

	if res.Error != nil {
		if isDuplicate(res.Error) {
			return domain.ErrDuplicateEmail
		}
		return res.Error
	}

	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var u domain.User
	err := r.db.WithContext(ctx).First(&u, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) List(ctx context.Context) ([]domain.User, error) {
	users := []domain.User{}

	if err := r.db.WithContext(ctx).Order("id").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func isDuplicate(err error) bool {
	var me *mysql.MySQLError
	return errors.As(err, &me) && me.Number == 1062
}
