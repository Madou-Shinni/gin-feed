package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Madou-Shinni/gin-quickstart/constants"
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/global"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools/pagelimit"
	"github.com/Madou-Shinni/go-logger"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// 定义接口
type MovingRepo interface {
	Create(ctx context.Context, moving domain.Moving) error
	Delete(ctx context.Context, moving domain.Moving) error
	Update(ctx context.Context, moving domain.Moving) error
	Find(ctx context.Context, moving domain.Moving) (domain.Moving, error)
	List(ctx context.Context, page domain.PageMovingSearch) ([]domain.Moving, int64, error)
	DeleteByIds(ctx context.Context, ids request.Ids) error
}

type MovingService struct {
	repo MovingRepo
}

func NewMovingService() *MovingService {
	return &MovingService{repo: &data.MovingRepo{}}
}

func (s *MovingService) Add(ctx context.Context, moving domain.Moving) error {
	// 3.持久化入库
	if err := s.repo.Create(ctx, moving); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(moving)", zap.Error(err), zap.Any("domain.Moving", moving))
		return err
	}

	return nil
}

func (s *MovingService) Delete(ctx context.Context, moving domain.Moving) error {
	if err := s.repo.Delete(ctx, moving); err != nil {
		logger.Error("s.repo.Delete(moving)", zap.Error(err), zap.Any("domain.Moving", moving))
		return err
	}

	return nil
}

func (s *MovingService) Update(ctx context.Context, moving domain.Moving) error {
	if err := s.repo.Update(ctx, moving); err != nil {
		logger.Error("s.repo.Update(moving)", zap.Error(err), zap.Any("domain.Moving", moving))
		return err
	}

	return nil
}

func (s *MovingService) Find(ctx context.Context, moving domain.Moving) (domain.Moving, error) {
	res, err := s.repo.Find(ctx, moving)

	if err != nil {
		logger.Error("s.repo.Find(moving)", zap.Error(err), zap.Any("domain.Moving", moving))
		return res, err
	}

	return res, nil
}

func (s *MovingService) List(ctx context.Context, page domain.PageMovingSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
		data    []domain.Moving
		count   int64
	)

	offset, l := pagelimit.OffsetLimit(page.PageNum, page.PageSize)

	// 查询目标是大up还是小up
	i, err := global.Rdb.HGet(constants.HFollowed, constants.GetFollowedKey(page.Uid)).Int()
	if err != nil && !errors.Is(err, redis.Nil) {
		return pageRes, err
	}

	if i > constants.GetDefaultFollower() {
		// 大up读扩散
		count = global.Rdb.ZCard(constants.GetMovingListKey(page.Uid)).Val()
		resp := global.Rdb.ZRevRangeWithScores(constants.GetMovingListKey(page.Uid), int64(offset), int64(l)).Val()
		if count < int64(l) {
			count = count + global.Rdb.ZCard(constants.GetMovingViewsKey(page.Uid)).Val()
			resp = append(resp, global.Rdb.ZRevRangeWithScores(constants.GetMovingViewsKey(page.Uid), 0, int64(l)-count).Val()...)
		}
		for _, z := range resp {
			var d domain.Moving
			json.Unmarshal([]byte(z.Member.(string)), &d)
			data = append(data, d)
		}
	} else {
		// 写扩散
		count = global.Rdb.ZCard(constants.GetMovingViewsKey(page.Uid)).Val()
		resp := global.Rdb.ZRevRangeWithScores(constants.GetMovingViewsKey(page.Uid), int64(offset), int64(l)).Val()
		for _, z := range resp {
			var d domain.Moving
			json.Unmarshal([]byte(z.Member.(string)), &d)
			data = append(data, d)
		}
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *MovingService) DeleteByIds(ctx context.Context, ids request.Ids) error {
	if err := s.repo.DeleteByIds(ctx, ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
