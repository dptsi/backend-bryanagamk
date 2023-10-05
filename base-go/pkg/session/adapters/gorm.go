package adapters

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"its.ac.id/base-go/pkg/session"
)

const TableName = "sessions"

type GormData struct {
	Id        string                 `gorm:"primaryKey"`
	Data      map[string]interface{} `gorm:"serializer:json"`
	ExpiredAt time.Time              `gorm:"index"`
	CSRFToken string
}

func (GormData) TableName() string {
	return TableName
}

type Gorm struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) *Gorm {
	db.AutoMigrate(&GormData{})
	return &Gorm{db}
}

func (g *Gorm) Get(ctx *gin.Context, id string) (*session.Data, error) {
	var data GormData
	if err := g.db.Table(TableName).First(&data, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if data.ExpiredAt.Before(time.Now()) {
		return nil, nil
	}

	sess := session.NewData(ctx, id, data.CSRFToken, data.Data, g, data.ExpiredAt)
	return sess, nil
}

func (g *Gorm) Save(ctx *gin.Context, id string, data map[string]interface{}, expiredAt time.Time, csrfToken string) error {
	return g.db.Table(TableName).Save(&GormData{id, data, expiredAt, csrfToken}).Error
}

func (g *Gorm) Delete(ctx *gin.Context, id string) error {
	return g.db.Table(TableName).Delete(&GormData{}, "id = ?", id).Error
}
