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

type VideoRepo struct {
}

func (s *VideoRepo) Create(ctx context.Context, video domain.Video) error {
	return global.DB.WithContext(ctx).Create(&video).Error
}

func (s *VideoRepo) Delete(ctx context.Context, video domain.Video) error {
	return global.DB.WithContext(ctx).Delete(&video).Error
}

func (s *VideoRepo) DeleteByIds(ctx context.Context, ids request.Ids) error {
	return global.DB.WithContext(ctx).Delete(&[]domain.Video{}, ids.Ids).Error
}

func (s *VideoRepo) Update(ctx context.Context, video domain.Video) error {
	if video.ID == 0 {
		return errors.New(fmt.Sprintf("missing %s.id", "video"))
	}
	return global.DB.WithContext(ctx).Model(&video).Scopes(scopes.UpdatesAllOmit()).Updates(&video).Error
}

func (s *VideoRepo) Find(ctx context.Context, video domain.Video) (domain.Video, error) {
	db := global.DB.WithContext(ctx).Model(&domain.Video{})
	// TODO：条件过滤

	res := db.First(&video)

	return video, res.Error
}

func (s *VideoRepo) List(ctx context.Context, page domain.PageVideoSearch) ([]domain.Video, int64, error) {
	var (
		videoList []domain.Video
		count     int64
		err       error
	)
	// db
	db := global.DB.WithContext(ctx).Model(&domain.Video{})

	// TODO：条件过滤

	err = db.Count(&count).Scopes(scopes.Paginate(page.PageSearch)).Find(&videoList).Error

	return videoList, count, err
}
