package dao

import (
	"context"
	"fmt"
	"projeto-pismo-api-go/pkg/api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoAccount(client *mongo.Client, db *mongo.Database) *Account {
	return &Account{client: client, collection: db.Collection("accounts")}
}

func (p *Account) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	if account == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, account)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var accountOut model.Account
	err = result.Decode(&accountOut)
	if err != nil {
		return nil, err
	}
	return &accountOut, nil
}

func (p *Account) GetByID(ctx context.Context, id string) (*model.Account, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": oID})
	var account model.Account
	err = result.Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (p *Account) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}
