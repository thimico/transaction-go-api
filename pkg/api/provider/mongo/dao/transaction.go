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

type Transaction struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoTransaction(client *mongo.Client, db *mongo.Database) *Transaction {
	return &Transaction{client: client, collection: db.Collection("transactions")}
}

func (p *Transaction) Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	if transaction == nil {
		return nil, fmt.Errorf("Error on parsing")
	}
	one, err := p.collection.InsertOne(ctx, transaction)
	if err != nil {
		return nil, err
	}
	result := p.collection.FindOne(ctx, bson.M{"_id": one.InsertedID.(primitive.ObjectID)})
	var transactionOut model.Transaction
	err = result.Decode(&transactionOut)
	if err != nil {
		return nil, err
	}
	return &transactionOut, nil
}

func (p *Transaction) Check(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, time.Second)
	return p.client.Ping(ctx, nil)
}
