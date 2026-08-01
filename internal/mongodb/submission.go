package mongodb

import (
	"context"
	logger "ieltsbeyond/internal/logging"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubmissionRepository struct {
	client *mongo.Client
}

func NewSubmissionRepository() *SubmissionRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("SUBMISSION_REPO_URL")))
	if err != nil {
		logger.Instance.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logger.Instance.Fatal(err)
		}
	}()
	return &SubmissionRepository{client: client}
}
