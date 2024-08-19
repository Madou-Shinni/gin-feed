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
type VideoRepo interface {
	Create(ctx context.Context, video domain.Video) error
	Delete(ctx context.Context, video domain.Video) error
	Update(ctx context.Context, video domain.Video) error
	Find(ctx context.Context, video domain.Video) (domain.Video, error)
	List(ctx context.Context, page domain.PageVideoSearch) ([]domain.Video, int64, error)
	DeleteByIds(ctx context.Context, ids request.Ids) error
}

type VideoService struct {
	repo VideoRepo
}

func NewVideoService() *VideoService {
	return &VideoService{repo: &data.VideoRepo{}}
}

func (s *VideoService) Add(ctx context.Context, video domain.Video) error {
	// 3.持久化入库
	if err := s.repo.Create(ctx, video); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Delete(ctx context.Context, video domain.Video) error {
	if err := s.repo.Delete(ctx, video); err != nil {
		logger.Error("s.repo.Delete(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Update(ctx context.Context, video domain.Video) error {
	if err := s.repo.Update(ctx, video); err != nil {
		logger.Error("s.repo.Update(video)", zap.Error(err), zap.Any("domain.Video", video))
		return err
	}

	return nil
}

func (s *VideoService) Find(ctx context.Context, video domain.Video) (domain.Video, error) {
	res, err := s.repo.Find(ctx, video)

	if err != nil {
		logger.Error("s.repo.Find(video)", zap.Error(err), zap.Any("domain.Video", video))
		return res, err
	}

	return res, nil
}

func (s *VideoService) List(ctx context.Context, page domain.PageVideoSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, count, err := s.repo.List(ctx, page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageVideoSearch", page))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *VideoService) DeleteByIds(ctx context.Context, ids request.Ids) error {
	if err := s.repo.DeleteByIds(ctx, ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
