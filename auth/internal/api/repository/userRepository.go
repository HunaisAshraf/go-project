package userRepository

import (
	"context"
	"fmt"
	"go-project/internal/api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//var userDB []model.User

type UserRepository interface {
	CreateNewUser(ctx context.Context, newUser model.User) (model.User, error)
	GetUser(ctx context.Context, email string) (model.User, bool)
	AddToken(ctx context.Context, token model.Token) (model.Token, error)
	GetToken(ctx context.Context, token string) (model.Token, error)
	UpdateToken(ctx context.Context, token model.Token) bool
}

type MongoUserRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
	token      *mongo.Collection
}

func NewMongoUserRepository(client *mongo.Client, dbName, collectionName string) *MongoUserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	token := client.Database(dbName).Collection("TOKENS")
	return &MongoUserRepository{client: client, collection: collection, token: token}
}

func (r *MongoUserRepository) CreateNewUser(ctx context.Context, newUser model.User) (model.User, error) {

	res, err := r.collection.InsertOne(ctx, newUser)
	if err != nil {
		return model.User{}, err
	}
	fmt.Println(res.InsertedID)
	newUser.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return newUser, nil
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

func (r *MongoUserRepository) AddToken(ctx context.Context, token model.Token) (model.Token, error) {
	_, err := r.token.InsertOne(ctx, token)
	if err != nil {
		return model.Token{}, err
	}

	return token, nil
}

func (r *MongoUserRepository) GetToken(ctx context.Context, token string) (model.Token, error) {
	filter := bson.M{"token": token}

	var foundToken model.Token

	err := r.token.FindOne(ctx, filter).Decode(&foundToken)
	if err != nil {
		return foundToken, err
	}

	return foundToken, nil
}

func (r *MongoUserRepository) UpdateToken(ctx context.Context, token model.Token) bool {
	return true
}
