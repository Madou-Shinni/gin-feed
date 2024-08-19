package handle

import (
	"errors"
	"github.com/Madou-Shinni/gin-quickstart/internal/domain"
	"github.com/Madou-Shinni/gin-quickstart/internal/service"
	"github.com/Madou-Shinni/gin-quickstart/pkg/constant"
	"github.com/Madou-Shinni/gin-quickstart/pkg/request"
	"github.com/Madou-Shinni/gin-quickstart/pkg/response"
	"github.com/Madou-Shinni/gin-quickstart/pkg/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FollowHandle struct {
	s *service.FollowService
}

func NewFollowHandle() *FollowHandle {
	return &FollowHandle{s: service.NewFollowService()}
}

// Add 创建Follow
// @Tags     Follow
// @Summary  创建Follow
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Follow true "创建Follow"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /follow [post]
func (cl *FollowHandle) Add(c *gin.Context) {
	var follow domain.Follow
	if err := c.ShouldBindJSON(&follow); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			response.Error(c, constant.CODE_INVALID_PARAMETER, tools.TransErrs(errs))
			return
		}
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(c.Request.Context(), follow); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Follow
// @Tags     Follow
// @Summary  删除Follow
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Follow true "删除Follow"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /follow [delete]
func (cl *FollowHandle) Delete(c *gin.Context) {
	var follow domain.Follow
	if err := c.ShouldBindJSON(&follow); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(c.Request.Context(), follow); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Follow
// @Tags     Follow
// @Summary  批量删除Follow
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     request.Ids true "批量删除Follow"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /follow/delete-batch [delete]
func (cl *FollowHandle) DeleteByIds(c *gin.Context) {
	var ids request.Ids
	if err := c.ShouldBindJSON(&ids); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.DeleteByIds(c.Request.Context(), ids); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Update 修改Follow
// @Tags     Follow
// @Summary  修改Follow
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Follow true "修改Follow"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /follow [put]
func (cl *FollowHandle) Update(c *gin.Context) {
	var follow domain.Follow
	if err := c.ShouldBindJSON(&follow); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(c.Request.Context(), follow); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Follow
// @Tags     Follow
// @Summary  查询Follow
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id path     uint true "查询Follow"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /follow/{id} [get]
func (cl *FollowHandle) Find(c *gin.Context) {
	var follow domain.Follow
	if err := c.ShouldBindUri(&follow); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(c.Request.Context(), follow)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Follow列表
// @Tags     Follow
// @Summary  查询Follow列表
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query     domain.PageFollowSearch true "查询Follow列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /follow/list [get]
func (cl *FollowHandle) List(c *gin.Context) {
	var follow domain.PageFollowSearch
	if err := c.ShouldBindQuery(&follow); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(c.Request.Context(), follow)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
