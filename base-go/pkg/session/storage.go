package session

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	Get(ctx *gin.Context, id string) (*Data, error)
	Save(ctx *gin.Context, id string, data map[string]interface{}, expiredAt time.Time, csrfToken string) error
	Delete(ctx *gin.Context, id string) error
}
