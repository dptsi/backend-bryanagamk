package adapters

import (
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"its.ac.id/base-go/pkg/session"
)

type FirestoreData struct {
	Data      map[string]interface{} `firestore:"data"`
	ExpiredAt time.Time              `firestore:"expired_at"`
	CSRFToken string                 `firestore:"csrf_token"`
}

type Firestore struct {
	client     *firestore.Client
	collection string
}

func NewFirestore(client *firestore.Client, collection string) *Firestore {
	return &Firestore{client, collection}
}

func (f *Firestore) Get(ctx *gin.Context, id string) (*session.Data, error) {
	ref := f.client.Collection(f.collection).Doc(id)
	doc, err := ref.Get(ctx)
	if err != nil && status.Code(err) == codes.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var data FirestoreData
	if err := doc.DataTo(&data); err != nil {
		return nil, err
	}
	if data.ExpiredAt.Before(time.Now()) {
		return nil, nil
	}

	sess := session.NewData(ctx, id, data.CSRFToken, data.Data, f, data.ExpiredAt)
	return sess, nil
}

func (f *Firestore) Save(ctx *gin.Context, id string, data map[string]interface{}, expiredAt time.Time, csrfToken string) error {
	ref := f.client.Collection(f.collection).Doc(id)
	_, err := ref.Set(ctx, FirestoreData{data, expiredAt, csrfToken})
	return err
}

func (f *Firestore) Delete(ctx *gin.Context, id string) error {
	ref := f.client.Collection(f.collection).Doc(id)
	_, err := ref.Delete(ctx)
	return err
}
