package session

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/samber/do"
	"its.ac.id/base-go/bootstrap/config"
)

type Data struct {
	ctx       *gin.Context
	id        string
	csrfToken string
	data      map[string]interface{}
	storage   Storage
	expiredAt time.Time
}

func (d *Data) Id() string {
	return d.id
}

func (d *Data) CSRFToken() string {
	return d.csrfToken
}

func (d *Data) Get(key string) (interface{}, bool) {
	data, ok := d.data[key]
	return data, ok
}

// Setiap set harus disertai dengan save
func (d *Data) Set(key string, value interface{}) {
	d.data[key] = value
}

// Setiap delete harus disertai dengan save
func (d *Data) Delete(key string) {
	delete(d.data, key)
}

// Setiap clear harus disertai dengan save
func (d *Data) Clear() {
	for key := range d.data {
		delete(d.data, key)
	}
}

func (d *Data) Save() error {
	return d.storage.Save(d.ctx, d.id, d.data, d.expiredAt, d.csrfToken)
}

// Setiap regenerate harus disertai dengan save dan add cookie to response
func (d *Data) Regenerate() {
	d.storage.Delete(d.ctx, d.id)
	d.id = uuid.NewString()
}

// Setiap invalidate harus disertai dengan save dan add cookie to response
func (d *Data) Invalidate() {
	d.storage.Delete(d.ctx, d.id)
	d.id = uuid.NewString()
	d.data = make(map[string]interface{})
}

// Setiap regenerate csrf harus disertai dengan save dan add cookie to response
func (d *Data) RegenerateCSRFToken() error {
	d.csrfToken = uuid.NewString()
	return d.Save()
}

func NewEmptyData(ctx *gin.Context, storage Storage) *Data {
	cfg := do.MustInvoke[config.Config](do.DefaultInjector).Session()
	expiredAt := time.Now().Add(time.Second * time.Duration(cfg.Lifetime))

	return &Data{
		ctx:       ctx,
		id:        uuid.NewString(),
		csrfToken: uuid.NewString(),
		data:      make(map[string]interface{}),
		storage:   storage,
		expiredAt: expiredAt,
	}
}

func NewData(ctx *gin.Context, id string, csrfToken string, data map[string]interface{}, storage Storage, expiredAt time.Time) *Data {
	return &Data{
		ctx:       ctx,
		id:        id,
		csrfToken: csrfToken,
		data:      data,
		storage:   storage,
		expiredAt: expiredAt,
	}
}

func Default(ctx *gin.Context) *Data {
	dataIf, exists := ctx.Get("session")
	if !exists {
		panic("session not found in context, make sure you have called session.StartSession middleware")
	}
	data, ok := dataIf.(*Data)
	if !ok {
		panic("session is not of type session.Data")
	}

	return data
}
