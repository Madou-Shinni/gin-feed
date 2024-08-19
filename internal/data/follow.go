package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/scopes"
	"gorm.io/gorm"
)

type FollowRepo struct {
}

func (s *FollowRepo) Create(ctx context.Context, follow domain.Follow) error {
	var err error
	err = global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&follow).Error
		if err != nil {
			return err
		}

		err = global.Rdb.HIncrBy(constants.HFollowed, constants.GetFollowedKey(follow.Followed), 1).Err()
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (s *FollowRepo) Delete(ctx context.Context, follow domain.Follow) error {
	return global.DB.WithContext(ctx).Delete(&follow).Error
}

func (s *FollowRepo) DeleteByIds(ctx context.Context, ids request.Ids) error {
	return global.DB.WithContext(ctx).Delete(&[]domain.Follow{}, ids.Ids).Error
}

func (s *FollowRepo) Update(ctx context.Context, follow domain.Follow) error {
	if follow.ID == 0 {
		return errors.New(fmt.Sprintf("missing %s.id", "follow"))
	}
	return global.DB.WithContext(ctx).Model(&follow).Scopes(scopes.UpdatesAllOmit()).Updates(&follow).Error
}

func (s *FollowRepo) Find(ctx context.Context, follow domain.Follow) (domain.Follow, error) {
	db := global.DB.WithContext(ctx).Model(&domain.Follow{})
	// TODO：条件过滤

	res := db.First(&follow)

	return follow, res.Error
}

func (s *FollowRepo) List(ctx context.Context, page domain.PageFollowSearch) ([]domain.Follow, int64, error) {
	var (
		followList []domain.Follow
		count      int64
		err        error
	)
	// db
	db := global.DB.WithContext(ctx).Model(&domain.Follow{})

	// TODO：条件过滤

	err = db.Count(&count).Scopes(scopes.Paginate(page.PageSearch)).Find(&followList).Error

	return followList, count, err
}
