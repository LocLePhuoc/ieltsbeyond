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

var _ writing.AssessmentRepository = (*AssessmentRepository)(nil)

type AssessmentRepository struct {
	client   *mongo.Client
	database *mongo.Database
}

func NewAssessmentRepository() *AssessmentRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("SUBMISSION_REPO_URL")))
	if err != nil {
		logger.Instance.Fatal(err)
	}
	database := client.Database(os.Getenv("SUBMISSION_REPO_DATABASE"))
	return &AssessmentRepository{client: client, database: database}
}

func (s *AssessmentRepository) UpsertWritingAssessment(ctx context.Context, collectionName string, assessment writing.Assessment) (*writing.Assessment, error) {
	collection := s.database.Collection(collectionName) // TODO: fill in collection name

	filter := bson.M{"_id": assessment.Id}
	update := bson.M{"$set": assessment}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (s *AssessmentRepository) GetAssessment(ctx context.Context, collectionName string, assessmentId string) (*writing.Assessment, error) {
	collection := s.database.Collection(collectionName)

	var assessment writing.Assessment
	if err := collection.FindOne(ctx, bson.M{"_id": assessmentId}).Decode(&assessment); err != nil {
		return nil, err
	}
	return &assessment, nil
}
