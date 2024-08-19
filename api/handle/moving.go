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

type MovingHandle struct {
	s *service.MovingService
}

func NewMovingHandle() *MovingHandle {
	return &MovingHandle{s: service.NewMovingService()}
}

// Add 创建Moving
// @Tags     Moving
// @Summary  创建Moving
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Moving true "创建Moving"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /moving [post]
func (cl *MovingHandle) Add(c *gin.Context) {
	var moving domain.Moving
	if err := c.ShouldBindJSON(&moving); err != nil {
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			response.Error(c, constant.CODE_INVALID_PARAMETER, tools.TransErrs(errs))
			return
		}
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Add(c.Request.Context(), moving); err != nil {
		response.Error(c, constant.CODE_ADD_FAILED, constant.CODE_ADD_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Delete 删除Moving
// @Tags     Moving
// @Summary  删除Moving
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Moving true "删除Moving"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /moving [delete]
func (cl *MovingHandle) Delete(c *gin.Context) {
	var moving domain.Moving
	if err := c.ShouldBindJSON(&moving); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Delete(c.Request.Context(), moving); err != nil {
		response.Error(c, constant.CODE_DELETE_FAILED, constant.CODE_DELETE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// DeleteByIds 批量删除Moving
// @Tags     Moving
// @Summary  批量删除Moving
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     request.Ids true "批量删除Moving"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /moving/delete-batch [delete]
func (cl *MovingHandle) DeleteByIds(c *gin.Context) {
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

// Update 修改Moving
// @Tags     Moving
// @Summary  修改Moving
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data body     domain.Moving true "修改Moving"
// @Success  200  {string} string            "{"code":200,"msg":"","data":{}"}"
// @Router   /moving [put]
func (cl *MovingHandle) Update(c *gin.Context) {
	var moving domain.Moving
	if err := c.ShouldBindJSON(&moving); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	if err := cl.s.Update(c.Request.Context(), moving); err != nil {
		response.Error(c, constant.CODE_UPDATE_FAILED, constant.CODE_UPDATE_FAILED.Msg())
		return
	}

	response.Success(c)
}

// Find 查询Moving
// @Tags     Moving
// @Summary  查询Moving
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    id path     uint true "查询Moving"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /moving/{id} [get]
func (cl *MovingHandle) Find(c *gin.Context) {
	var moving domain.Moving
	if err := c.ShouldBindUri(&moving); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.Find(c.Request.Context(), moving)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}

// List 查询Moving列表
// @Tags     Moving
// @Summary  查询Moving列表
// @accept   application/json
// @Produce  application/json
// @Security ApiKeyAuth
// @Param    data query     domain.PageMovingSearch true "查询Moving列表"
// @Success  200  {string} string            "{"code":200,"msg":"查询成功","data":{}"}"
// @Router   /moving/list [get]
func (cl *MovingHandle) List(c *gin.Context) {
	var moving domain.PageMovingSearch
	if err := c.ShouldBindQuery(&moving); err != nil {
		response.Error(c, constant.CODE_INVALID_PARAMETER, constant.CODE_INVALID_PARAMETER.Msg())
		return
	}

	res, err := cl.s.List(c.Request.Context(), moving)

	if err != nil {
		response.Error(c, constant.CODE_FIND_FAILED, constant.CODE_FIND_FAILED.Msg())
		return
	}

	response.Success(c, res)
}
