package dto

import (
	"invertory/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationRequest(context *gin.Context) *domain.Pagination {
	// convert query parameter string to int
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(context.DefaultQuery("page", "0"))
	return &domain.Pagination{Limit: limit, Page: page}
}
