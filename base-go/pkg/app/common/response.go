package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"its.ac.id/base-go/pkg/app/common/errors"
)

var UnauthorizedResponse = gin.H{
	"code":    http.StatusUnauthorized,
	"message": "unauthorized",
	"data":    nil,
}

var InternalServerErrorResponse = gin.H{
	"code":    http.StatusInternalServerError,
	"message": "internal_server_error",
	"data":    nil,
}

func AbortAndResponseErrorWithJSON(c *gin.Context, err error) {
	if notFound, isNotFound := err.(*errors.NotFoundError); isNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": notFound.Error(),
			"data":    nil,
		})
		return
	}
	if mismatch, isVersionMismatch := err.(*errors.AggregateVersionMismatchError); isVersionMismatch {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"code":    http.StatusConflict,
			"message": mismatch.Error(),
			"data":    nil,
		})
		return
	}
	if invariantError, isInvariantError := err.(*errors.InvariantError); isInvariantError {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": invariantError.Error(),
			"data":    nil,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerErrorResponse)
}
