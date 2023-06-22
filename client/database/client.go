package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

type Database struct {
	*zap.Logger
	*mongo.Client
	database string
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func NewDatabase(logger *zap.Logger, url, name string, check ...string) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		<-ctx.Done()
		cancel()
	}()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		logger.Panic("failed to connect to database", zap.Error(err))
	}

	for _, c := range check {
		client.Database(name).CreateCollection(ctx, c)
	}

	return &Database{logger, client, name}
}

func (d *Database) Write(collection string, data any) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = d.Database(d.database).Collection(collection).InsertOne(ctx, data)
	return
}

func (d *Database) Read(collection string, filter bson.M, data any) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return d.Database(d.database).Collection(collection).FindOne(ctx, filter).Decode(&data)
}
