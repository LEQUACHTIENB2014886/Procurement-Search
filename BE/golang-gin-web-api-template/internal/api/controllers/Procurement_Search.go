package controllers

import (
	"net/http"

	"web-api/internal/api/services"
	"web-api/internal/pkg/models/request"
	"web-api/internal/pkg/models/response"

	"github.com/gin-gonic/gin"
)

type ProcController struct {
	*BaseController
}

var Proc = &ProcController{}

func (c *ProcController) GetPurListData(ctx *gin.Context) {
	var requestParams request.ProcListRequest

	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	result, err := services.Proc.GetPurList(requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response.OkWithData(ctx, result)
}

func (c *ProcController) GetAverageDiffDay(ctx *gin.Context) {
	var requestParams request.ProcListRequest

	// Parse và validate request parameters
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.FailWithDetailed(ctx, http.StatusBadRequest, nil, err.Error())
		return
	}

	// Gọi service để tính toán giá trị AverageDiffDay
	averageDiffDay, err := services.Proc.GetAverageDiffDay(requestParams)
	if err != nil {
		response.FailWithDetailed(ctx, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// Trả về kết quả
	response.OkWithData(ctx, gin.H{"averageDiffDay": averageDiffDay})
}
