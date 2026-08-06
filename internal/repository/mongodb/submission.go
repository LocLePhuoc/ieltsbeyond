package mongodb

import (
	"context"
	logger "ieltsbeyond/internal/logging"
	"ieltsbeyond/internal/writing"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubmissionRepository struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewSubmissionRepository() *SubmissionRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("SUBMISSION_REPO_URL")))
	if err != nil {
		logger.Instance.Fatal(err)
	}
	database := client.Database(os.Getenv("SUBMISSION_REPO_DATABASE"))
	return &SubmissionRepository{client: client, database: database}
}

func (s *SubmissionRepository) UpsertWritingSubmission(ctx context.Context, collectionName string, submission writing.Submission) (writing.Submission, error) {
	collection := s.database.Collection(collectionName) // TODO: fill in collection name

	filter := bson.M{"_id": submission.Id}
	update := bson.M{"$set": submission}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return writing.Submission{}, err
	}
	return submission, nil
}

func (s *SubmissionRepository) GetSubmission(ctx context.Context, collectionName string, submissionId string) (writing.Submission, error) {
	collection := s.database.Collection(collectionName)

	var submission writing.Submission
	if err := collection.FindOne(ctx, bson.M{"_id": submissionId}).Decode(&submission); err != nil {
		return writing.Submission{}, err
	}
	return submission, nil
}
