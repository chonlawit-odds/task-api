package item

import (
	"net/http"
	"strconv"

	"github.com/chonlawit-odds/task-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Controller struct {
	Service Service
}

func NewController(dbconn *gorm.DB) Controller {
	return Controller{
		Service: NewService(dbconn),
	}
}

func (ctrl Controller) CreateItem(ctx *gin.Context) {
	// Bind request body
	var (
		request model.RequestCreateItem
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Bind body").Error(),
			},
		)
		return
	}

	// Validate
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}

	// Create
	result, err := ctrl.Service.Create(request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Create item").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[model.Item]{
			Data: result,
		},
	)
}

func (ctrl Controller) GetItems(ctx *gin.Context) {
	// Bind request body
	var (
		request model.RequestFindItem
	)

	if err := ctx.BindQuery(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Bind query").Error(),
			},
		)
		return
	}

	// Validate
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}

	// Find
	results, err := ctrl.Service.Find(request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Find items").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		model.BaseResponse[model.BaseResponseList[[]model.Item]]{
			Data: model.BaseResponseList[[]model.Item]{
				Count:   len(results),
				Results: results,
			},
		},
	)
}

func (ctrl Controller) ReplaceItem(ctx *gin.Context) {
	// Path params
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// Bind request body
	var (
		request model.RequestCreateItem
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Bind body").Error(),
			},
		)
		return
	}

	// Validate
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}

	// Replace
	result, err := ctrl.Service.Replace(uint(id), request)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Replace item").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[model.Item]{
			Data: result,
		},
	)
}

func (ctrl Controller) UpdateItemStatus(ctx *gin.Context) {
	// Path params
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// Bind request body
	var (
		request model.RequestUpdateItem
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Bind body").Error(),
			},
		)
		return
	}

	// Validate
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}

	// Replace
	result, err := ctrl.Service.UpdateStatus(uint(id), request.Status)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Update item status").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[model.Item]{
			Data: result,
		},
	)
}

func (ctrl Controller) DeleteItem(ctx *gin.Context) {
	// Path params
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	// Replace
	if err := ctrl.Service.Delete(uint(id)); err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Update item status").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[any]{
			Message: "success",
		},
	)
}
