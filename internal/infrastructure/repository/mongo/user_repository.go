package mongo

import (
	"context"
	"go-hexagonal/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	client *mongo.Client
}

// NewMongoUserRepository constructor
func NewMongoUserRepository(client *mongo.Client) *MongoUserRepository {
	return &MongoUserRepository{client: client}
}

// Save saves a new user into the database
func (m *MongoUserRepository) Save(user domain.User) error {
	collection := m.client.Database("mydb").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

// FindByID finds a user by their ID
func (m *MongoUserRepository) FindByID(id string) (*domain.User, error) {
	collection := m.client.Database("mydb").Collection("users")
	var user domain.User
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
