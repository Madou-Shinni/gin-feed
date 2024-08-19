package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Madou-Shinni/gin-quickstart/common"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/scopes"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type MovingRepo struct {
}

func (s *MovingRepo) Create(ctx context.Context, moving domain.Moving) error {
	err := global.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 创建动态
		err := tx.Create(&moving).Error
		if err != nil {
			return err
		}

		// 发布动态
		// 查询粉丝数量
		i, err := global.Rdb.HGet(constants.HFollowed, constants.GetFollowedKey(moving.Uid)).Int()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		movingMarshal, _ := json.Marshal(moving)
		if i > constants.GetDefaultFollower() {
			// 大up 读扩散
			err = global.Rdb.ZAdd(constants.GetMovingListKey(moving.Uid), redis.Z{
				Score:  0,
				Member: movingMarshal,
			}).Err()
			if err != nil {
				return err
			}
		} else {
			// 小up 写扩散
			// 查询所有粉丝
			var fans []uint
			var keys []string
			err = tx.Model(&domain.Follow{}).Select("follower").Find(&fans, "followed = ?", moving.Uid).Error
			if err != nil {
				return err
			}
			for _, v := range fans {
				keys = append(keys, constants.GetMovingViewsKey(v))
			}

			err = global.Rdb.Eval(common.AddMovingLua, keys, []string{"0", string(movingMarshal)}).Err()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (s *MovingRepo) Delete(ctx context.Context, moving domain.Moving) error {
	return global.DB.WithContext(ctx).Delete(&moving).Error
}

func (s *MovingRepo) DeleteByIds(ctx context.Context, ids request.Ids) error {
	return global.DB.WithContext(ctx).Delete(&[]domain.Moving{}, ids.Ids).Error
}

func (s *MovingRepo) Update(ctx context.Context, moving domain.Moving) error {
	if moving.ID == 0 {
		return errors.New(fmt.Sprintf("missing %s.id", "moving"))
	}
	return global.DB.WithContext(ctx).Model(&moving).Scopes(scopes.UpdatesAllOmit()).Updates(&moving).Error
}

func (s *MovingRepo) Find(ctx context.Context, moving domain.Moving) (domain.Moving, error) {
	db := global.DB.WithContext(ctx).Model(&domain.Moving{})
	// TODO：条件过滤

	res := db.First(&moving)

	return moving, res.Error
}

func (s *MovingRepo) List(ctx context.Context, page domain.PageMovingSearch) ([]domain.Moving, int64, error) {
	var (
		movingList []domain.Moving
		count      int64
		err        error
	)
	// db
	db := global.DB.WithContext(ctx).Model(&domain.Moving{})

	// TODO：条件过滤

	err = db.Count(&count).Scopes(scopes.Paginate(page.PageSearch)).Find(&movingList).Error

	return movingList, count, err
}
