package web

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/samber/do"
	"its.ac.id/base-go/bootstrap/config"
	"its.ac.id/base-go/pkg/session/adapters"
)

var ErrProjectIDNotConfigured = errors.New("firestore project ID not configured. please set SESSION_FIRESTORE_PROJECT_ID in .env file")

func setupFirestoreSessionAdapter(i *do.Injector) (*adapters.Firestore, error) {
	ctx := context.Background()
	cfg := do.MustInvoke[config.Config](i).Session()
	if cfg.FirestoreProjectID == "" {
		return nil, ErrProjectIDNotConfigured
	}

	client, err := firestore.NewClient(ctx, cfg.FirestoreProjectID)
	if err != nil {
		return nil, err
	}
	return adapters.NewFirestore(client, cfg.FirestoreCollection), nil
}
