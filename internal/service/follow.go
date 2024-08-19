package service

import (
	"context"
	"github.com/Madou-Shinni/gin-quickstart/internal/data"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type FollowRepo interface {
	Create(ctx context.Context, follow domain.Follow) error
	Delete(ctx context.Context, follow domain.Follow) error
	Update(ctx context.Context, follow domain.Follow) error
	Find(ctx context.Context, follow domain.Follow) (domain.Follow, error)
	List(ctx context.Context, page domain.PageFollowSearch) ([]domain.Follow, int64, error)
	DeleteByIds(ctx context.Context, ids request.Ids) error
}

type FollowService struct {
	repo FollowRepo
}

func NewFollowService() *FollowService {
	return &FollowService{repo: &data.FollowRepo{}}
}

func (s *FollowService) Add(ctx context.Context, follow domain.Follow) error {
	// 3.持久化入库
	if err := s.repo.Create(ctx, follow); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(follow)", zap.Error(err), zap.Any("domain.Follow", follow))
		return err
	}

	return nil
}

func (s *FollowService) Delete(ctx context.Context, follow domain.Follow) error {
	if err := s.repo.Delete(ctx, follow); err != nil {
		logger.Error("s.repo.Delete(follow)", zap.Error(err), zap.Any("domain.Follow", follow))
		return err
	}

	return nil
}

func (s *FollowService) Update(ctx context.Context, follow domain.Follow) error {
	if err := s.repo.Update(ctx, follow); err != nil {
		logger.Error("s.repo.Update(follow)", zap.Error(err), zap.Any("domain.Follow", follow))
		return err
	}

	return nil
}

func (s *FollowService) Find(ctx context.Context, follow domain.Follow) (domain.Follow, error) {
	res, err := s.repo.Find(ctx, follow)

	if err != nil {
		logger.Error("s.repo.Find(follow)", zap.Error(err), zap.Any("domain.Follow", follow))
		return res, err
	}

	return res, nil
}

func (s *FollowService) List(ctx context.Context, page domain.PageFollowSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, count, err := s.repo.List(ctx, page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageFollowSearch", page))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *FollowService) DeleteByIds(ctx context.Context, ids request.Ids) error {
	if err := s.repo.DeleteByIds(ctx, ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
