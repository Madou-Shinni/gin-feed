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
type UserRepo interface {
	Create(ctx context.Context, user domain.User) error
	Delete(ctx context.Context, user domain.User) error
	Update(ctx context.Context, user domain.User) error
	Find(ctx context.Context, user domain.User) (domain.User, error)
	List(ctx context.Context, page domain.PageUserSearch) ([]domain.User, int64, error)
	DeleteByIds(ctx context.Context, ids request.Ids) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService() *UserService {
	return &UserService{repo: &data.UserRepo{}}
}

func (s *UserService) Add(ctx context.Context, user domain.User) error {
	// 3.持久化入库
	if err := s.repo.Create(ctx, user); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, user domain.User) error {
	if err := s.repo.Delete(ctx, user); err != nil {
		logger.Error("s.repo.Delete(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Update(ctx context.Context, user domain.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		logger.Error("s.repo.Update(user)", zap.Error(err), zap.Any("domain.User", user))
		return err
	}

	return nil
}

func (s *UserService) Find(ctx context.Context, user domain.User) (domain.User, error) {
	res, err := s.repo.Find(ctx, user)

	if err != nil {
		logger.Error("s.repo.Find(user)", zap.Error(err), zap.Any("domain.User", user))
		return res, err
	}

	return res, nil
}

func (s *UserService) List(ctx context.Context, page domain.PageUserSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, count, err := s.repo.List(ctx, page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageUserSearch", page))
		return pageRes, err
	}

	pageRes.List = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *UserService) DeleteByIds(ctx context.Context, ids request.Ids) error {
	if err := s.repo.DeleteByIds(ctx, ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
