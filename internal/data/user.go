package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/scopes"
)

type UserRepo struct {
}

func (s *UserRepo) Create(ctx context.Context, user domain.User) error {
	return global.DB.WithContext(ctx).Create(&user).Error
}

func (s *UserRepo) Delete(ctx context.Context, user domain.User) error {
	return global.DB.WithContext(ctx).Delete(&user).Error
}

func (s *UserRepo) DeleteByIds(ctx context.Context, ids request.Ids) error {
	return global.DB.WithContext(ctx).Delete(&[]domain.User{}, ids.Ids).Error
}

func (s *UserRepo) Update(ctx context.Context, user domain.User) error {
	if user.ID == 0 {
		return errors.New(fmt.Sprintf("missing %s.id", "user"))
	}
	return global.DB.WithContext(ctx).Model(&user).Scopes(scopes.UpdatesAllOmit()).Updates(&user).Error
}

func (s *UserRepo) Find(ctx context.Context, user domain.User) (domain.User, error) {
	db := global.DB.WithContext(ctx).Model(&domain.User{})
	// TODO：条件过滤

	res := db.First(&user)

	return user, res.Error
}

func (s *UserRepo) List(ctx context.Context, page domain.PageUserSearch) ([]domain.User, int64, error) {
	var (
		userList []domain.User
		count    int64
		err      error
	)
	// db
	db := global.DB.WithContext(ctx).Model(&domain.User{})

	// TODO：条件过滤

	err = db.Count(&count).Scopes(scopes.Paginate(page.PageSearch)).Find(&userList).Error

	return userList, count, err
}
