package userRepository

import (
	"context"
	"fmt"
	"go-project/internal/api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//var userDB []model.User

type UserRepository interface {
	CreateNewUser(ctx context.Context, newUser model.User) bool
	GetUser(ctx context.Context, email string) (model.User, bool)
}

type MongoUserRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client, dbName, collectionName string) *MongoUserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoUserRepository{client: client, collection: collection}
}

func (r *MongoUserRepository) CreateNewUser(ctx context.Context, newUser model.User) bool {

	_, err := r.collection.InsertOne(ctx, newUser)
	if err != nil {
		return false
	}
	return true
}

func (r *MongoUserRepository) GetUser(ctx context.Context, email string) (model.User, bool) {

	var user model.User

	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return user, false
	}

	fmt.Println(user)
	return user, true
}
